package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	static := martini.Static("public", martini.StaticOptions{IndexFile: "index.html"})
	m.Use(static)
	m.Run()
}
