## HTTP API
## Create node
```
curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:na -d '{
  "statements": [{
    "statement":"CREATE(n) RETURN id(n)"
  }]
}' http://172.31.86.190:7474/db/data/transaction/commit
```
## Execute multiple statements
```
curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:na -d '{
  "statements": [{
    "statement":"CREATE(n) RETURN id(n)"
  },{
    "statement":"CREATE(n {props}) RETURN n",
    "parameters":{
      "props":{
        "name":"My node"
      }
    }
  }]
}' http://172.31.86.190:7474/db/data/transaction/commit
```
## Return results in graph format
```
curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:na -d '{
  "statements" : [ {
    "statement" : "CREATE ( bike:Bike { weight: 10 } ) CREATE ( frontWheel:Wheel { spokes: 3 } ) CREATE ( backWheel:Wheel { spokes: 32 } ) CREATE p1 = (bike)-[:HAS { position: 1 } ]->(frontWheel) CREATE p2 = (bike)-[:HAS { position: 2 } ]->(backWheel) RETURN bike, p1, p2",
    "resultDataContents" : [ "row", "graph" ]
  }]
}' http://172.31.86.190:7474/db/data/transaction/commit
```
## n1:Leader n2:Member n3:Member 後面的group表示標籤tag
```
curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:na -d '{
  "statements" : [ {
    "statement" : "CREATE ( n1:Leader { domainId: 127 , name: 123456} ) CREATE ( n2:Member { domainId: 3 } ) CREATE ( n3:Member { domainId: 32 } ) CREATE p1 = (n1)-[:HAS]->(n2) CREATE p2 = (n1)-[:HAS]->(n3) RETURN n1, p1, p2",
    "resultDataContents" : [ "row", "graph" ]
  }]
}' http://172.31.86.190:7474/db/data/transaction/commit
```