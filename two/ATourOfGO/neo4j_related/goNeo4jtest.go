package main

import (
	"fmt"
)

// claim node : would declare properties that would be writed to neo4j
type nodeInfo struct {
	NodeNum  string
	Type     string
	DomainID string
	Name     string
	Labels   string
	// Link     string
	// Relation string
}

// declare to parse json format while use NewRequest
type Payload struct {
	Query string `json:"query"`
}

func neo4jBodyString(c nodeInfo) string {
	var bodyString = "merge (" + c.NodeNum + ":" + c.Labels + " {domainId:'" + c.DomainID + "', name:'" + c.Name + "'})"
	return bodyString
}

// func neo4jLinkString(c nodeInfo) string {
// 	var linkString string
// 	if len(strings.Split(c.Link, "|")) == 1 {
// 		linkString = "merge " + c.Link + " = (" + strings.Split(c.Link, "_")[0] + ")-[:" + c.Relation + "]->(" + strings.Split(c.Link, "_")[1] + ")"
// 	} else {
// 		linkArray := strings.Split(c.Link, "|")
// 		linkRelation := strings.Split(c.Relation, "|")
// 		for i := 0; i < len(linkArray); i++ {
// 			tempString := "merge " + linkArray[i] + " = (" + strings.Split(linkArray[i], "_")[0] + ")-[:" + linkRelation[i] + "]->(" + strings.Split(linkArray[i], "_")[1] + ")"
// 			linkString = linkString + " " + tempString
// 		}
// 	}
// 	return linkString
// }

func neo4jMainFunc(m map[string]nodeInfo) string {
	// define parse body
	var parseBodyString string
	var parselinkString string
	for value := range m {

		nodesCreated := nodeInfo{
			NodeNum:  m[value].NodeNum,
			Type:     m[value].Type,
			DomainID: m[value].DomainID,
			Name:     m[value].Name,
			Labels:   m[value].Labels,
		}

		fmt.Println(nodesCreated)
		// 	// if !(nodesCreated.Link == "") {
		// 	// 	parselinkString = parselinkString + " " + neo4jLinkString(nodesCreated)
		// 	// }
		// 	// parseBodyString = parseBodyString + " " + neo4jBodyString(nodesCreated)
	}
	parseBodyString = parseBodyString + parselinkString
	return parseBodyString
}

func main() {
	// var onNodeString string

	var neo4jInput = map[string]nodeInfo{
		// "n1": {"n1", "testDomainID", "testName", "testTag", "n1_n2", "belong"},
		// "n2": {"n2", "testDomainID2", "testName2", "testTag2", "n2_n3", "follow"},
		// "n3": {"n3", "testDomainID3", "testName3", "testTag3", "n3_n1", "take"},
		"n1": {"n1", "node", "testDomainID", "testName", "testTag"},
		"n2": {"n2", "node", "testDomainID2", "testName2", "testTag2"},
		"n3": {"n3", "node", "testDomainID3", "testName3", "testTag3"},
		// "r1": {"n1_n2", "belong"},
		// "r2": {"n2_n3", "follow"},
		// "r3": {"n3_n1", "take"},
	}

	fmt.Println("neo4jInput is", neo4jInput)
	fmt.Println("----------------------------------------------------------------")

	neo4jMainFunc(neo4jInput)

	// data := Payload{
	// 	Query: onNodeString,
	// }
	// fmt.Println(data)
}
