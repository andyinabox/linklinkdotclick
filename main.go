package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/apirouter"
	"github.com/andyinabox/linkydink/app/approuter"
	"github.com/andyinabox/linkydink/app/authhelper"
	"github.com/andyinabox/linkydink/app/htmlresponsehelper"
	"github.com/andyinabox/linkydink/app/jsonresponsehelper"
	"github.com/andyinabox/linkydink/app/linkrepository"
	"github.com/andyinabox/linkydink/app/linkservice"
	"github.com/andyinabox/linkydink/app/servicecontainer"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/app/userservice"
	"github.com/andyinabox/linkydink/pkg/logservice"
	"github.com/andyinabox/linkydink/pkg/tokenstore"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

//go:embed res/*
var res embed.FS

//go:embed VERSION
var version string

const templatesGlob = "res/tmpl/*.tmpl"

func init() {
	godotenv.Load()
	version = strings.TrimSpace(version)
}

// registerConfigVar registers both env vars and command-line flags,
// falling back to a default value if neighter are found
// command-line flags are given priority, env variables second, default value last
// env vars should be in the format `LINKY_<NAME>` with `<NAME>` being an uppercase
// version of the CLI flag
func registerConfigVar(variable *string, name string, def string, description string) {
	defaultValue := os.Getenv("LINKY_" + strings.ToUpper(name))
	if defaultValue == "" {
		defaultValue = def
	}
	flag.StringVar(variable, name, defaultValue, description)
}

func main() {

	var domain string
	var port string
	var dbfile string
	var mode string
	var defaultemail string
	var defaultusertitle string
	var smtpaddr string
	var secret string

	registerConfigVar(&domain, "domain", "linklink.click", "the domain the site is hosted on (linklink.click)")
	registerConfigVar(&port, "port", "8080", "port to run the webserver on")
	registerConfigVar(&dbfile, "dbfile", "db/linkydink.db", "location on sqlite db")
	registerConfigVar(&mode, "mode", "debug", "run mode, use 'release' for production")
	registerConfigVar(&defaultemail, "defaultemail", "linkydink@linkydink.tld", "an email for the default user that appears when not logged in")
	registerConfigVar(&defaultusertitle, "defaultusertitle", "📚 my reading list", "the default user's site title")
	registerConfigVar(&smtpaddr, "smtpaddr", "127.0.0.1:1025", "smtp server")
	registerConfigVar(&secret, "secret", "", "secret to use for cookie encryption")
	flag.Parse()

	if secret == "" {
		secret = domain + port + dbfile + mode
	}

	fmt.Printf(`
	
                🖇✨ linkydink starting ✨🖇
								
                         %s

	                Port: %s
	                Mode: %s
	              DbFile: %s
	              Domain: %s
	            SmtpAddr: %s
	    DefaultUserEmail: %s
	DefaultUSerSiteTitle: %s

	`, version, port, mode, dbfile, domain, smtpaddr, defaultemail, defaultusertitle)

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
	templates, err := template.ParseFS(res, templatesGlob)
	if err != nil {
		panic(err)
	}

	// create session store
	sessionStore := cookie.NewStore([]byte(secret))

	// crewate log service
	logService := logservice.New()

	// create user repository
	userRepository := userrepository.New(db)

	// create user service
	tokenStore := tokenstore.New(db, &tokenstore.Config{
		ExpireseIn: 10 * time.Minute,
	})
	userServiceConfig := &userservice.Config{
		DefaultUserEmail:     defaultemail,
		DefaultUserSiteTitle: defaultusertitle,
	}
	userService := userservice.New(userRepository, tokenStore, userServiceConfig)

	// create link service
	linkRepository := linkrepository.New(db)
	linkService := linkservice.New(
		linkRepository,
		logService,
		&linkservice.Config{
			LinkRefreshBuffer: 10 * time.Minute,
		},
	)

	// create service container
	serviceContainer := servicecontainer.New(
		userService,
		linkService,
		logService,
	)

	authHelper := authhelper.New(serviceContainer, &authhelper.Config{
		SessionUserKey: "user",
	})

	jsonResponseHelper := jsonresponsehelper.New()

	title := "link link dot click"
	htmlResponseHelper := htmlresponsehelper.New(&htmlresponsehelper.Config{
		SiteTitle:         title,
		Description:       "Somewhere in-between a blogroll and an RSS reader",
		FavIconUrl:        "/static/favicon.ico",
		AppleTouchIconUrl: "/static/apple-touch-icon.png",
		ManifestUrl:       "/static/site.webmanifest",
		OgImageUrl:        "/static/android-chrome-512x512.png",
		OgImageAlt:        "Two paperclips entwined",
		InfoPageSuccessOptions: &app.HtmlInfoMessageOptions{
			LinkText: "Back to the main page",
			LinkUrl:  "/",
		},
		InfoPageErrorOptions: &app.HtmlInfoMessageOptions{
			Message:  "🫠 Uh-oh, something went wrong...",
			LinkText: "Back to safety",
			LinkUrl:  "/",
		},
	})

	// create routers
	appRouter := approuter.New(
		serviceContainer,
		authHelper,
		htmlResponseHelper,
		&approuter.Config{
			Templates: templates,
			Version:   version,
			SmtpAddr:  smtpaddr,
		},
	)

	apiRouter := apirouter.New(
		serviceContainer,
		authHelper,
		jsonResponseHelper,
		&apirouter.Config{
			Domain: domain,
			Mode:   mode,
		},
	)
	routers := []app.RouterGroup{appRouter, apiRouter}

	// create app
	appConfig := &app.Config{
		Domain:      domain,
		Port:        port,
		Mode:        mode,
		Resources:   res,
		Templates:   templates,
		Version:     version,
		SessionName: "session",
	}
	appInstance := app.New(
		sessionStore,
		routers,
		appConfig,
	)

	log.Fatal(appInstance.Start())
}
