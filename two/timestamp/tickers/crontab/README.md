# intro
使用ticker製作一個crontab定時去拿取數據，並且把數據送到kafka。
> 透過consumer消費kafka裡面的topic可以拿到對應的數據

# architecture
ticker -> sendRequest -> sendDataToKafka
- 驗證: 透過製作一個consumer消費kafka的topic來看數據是否有正確送達
> docker exec -it kafka_name ./opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic sarama

# extend-refer:
- https://medium.com/@petermelias/kafka-consumer-patterns-and-gotchas-1bfc04cd643b
- https://godoc.org/github.com/Shopify/sarama?fbclid=IwAR3nPogQ9ps9s_i4NtFJ5ql5rC4zdLVZiiM_9H79Gb9bwmyAqXzexZZmURI