package main

import (
	"embed"
	"log"

	"github.com/andyinabox/linkydink-sketch/pkg"
)

//go:embed res/*
var res embed.FS

func main() {

	server := pkg.NewServer(&pkg.ServerConfig{
		Host:          "127.0.0.1",
		Port:          "8000",
		Res:           res,
		TemplatesGlob: "res/tmpl/*.tmpl",
	})

	log.Fatal(server.Run())
}
