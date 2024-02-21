package htmlresponsehelper

import "github.com/andyinabox/linkydink/app"

type page string

const (
	about       page = "about"
	home        page = "home"
	info        page = "info"
	styleEditor page = "styleEditor"
)

type Config struct {
	SiteTitle              string
	Description            string
	FavIconUrl             string
	AppleTouchIconUrl      string
	ManifestUrl            string
	OgImageUrl             string
	OgImageAlt             string
	AppVersion             string
	InfoPageSuccessOptions *app.HtmlInfoMessageOptions
	InfoPageErrorOptions   *app.HtmlInfoMessageOptions
}

type Helper struct {
	sc   app.ServiceContainer
	ah   app.AuthHelper
	conf *Config
}

func New(sc app.ServiceContainer, ah app.AuthHelper, conf *Config) *Helper {
	return &Helper{sc, ah, conf}
}
