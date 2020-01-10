package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetJsonConfig(folderPath string) (*[]string, error) {
	filelist := []string{}
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		filelist = append(filelist, fmt.Sprintf("%s", folderPath+"/"+f.Name()))
	}

	return &filelist, nil
}

func ReadJson(path string) (interface{}, error) {
	confBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var f interface{}
	jsonStruct := f
	if err = json.Unmarshal(confBytes, &jsonStruct); err != nil {
		return nil, err
	}
	fmt.Printf("Test value of jsonStruct %v\n", jsonStruct)
	return jsonStruct, nil
}

func DumpJsonValue(jsonStruct interface{}) {
	structJson := make(map[string]interface{})

	switch jsonStruct.(type) {
	case map[string]interface{}:
		structJson = jsonStruct.(map[string]interface{})

	default:
		fmt.Println("another type")
	}

	for i, j := range structJson {
		fmt.Printf("%v___%v\n", i, j)
	}
}
