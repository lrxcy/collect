// retrun a map with both keys and values
package main

func main() {
	// testArray := [1,2,3,4,5]
	tags := map[string]string{
		"host": "localhost",
	}
	for k, v := range tags {
		println(k)
		println(v)
	}

}
