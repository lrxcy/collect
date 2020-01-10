package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// 背景執行: 開始執行ticker，每500 milesecond即執行一次
	go func() {
		for {
			select {
			// 如果done沒有值，則不進行return，如果return則結束執行程序;使用done來控制背景線程，間接控制整個主線程
			case <-done:
				return

			// 當ticker有值進來的時候進行打印
			case t := <-ticker.C:
				fmt.Println("Tick at ", t)
				res := sendRequest()
				send("localhost:9092", "sarama", []byte(res))
			}
		}
	}()

	// 主線程停1600秒後停止ticker，並且傳送bool(true)給done，結束整個線程
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func sendRequest() string {
	req, _ := http.NewRequest("GET", "http://jimqaweb.mlytics.ai/cache.txt", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(bodyBytes))
	return string(bodyBytes)

}

var signals = make(chan os.Signal, 1)

// url: borker, topic: consumed topic, chars: sended data
func send(url string, topic string, chars []byte) {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewAsyncProducer([]string{url}, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	signal.Notify(signals, os.Interrupt)

	var enqueued, errors int
	doneCh := make(chan struct{})

	// buf := make([]byte, 1024)
	// for i := 0; i < 4; i++ {
	// buf[i] = chars[rand.Intn(len(chars))]
	// }

	strTime := strconv.Itoa(int(time.Now().Unix()))
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(strTime),
		Value: sarama.StringEncoder(),
	}
	select {
	case producer.Input() <- msg:
		enqueued++
		fmt.Printf("Produce message: %s\n", buf)
	case err := <-producer.Errors():
		errors++
		fmt.Println("Failed to produce message:", err)
	case <-signals:
		doneCh <- struct{}{}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)

}
