# quick start
1. 啟動docker-compose帶起對應的kafka以及zookeeper
> docker-compose up -d

2. 啟動producer.go往已經建立好的topic(sarama)送資料
> go run producer.go localhost:9092 sarama

3. 啟動consumer.go來消費線上的kafka topic
> go run consumer.go localhost:9092 1 sarama

# refer:
- https://github.com/Shopify/sarama
- https://juejin.im/post/5d40f179f265da038f47e9eb

# kafka-docker-refer:
- https://github.com/wurstmeister/kafka-docker


 
# Automatically create topics
If you want to have kafka-docker automatically create topics in Kafka during creation, a 
KAFKA_CREATE_TOPICS environment variable can be added in docker-compose.yml.

Here is an example snippet from docker-compose.yml:
```sh
    environment:
      KAFKA_CREATE_TOPICS: "Topic1:1:3,Topic2:1:1:compact"
```

Topic 1 will have 1 partition and 3 replicas,
Topic 2 will have 1 partition, 1 replica and a cleanup.policy set to compact. 
