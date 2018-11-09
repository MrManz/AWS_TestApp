package main

import (
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()
	static := martini.Static("public", martini.StaticOptions{IndexFile: "index.html", Exclude: "/data"})
	m.NotFound(static, http.NotFound)
	m.Use(static)
	m.Get("/data", func() string {
		return "Hello AWS world!"
	})
	m.Run()
}
