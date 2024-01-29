package main

import (
	"embed"
	"log"

	"github.com/andyinabox/linkydink-sketch/pkg/adapters/server"
	"github.com/andyinabox/linkydink-sketch/pkg/domain"
)

//go:embed res/*
var res embed.FS

func main() {

	s := server.NewServer(&server.ServerConfig{
		Host:           "127.0.0.1",
		Port:           "8000",
		EmbedFS:        res,
		TemplatesGlob:  "res/tmpl/*.tmpl",
		StaticDirName:  "/static/",
		EmbedFSRootDir: "res",
	})

	s.Route("/", domain.GetIndex, &server.RouteOptions{})

	log.Fatal(s.Serve())
}
