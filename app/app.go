package app

import (
	"embed"
	"fmt"
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
	Mode      string
	DbFile    string
	Resources embed.FS
}

type Link struct {
	// gorm fields
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// domain fields
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

	staticFiles, err := fs.Sub(fs.FS(conf.Resources), "res/static")
	if err != nil {
		panic(err)
	}

	reader := feedreader.New()

	gin.SetMode(conf.Mode)
	router := gin.Default()
	router.SetHTMLTemplate(templates)
	router.StaticFS("/static", http.FS(staticFiles))

	app := &App{conf, router, reader, db}

	router.GET("/", app.IndexGet)

	api := router.Group("/api")
	api.GET("/links", app.ApiLinksGet)
	api.POST("/links", app.ApiLinksPost)
	api.GET("/links/:id", app.ApiLinksIdGet)
	api.PUT("/links/:id", app.ApiLinksIdPut)
	api.DELETE("/links/:id", app.ApiLinksIdDelete)

	return app
}

func (a *App) Start() error {

	a.db.AutoMigrate(&Link{})

	return a.router.Run(fmt.Sprintf("%s:%s", a.conf.Host, a.conf.Port))
}
