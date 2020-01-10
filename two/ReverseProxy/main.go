package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-log"
)

var logger *log.Logger = log.NewOrExit()

func main() {
	var level, listen, backend string
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.StringVar(&listen, "listen", "0.0.0.0:80", "Listen for connections on this address.")
	flags.StringVar(&backend, "backend", "0.0.0.0:8080", "The address of the backend to forward to.")
	flags.StringVar(&level, "level", "info", "The logging level.")
	flags.Parse(os.Args[1:])
	logger.SetLevel(log.NewLevel(level))

	if listen == "" || backend == "" {
		fmt.Fprintln(os.Stderr, "listen and backend options required")
		os.Exit(1)
	}

	p := Proxy{Listen: listen, Backend: backend}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		if err := p.Close(); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	if err := p.Run(); err != nil {
		logger.Fatal(err.Error())
	}
}

func Pipe(a, b net.Conn) error {
	done := make(chan error, 1)

	cp := func(r, w net.Conn) {
		n, err := io.Copy(r, w)
		logger.Debugf("copied %d bytes from %s to %s", n, r.RemoteAddr(), w.RemoteAddr())
		done <- err
	}

	go cp(a, b)
	go cp(b, a)
	err1 := <-done
	err2 := <-done
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

// Proxy connections from Listen to Backend.
type Proxy struct {
	Listen   string
	Backend  string
	listener net.Listener
}

func (p *Proxy) Run() error {
	var err error
	if p.listener, err = net.Listen("tcp", p.Listen); err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	for {
		if conn, err := p.listener.Accept(); err == nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				p.handle(conn)
			}()
		} else {
			return nil
		}
	}
	wg.Wait()
	return nil
}

func (p *Proxy) Close() error {
	return p.listener.Close()
}

func (p *Proxy) handle(upConn net.Conn) {
	defer upConn.Close()
	logger.Infof("accepted: %s", upConn.RemoteAddr())
	downConn, err := net.Dial("tcp", p.Backend)
	if err != nil {
		logger.Errorf("unable to connect to %s: %s", p.Backend, err)
		return
	}
	defer downConn.Close()
	if err := Pipe(upConn, downConn); err != nil {
		logger.Errorf("pipe failed: %s", err)
	} else {
		logger.Infof("disconnected: %s", upConn.RemoteAddr())
	}
}
