package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jimweng/bookingsystem/conf"

	. "github.com/jimweng/bookingsystem/internal"
	. "github.com/jimweng/bookingsystem/logger"
	. "github.com/jimweng/bookingsystem/models"
	. "github.com/jimweng/bookingsystem/utils"


	_ "github.com/jimweng/bookingsystem/plugins/inputs/all"
	_ "github.com/jimweng/bookingsystem/plugins/outputs/all"
)

var (
	// HTTPAddr    = flag.String("http", "0.0.0.0:9090", "Address to listen for HTTP requests on")
	confPath    = flag.String("config", "./conf/active.ini", "config location")
	checkcommit = flag.Bool("version", false, "burry code for check version")

	confInfo     *conf.Config
	gitcommitnum string
)

func checkComimit() {
	log.Println(gitcommitnum)
}

func Init() error {
	flag.Parse()
	// if there is a needed to check git commit num ... print it out and exit
	if *checkcommit {
		checkComimit()
		os.Exit(1)
	}

	// read config and pass variables ...
	var err error
	confInfo, err = conf.InitConfig(confPath)
	if err != nil {
		return fmt.Errorf("Init config err: %v", err)
	}

	// initialize logger
	if err = InitLog(confInfo.LogConf.LogPath, confInfo.LogConf.LogLevel); err != nil {
		return fmt.Errorf("init logger err: %v", err)
	}

	// initialize both mysql and redis db
	if err = InitDb(&confInfo.DbConf, &confInfo.RedisConf); err != nil {
		return fmt.Errorf("init db err: %v", err)
	}

	// initialize workerpools
	if err :=InitWorkerPools(confInfo.WorkersPool.Tag, confInfo.WorkersPool.NumWorkers); err != nil{
		return fmt.Errorf("init workers pool err: %v", err)		
	}

	return nil
}

func main() {
	//catch global panic
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic err: %v", err)
		}
	}()

	err := Init()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	StartAgent()

	GracefulShutdown(StopDispatcher)
}

