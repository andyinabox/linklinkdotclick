package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
<<<<<<< HEAD
	"os"
	"path"
	"time"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/apirouter"
	"github.com/andyinabox/linkydink/app/approuter"
	"github.com/andyinabox/linkydink/app/handlerhelper"
	"github.com/andyinabox/linkydink/app/servicecontainer"
	"github.com/andyinabox/linkydink/app/tokenstore"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/app/userservice"
	"github.com/andyinabox/linkydink/pkg/mailservice"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
=======
	"path"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/app/userservice"
	"github.com/gin-contrib/sessions/cookie"
>>>>>>> main
)

//go:embed res/*
var res embed.FS

func main() {
	var domain string
	var port string
	var dbfile string
	var mode string
	var defaultemail string
<<<<<<< HEAD
	var smtpaddr string
=======
>>>>>>> main

	flag.StringVar(&domain, "domain", "linklink.click", "the domain the site is hosted on (linklink.click)")
	flag.StringVar(&port, "port", "8080", "port to run the webserver on")
	flag.StringVar(&dbfile, "dbfile", "db/linkydink.db", "location on sqlite db")
	flag.StringVar(&mode, "mode", "debug", "run mode, use 'release' for production")
	flag.StringVar(&defaultemail, "defaultemail", "linkydink@linkydink.tld", "an email for the default user that appears when not logged in")
<<<<<<< HEAD
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

	// create session store
	sessionStore := cookie.NewStore([]byte(domain + port + dbfile + mode + defaultemail))

	// create user repository
	userRepository := userrepository.New(db)

	// create user service
	tokenStore := tokenstore.New(db, &tokenstore.Config{
		ExpireseIn: 10 * time.Minute,
	})
	userServiceConfig := &userservice.Config{
		UserDbPath:       userDbPath,
		DefaultUserEmail: defaultemail,
	}
	userService := userservice.New(userRepository, tokenStore, userServiceConfig)

	// create mail service
	mailService := mailservice.New(&mailservice.Config{
		SmtpAddr: smtpaddr,
	})

	// get default linkService
	user, err := userService.EnsureDefaultUser()
	if err != nil {
		panic(err)
	}
	linkService, err := userService.GetUserLinkService(user)
	if err != nil {
		panic(err)
	}

	// create service container
	serviceContainer := servicecontainer.New(
		userService,
		linkService,
		mailService,
	)

	// create handler helper
	handlerHelper := handlerhelper.New(serviceContainer)

	// create routers
	appRouter := approuter.New(serviceContainer, handlerHelper)
	apiRouter := apirouter.New(serviceContainer, handlerHelper, &apirouter.Config{
		Domain: domain,
		Mode:   mode,
	})
	routers := []app.RouterGroup{appRouter, apiRouter}
=======
	flag.Parse()

	// maybe not the most secure?
	store := cookie.NewStore([]byte(domain + port + dbfile + mode + defaultemail))

	userRepository, err := userrepository.New(&userrepository.Config{
		DbFile: dbfile,
	})
	if err != nil {
		panic(err)
	}

	userService := userservice.New(&userservice.Config{
		UserDbPath:       path.Join(path.Dir(dbfile), "usr"),
		DefaultUserEmail: defaultemail,
	}, userRepository)
>>>>>>> main

	// create app
	appConfig := &app.Config{
		Domain:    domain,
		Port:      port,
		Mode:      mode,
		Resources: res,
<<<<<<< HEAD
	}
	appInstance := app.New(
		sessionStore,
		routers,
		appConfig,
	)
=======
	}, userService, store)
>>>>>>> main

	log.Fatal(appInstance.Start())
}
