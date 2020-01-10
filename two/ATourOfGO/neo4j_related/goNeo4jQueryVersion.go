package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// claim node : would declare properties that would be writed to neo4j
type neo4jNodeInfo struct {
	NodeNum  string
	DomainID string
	Name     string
	Labels   string
	Link     string
	Relation string
}

func main() {
	// assume telegraf.Metric body would be like this
	// neo4j telegraf.Metric would be pass with map[name]field ; e.g. neo4j,nodeNum=n1,Labels=testTag ,domainID=testDomainID, name=testName,link=n2_n3 (rmk: n2 belong to n3),relation=belong
	var neo4jInput = map[string]neo4jNodeInfo{
		"n1": {"n1", "testDomainID", "testName", "testTag", "n1_n2", "belong"},
		"n2": {"n2", "testDomainID2", "testName2", "testTag3", "n2_n3", "follow"},
		"n3": {"n3", "testDomainID3", "testName3", "testTag3", "n3_n1", "take"},
	}
	// var keys []string
	// for k, v := range []string{"n1", "n2", "n3"} {
	// for v := range neo4jInput {
	// 	fmt.Println(v)
	// 	fmt.Println(reflect.TypeOf(v))
	// }
	// fmt.Println(len(neo4jInput))

	// fmt.Println(keys != nil)

	neo4jQueryNodes(neo4jInput["n1"])
	// neo4jQueryLink(neo4jInput["n1"])
	// fmt.Println(neo4jQueryLink(neo4jInput["n1"]))

	// v,i:=range neo4jInput
	// fmt.Println(neo4jInput.len)

}

func neo4jQueryNodes(queryNodes neo4jNodeInfo) bool {

	criterion := false

	type Payload struct {
		Query string `json:"query"`
	}
	data := Payload{
		Query: "match(n:" + queryNodes.Labels + ") where n.domainId='" + queryNodes.DomainID + "' and n.name='" + queryNodes.Name + "' return n.domainId",
	}
	fmt.Println(data)
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://172.31.86.190:7474/db/data/cypher", body)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth("neo4j", "na")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	// var f map[string]interface{}
	var f interface{}

	jsonResponse := f
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &jsonResponse)
	// print out jsonResponse
	fmt.Println(jsonResponse)
	fmt.Println(jsonResponse.(map[string]interface{})["data"].([]interface{})[0].([]interface{})[0])
	// fmt.Println(queryNodes.DomainID)

	criterion = (jsonResponse.(map[string]interface{})["data"].([]interface{})[0].([]interface{})[0] == queryNodes.DomainID)
	fmt.Println(criterion)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	return criterion
}

func neo4jQueryLink(queryNodes neo4jNodeInfo) bool {
	type Payload struct {
		Query string `json:"query"`
	}
	data := Payload{
		Query: "match p=(n1)-[r]->(n2) where type(r)='" + queryNodes.Relation + "' return n1.domainId,n2.domainId ,type(r)",
	}
	// fmt.Println(queryNodes.DomainID)
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://172.31.86.190:7474/db/data/cypher", body)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth("neo4j", "na")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	// var f map[string][]interface{}
	var f interface{}
	jsonResponse := f
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &jsonResponse)
	// fmt.Println(jsonResponse["data"].([]interface{}))
	// fmt.Println(reflect.TypeOf(jsonResponse["data"]))
	// criterion := true
	criterion := (jsonResponse.(map[string]interface{})["data"].([]interface{})[0].([]interface{})[0] == queryNodes.DomainID)

	// value, criterion := jsonResponse.(map[string]interface{})["data"]
	// fmt.Println(value)

	// criterion := (jsonResponse.(map[string]interface{})["data"] != nil)

	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	return criterion
	// return true
}
