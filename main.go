package main

import (
	"github.com/codegangsta/martini"
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	m := martini.Classic()
	m.Get("/:name", func(params martini.Params) []byte {
		b, err := ioutil.ReadFile(params["name"])
		PanicIf(err)
		return blackfriday.MarkdownBasic([]byte(b))
	})

	m.Run()
}
