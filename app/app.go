package app

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"time"

	"github.com/andyinabox/linkydink/pkg/feedreader"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type App struct {
	conf       *Config
	router     *gin.Engine
	feedreader *feedreader.Reader
	db         *gorm.DB
}

type Config struct {
	Host      string
	Port      string
	DbFile    string
	Resources embed.FS
}

type Link struct {
	gorm.Model
	SiteName    string    `json:"siteName"`
	SiteUrl     string    `json:"siteUrl"`
	FeedUrl     string    `json:"feedUrl"`
	OriginalUrl string    `json:"originalUrl"`
	UnreadCount int16     `json:"unreadCount"`
	LastClicked time.Time `json:"lastClicked"`
	LastFetched time.Time `json:"lastFetched"`
}

func New(conf *Config) *App {

	db, err := gorm.Open(sqlite.Open(conf.DbFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	templates, err := template.ParseFS(conf.Resources, "res/tmpl/*.tmpl")
	if err != nil {
		panic(err)
	}

	fSys, err := fs.Sub(fs.FS(conf.Resources), "res")
	if err != nil {
		panic(err)
	}

	reader := feedreader.New()

	// server := simpleserver.New(&simpleserver.Config{
	// 	Host:           conf.Host,
	// 	Port:           conf.Port,
	// 	Resources:      conf.Resources,
	// 	TemplatesGlob:  "res/tmpl/*.tmpl",
	// 	StaticDirName:  "/static/",
	// 	EmbedFSRootDir: "res",
	// })

	router := gin.Default()
	router.StaticFS("/static", http.FS(fSys))

	app := &App{conf, router, reader, db}

	router.SetHTMLTemplate(templates)

	router.GET("/", app.IndexGet)

	api := router.Group("/api")
	api.GET("/links", app.ApiLinksGet)
	api.POST("/links", app.ApiLinksPost)
	api.GET("/links/{id}", app.ApiLinksIdGet)
	api.PUT("/links/{id}", app.ApiLinksIdPut)
	api.DELETE("/links/{id}", app.ApiLinksIdPut)

	// main page
	// server.Route("/", app.IndexGet, &simpleserver.RouteOptions{})

	// api
	// server.Route("/api/links", app.ApiLinksGet, &simpleserver.RouteOptions{
	// 	Methods: []string{http.MethodGet},
	// })
	// server.Route("/api/links", app.ApiLinksPost, &simpleserver.RouteOptions{
	// 	Methods: []string{http.MethodPost},
	// })
	// server.Route("/api/links/{id}", app.ApiLinksIdGet, &simpleserver.RouteOptions{
	// 	Methods: []string{http.MethodGet},
	// })
	// server.Route("/api/links/{id}", app.ApiLinksIdPut, &simpleserver.RouteOptions{
	// 	Methods: []string{http.MethodPut},
	// })
	// server.Route("/api/links/{id}", app.ApiLinksIdDelete, &simpleserver.RouteOptions{
	// 	Methods: []string{http.MethodDelete},
	// })

	return app
}

func (a *App) Start() error {

	a.db.AutoMigrate(&Link{})

	return a.router.Run()
}
