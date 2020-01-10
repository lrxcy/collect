package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// claim node : would declare properties that would be writed to neo4j
type nodeInfo struct {
	NodeNum  string
	DomainID string
	Name     string
	TAG      string
	Link     string
	Relation string
}

func neo4jBodyString(c nodeInfo) string {
	var bodyString = "merge (" + c.NodeNum + ":" + c.TAG + " {domainId:'" + c.DomainID + "', name:'" + c.Name + "'})"
	return bodyString
}

func neo4jLinkString(c nodeInfo) string {
	var linkString string
	if len(strings.Split(c.Link, "|")) == 1 {
		linkString = "merge " + c.Link + " = (" + strings.Split(c.Link, "_")[0] + ")-[:" + c.Relation + "]->(" + strings.Split(c.Link, "_")[1] + ")"
	} else {
		linkArray := strings.Split(c.Link, "|")
		linkRelation := strings.Split(c.Relation, "|")
		for i := 0; i < len(linkArray); i++ {
			tempString := "merge " + linkArray[i] + " = (" + strings.Split(linkArray[i], "_")[0] + ")-[:" + linkRelation[i] + "]->(" + strings.Split(linkArray[i], "_")[1] + ")"
			linkString = linkString + " " + tempString
		}
	}
	fmt.Println(linkString)
	return linkString
}

func CreateNodes(s string) {

	type Payload struct {
		Query string `json:"query"`
	}

	data := Payload{
		Query: s,
	}

	payloadBytes, err := json.Marshal(data)

	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://127.0.0.1:7474/db/data/cypher", body)
	if err != nil {
		// handle err
	}

	req.SetBasicAuth("neo4j", "na")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func main() {

	// assume telegraf.Metric body would be like this
	// neo4j telegraf.Metric woudl be pass with map[name]field ; e.g. neo4j,nodeNum=n1,tag=testTag ,domainID=testDomainID, name=testName,link=n2_n3 (rmk: n2 belong to n3),relation=belong
	var testString = "mem,host=MacPro used_percent=57.654523849487305,total=8589934592i,cached=0i,inactive=3101143040i,slab=0i,active=3521380352i,available_percent=42.345476150512695,available=3637448704i,used=4952485888i,free=536305664i,buffered=0i 1513149940000000000\ndisk,path=/,device=disk1s1,fstype=apfs,mode=rw,host=MacPro inodes_free=9223372036853631618i,inodes_used=1144189i,total=119824367616i,free=71182340096i,used=42605252608i,used_percent=37.44279283491888,inodes_total=9223372036854775807i 1513149940000000000\nsystem,host=MacPro load15=3.35,n_users=1i,n_cpus=4i,load1=4.48,load5=4.86 1513566990000000000\nsystem,host=MacPro uptime=6809i,uptime_format= 1513566990000000000,\nneo4j,nodeNum=n1,tag=testTag,domainID=testDomainID, name=testName,link=n1_n2|n3_n1,relation=belong|take,nodeNum=n2,tag=testTag2,domainID=testDomainID2, name=testName2,link=n2_n3,relation=follow,nodeNum=n3,tag=testTag3,domainID=testDomainID3, name=testName3,link=n1_n3,relation=after"

	// define commitArray
	var parseBodyString string
	var parselinkString string
	var mesurementStringArray = strings.Split(testString, "\n")
	// create request body
	for index := 0; index < len(mesurementStringArray); index++ {
		neededMeasurement := strings.Split(mesurementStringArray[index], ",")
		if neededMeasurement[0] == "neo4j" {
			for i := 0; i < (len(neededMeasurement)-1)/6; i++ {
				neo4jNum := strings.Split(neededMeasurement[6*i+1], "=")[1]
				neo4jTag := strings.Split(neededMeasurement[6*i+2], "=")[1]
				neo4jDomainID := strings.Split(neededMeasurement[6*i+3], "=")[1]
				neo4jName := strings.Split(neededMeasurement[6*i+4], "=")[1]
				neo4jlink := strings.Split(neededMeasurement[6*i+5], "=")[1]
				neo4jRelation := strings.Split(neededMeasurement[6*i+6], "=")[1]
				nodesCreated := nodeInfo{
					NodeNum:  neo4jNum,
					TAG:      neo4jTag,
					DomainID: neo4jDomainID,
					Name:     neo4jName,
					Link:     neo4jlink,
					Relation: neo4jRelation,
				}
				neo4jBodyString(nodesCreated)

				if !(neo4jlink == "") {
					parselinkString = parselinkString + " " + neo4jLinkString(nodesCreated)
				}
				parseBodyString = parseBodyString + " " + neo4jBodyString(nodesCreated)
			}
			parseBodyString = parseBodyString + parselinkString
			CreateNodes(parseBodyString)
		}
	}
}
