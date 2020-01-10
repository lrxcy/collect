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
	domainID string
	name     string
	TAG      string
	link     map[string]string
}

func CreateMultiNodes(cA []nodeInfo) {
	for index := 0; index < len(cA); index++ {
		c := cA[index]
		CreateNodes(c)
	}
}

func CreateNodes(c nodeInfo) {

	// Define create one node strings
	var oneNodeString = "create (n1:" + c.TAG + " {domainId:'" + c.domainID + "', name:'" + c.name + "'})"
	type Payload struct {
		Query string `json:"query"`
	}

	data := Payload{
		Query: oneNodeString,
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
	// neo4j telegraf.Metric woudl be pass with map[name]field ; e.g. neo4j,tag=testTag ,domainID=testDomainID, name=testName,link=testLink
	var testString = "mem,host=MacPro used_percent=57.654523849487305,total=8589934592i,cached=0i,inactive=3101143040i,slab=0i,active=3521380352i,available_percent=42.345476150512695,available=3637448704i,used=4952485888i,free=536305664i,buffered=0i 1513149940000000000\ndisk,path=/,device=disk1s1,fstype=apfs,mode=rw,host=MacPro inodes_free=9223372036853631618i,inodes_used=1144189i,total=119824367616i,free=71182340096i,used=42605252608i,used_percent=37.44279283491888,inodes_total=9223372036854775807i 1513149940000000000\nsystem,host=MacPro load15=3.35,n_users=1i,n_cpus=4i,load1=4.48,load5=4.86 1513566990000000000\nsystem,host=MacPro uptime=6809i,uptime_format= 1513566990000000000,\nneo4j,tag=testTag,domainID=testDomainID, name=testName,link=testLink,tag=testTag2,domainID=testDomainID2, name=testName2,link=testLink,tag=testTag3,domainID=testDomainID3, name=testName3,link=testLink"

	// define commitArray
	var mesurementStringArray = strings.Split(testString, "\n")
	// create request body also make criterion to choose

	for index := 0; index < len(mesurementStringArray); index++ {
		neededMeasurement := strings.Split(mesurementStringArray[index], ",")
		if neededMeasurement[0] == "neo4j" {
			if len(neededMeasurement) > 4 {
				fmt.Println("multinodes")
				for i := 0; i < (len(neededMeasurement)-1)/4; i++ {
					neo4jTag := strings.Split(neededMeasurement[4*i+1], "=")[1]
					neo4jDomainId := strings.Split(neededMeasurement[4*i+2], "=")[1]
					neo4jName := strings.Split(neededMeasurement[4*i+3], "=")[1]
					nodesCreated := nodeInfo{
						TAG:      neo4jTag,
						domainID: neo4jDomainId,
						name:     neo4jName,
					}
					CreateNodes(nodesCreated)
				}
			} else {
				fmt.Println("one node")
				neo4jTag := strings.Split(neededMeasurement[1], "=")[1]
				neo4jDomainId := strings.Split(neededMeasurement[2], "=")[1]
				neo4jName := strings.Split(neededMeasurement[3], "=")[1]
				nodesCreated := nodeInfo{
					TAG:      neo4jTag,
					domainID: neo4jDomainId,
					name:     neo4jName,
				}
				CreateNodes(nodesCreated)
			}
		}
	}
}
