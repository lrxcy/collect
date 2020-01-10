## Cypher API
## curl write data to influxdb db=mydb
> curl -i -XPOST 'http://172.31.86.190:8086/write?db=mydb' --data-binary 'cpu_load_short,host=server01,region=us-west value=0.64 1434015562000000000'

## curl write data to noetj db=mydb 
> curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --data-binary '{"query":"match(n) return n"}' http://172.31.86.190:7474/db/data/cypher?u=neo4j&p=na

> curl -i XPOST -H "Accept: application/json" -H "Content-Type: application/json" -H "X-stream: true" --user neo4j:na http://172.31.86.190:7474/db/data/





