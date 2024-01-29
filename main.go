package main

import (
	"embed"
	"log"

	"github.com/andyinabox/linkydink-sketch/app"
	"github.com/andyinabox/linkydink-sketch/pkg/simpleserver"
)

//go:embed res/*
var res embed.FS

func main() {

	server := simpleserver.NewServer(&simpleserver.Config{
		Host:           "127.0.0.1",
		Port:           "8000",
		EmbedFS:        res,
		TemplatesGlob:  "res/tmpl/*.tmpl",
		StaticDirName:  "/static/",
		EmbedFSRootDir: "res",
	})

	server.Route("/", app.GetIndex, &simpleserver.RouteOptions{})

	log.Fatal(server.Serve())
}
