package main

import (
	"fmt"
	"net/url"
)

type WriteParams struct {
	Database        string
	RetentionPolicy string
	Precision       string
	Consistency     string
}

func main() {
	wp := WriteParams{Database: "d1", RetentionPolicy: "1d", Precision: "95", Consistency: "any"}
	fmt.Println(wp)
	params := url.Values{}

	fmt.Println("Before params.Set, the map should be empty", params)
	// add database inside params with wp.Database value but "db" name
	params.Set("db", wp.Database)
	fmt.Println("After params.Set, the map should be filled with map[db:[d1]]", params)

	params.Encode()
	fmt.Println(params)

}
