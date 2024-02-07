package main

import (
	"embed"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path"
	"time"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/apirouter"
	"github.com/andyinabox/linkydink/app/approuter"
	"github.com/andyinabox/linkydink/app/feedservice"
	"github.com/andyinabox/linkydink/app/handlerhelper"
	"github.com/andyinabox/linkydink/app/linkrepository"
	"github.com/andyinabox/linkydink/app/linkservice"
	"github.com/andyinabox/linkydink/app/servicecontainer"
	"github.com/andyinabox/linkydink/app/tokenstore"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/app/userservice"
	"github.com/andyinabox/linkydink/pkg/mailservice"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

//go:embed res/*
var res embed.FS

func main() {
	var domain string
	var port string
	var dbfile string
	var mode string
	var defaultemail string
	var smtpaddr string

	flag.StringVar(&domain, "domain", "linklink.click", "the domain the site is hosted on (linklink.click)")
	flag.StringVar(&port, "port", "8080", "port to run the webserver on")
	flag.StringVar(&dbfile, "dbfile", "db/linkydink.db", "location on sqlite db")
	flag.StringVar(&mode, "mode", "debug", "run mode, use 'release' for production")
	flag.StringVar(&defaultemail, "defaultemail", "linkydink@linkydink.tld", "an email for the default user that appears when not logged in")
	flag.StringVar(&smtpaddr, "smtpaddr", "127.0.0.1:1025", "smtp server")
	flag.Parse()

	// setup users db
	userDbPath := path.Join(path.Dir(dbfile), "usr")
	err := os.MkdirAll(path.Dir(userDbPath), fs.ModePerm)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// load templates
	templates, err := template.ParseFS(res, "res/tmpl/*.tmpl")
	if err != nil {
		panic(err)
	}

	// create session store
	sessionStore := cookie.NewStore([]byte(domain + port + dbfile + mode + defaultemail))

	// create user repository
	userRepository := userrepository.New(db)

	// create user service
	tokenStore := tokenstore.New(db, &tokenstore.Config{
		ExpireseIn: 10 * time.Minute,
	})
	userServiceConfig := &userservice.Config{
		DefaultUserEmail:     defaultemail,
		DefaultUserSiteTitle: "🖇 linklink.click",
	}
	userService := userservice.New(userRepository, tokenStore, userServiceConfig)

	// create mail service
	mailService := mailservice.New(&mailservice.Config{
		SmtpAddr: smtpaddr,
	})

	// create link service
	feedService := feedservice.New()
	linkRepository := linkrepository.New(db)
	linkService := linkservice.New(linkRepository, feedService)

	// create service container
	serviceContainer := servicecontainer.New(
		userService,
		linkService,
		mailService,
	)

	// create handler helper
	handlerHelper := handlerhelper.New(serviceContainer)

	// create routers
	appRouter := approuter.New(serviceContainer, handlerHelper, &approuter.Config{
		Templates: templates,
	})
	apiRouter := apirouter.New(serviceContainer, handlerHelper, &apirouter.Config{
		Domain: domain,
		Mode:   mode,
	})
	routers := []app.RouterGroup{appRouter, apiRouter}

	// create app
	appConfig := &app.Config{
		Domain:    domain,
		Port:      port,
		Mode:      mode,
		Resources: res,
		Templates: templates,
	}
	appInstance := app.New(
		sessionStore,
		routers,
		appConfig,
	)

	log.Fatal(appInstance.Start())
}
