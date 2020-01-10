// http://spf13.com/post/is-go-object-oriented/
package main

import "time"
import "fmt"

// claim a time Duration struct
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalTOML(b []byte) error {
	dur, err := time.ParseDuration(string(b[1 : len(b)-1]))
	if err != nil {
		return err
	}

	d.Duration = dur

	return nil
}

func main() {
	test := Duration{time.Hour}
	fmt.Println(test)

	fmt.Println("UnmarshalTOML value is", test.UnmarshalTOML([]byte("12ms")))
}
