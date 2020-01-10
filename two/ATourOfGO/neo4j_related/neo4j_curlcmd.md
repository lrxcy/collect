## refer website
- https://neo4j.com/docs/rest-docs/current/#rest-api-service-root
- https://neo4j.com/blog/the-neo4j-rest-server-part1-get-it-going/

## This command would query mine eno4j db with --user loginID:passWD
## ,however this is the basic of neo4j need to implement some detail criterion.
> curl -i XPOST --user neo4j:password http://172.31.86.190:7474/db/data/
> curl -H "Content-Type: text/plain" -X POST --user neo4j:password -o result.tde http://172.31.86.178:7474/db/data/
## add an .json file to note the info of query item
> curl -i XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:password http://172.31.86.178:7474/db/data/

## add X-stream to make lower memory overhead
> curl -i XPOST -H "Accept: application/json" -H "Content-Type: application/json" -H "X-stream: true" --user neo4j:password http://172.31.86.190:7474/db/data/

## addd --data 
> curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:password -o output.log --data @create.json http://172.31.86.178:7474/db/data/cypher

## another way to create node instead of create a .json file
## CREATE (n:Person { name : {name} }) RETURN n
```
curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:password -d '{
  "query": "CREATE (n:Person { name : {name} }) RETURN n",
  "params": {
    "name": "John"
  }
}' http://172.31.86.178:7474/db/data/cypher
```
## MATCH(n) RETURN n
> curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" --user neo4j:na -d '{"query":"match(n) return n"}' http://172.31.86.190:7474/db/data/cypher

> curl -i -XPOST -H "Accept: application/json" -H "Content-Type: application/json" -u neo4j:na --data-binary '{"query":"match(n) return n"}' http://172.31.86.190:7474/db/data/cypher

## http://172.31.86.190:7474/browser/?u=neo4j&p=na

## Cypher langeuage interface operation
## add an extra .json file 
# create.json
```
{
  "query": "CREATE (n:Person { name : {name} }) RETURN n",
  "params": {
    "name": "Jim"
  }
}
```

## show all nodes
match(n) return n

## create an node
CREATE (n:Person { name : "create" }) RETURN n

## delete all nodesgo gogogo
match(n) detach delete n

## create a relationship
match(a:Person),(b:Person)
where a.name="create" and b.name="John"
create (b)-[r:enemy]->(a)

## create nodes and relationship together
```
CREATE (js:Person { name: "java script", from: "Sweden", learn: "surfing" }),
(ir:Person { name: "i read a book", from: "England", title: "author" }),
(rvb:Person { name: "read virtual book", from: "Belgium", pet: "Orval" }),
(ally:Person { name: "a lulu yy", from: "California", hobby: "surfing" }),
(ir)-[:KNOWS {since: 2001}]->(js)-[:KNOWS {rating: 5}]->(ir),
(js)-[:KNOWS]->(ir),(js)-[:KNOWS]->(rvb),
(ir)-[:KNOWS]->(js),(ir)-[:KNOWS]->(ally),
(rvb)-[:KNOWS]->(ally)
```