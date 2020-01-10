## query with cypher syntax match(n) return n 
curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:na -d '{"query":"match(n) return n"}' http://172.31.86.190:7474/db/data/cypher

## create two nodes with form after -d '{ }'
```
curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:na -d '{
  "query": "CREATE (n1:Person { domainId : {domainId1} , name : {name1}}) CREATE (n2:Person { domainId : {domainId2} , name : {name2}}) CREATE p1 = (n1)-[:HAS]->(n2) RETURN p1",
  "params": {
    "domainId1": "Global",
    "name1":"Esxi",
    "domainId2": "sub-domainID",
    "name2":"Ubuntu"
  }
}' http://172.31.86.190:7474/db/data/cypher
```
## search nodes and create relationship
```
curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:na -d '{
  "query": "match (n1:Person { domainId : {domainId1} , name : {name1}}) match (n2:Person { domainId : {domainId2} , name : {name2}}) CREATE p2 = (n2)-[:BELONG]->(n1) RETURN p2",
  "params": {
    "domainId1": "Global",
    "name1":"Esxi",
    "domainId2": "sub-domainID",
    "name2":"Ubuntu"
  }
}' http://172.31.86.190:7474/db/data/cypher
```

