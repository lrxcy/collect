// Refer to website
// https://studygolang.com/articles/3496
// http://www.sharejs.com/codes/go/4402
// https://blog.golang.org/go-maps-in-action

package main

import (
	"fmt"
)

// claim a struct obj to use Map
type MapStruc struct {
	Lat  string
	Long float64
}

func main() {

	// claim variable m as map with structure map[string]MapStruc
	var m map[string]MapStruc

	// use 'make' to build a dynamic array (obj)
	// func make(Type, size IntegerType) Type
	// obbiously, it would return a Type alike input Type
	m = make(map[string]MapStruc)

	// Insert some value into map structure
	m["Height Jim"] = MapStruc{
		"jim", 171,
	}
	m["Weight Jim"] = MapStruc{
		"jim", 61,
	}

	fmt.Println(m["Weight Jim"])
	fmt.Println(m)

	// another method to claim a map
	var m_1 map[int]MapStruc
	m_1 = map[int]MapStruc{
		1: {"jim", 170},
		2: {"bob", 171},
	}

	fmt.Println(m_1)

	// delete map element
	delete(m_1, 1)

	fmt.Println(m_1)

	// check value with key
	tempValue, ok := m_1[2]
	fmt.Println("The value:", tempValue, "Present?", ok)
	tempValue2, crit := m_1[1]
	fmt.Println("The value:", tempValue2, "Present?", crit)
}
