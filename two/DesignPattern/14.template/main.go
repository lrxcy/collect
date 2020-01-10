package main

import "fmt"

type Downloader interface {
	Download(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Println("Prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Println("Finish downloading\n")
}

func main() {

}
