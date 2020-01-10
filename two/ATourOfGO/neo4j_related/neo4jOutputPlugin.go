package neo4j

// testing plugin

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/outputs"
)

// Neo4j struct is the primary data structure for the plugin
type Neo4j struct {
	// URL is only for backwards compatibility
}

var sampleConfig = `
namepass = ["db_relay"]
`

func (i *Neo4j) Description() string {
	return "a neo4j output"
}

func (i *Neo4j) SampleConfig() string {
	return sampleConfig
}

func (i *Neo4j) Connect() error {
	return nil
}

func (i *Neo4j) Close() error {
	// Close connection to the URL here
	return nil
}

func (i *Neo4j) Write(metrics []telegraf.Metric) error {
	// make sure only ["neo4j"] measurement would be passed
	// fmt.Printf("%v\n", metrics)
	for _, neo4jMap := range metrics {
		// neo4j implement
		neo4jInputFields := neo4jMap.Fields()
		tempString := neo4jInputFields["cmd"].(string)
		CreateNodes(tempString)

		// fmt.Printf("%v\n", tempString)
		// fmt.Println("================")
	}
	return nil
}

func init() {
	outputs.Add("neo4j", func() telegraf.Output { return &Neo4j{} })
}

func CreateNodes(s string) {
	// Define create one node strings
	var oneNodeString = s
	type Payload struct {
		Query string `json:"query"`
	}
	data := Payload{
		// fill struct
		Query: oneNodeString,
	}
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

	if err != nil {
		// handle err
	}

	defer resp.Body.Close()
}
