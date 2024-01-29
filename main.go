package main

import (
	"embed"
	"log"

	"github.com/andyinabox/linkydink/app"
)

//go:embed res/*
var res embed.FS

func main() {

	appInstance := app.New(&app.Config{
		Host:      "127.0.0.1",
		Port:      "8000",
		Resources: res,
	})

	log.Fatal(appInstance.Start())
}
