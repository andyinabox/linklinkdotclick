package app

import "time"

type Link struct {
	// gorm fields
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// domain fields
	FeedFound   bool      `json:"feedFound"`
	SiteName    string    `json:"siteName"`
	SiteUrl     string    `json:"siteUrl"`
	FeedUrl     string    `json:"feedUrl"`
	OriginalUrl string    `json:"originalUrl"`
	UnreadCount int16     `json:"unreadCount"`
	LastClicked time.Time `json:"lastClicked"`
	LastFetched time.Time `json:"lastFetched"`
}

type LinkRepository interface {
	FetchLinks() ([]Link, error)
	CreateLink(Link) (*Link, error)
	FetchLink(id uint) (*Link, error)
	UpdateLink(Link) (*Link, error)
	DeleteLink(id uint) (uint, error)
}

type LinkService interface {
	FetchLinks() ([]Link, error)
	CreateLink(url string) (*Link, error)
	FetchLink(id uint, refresh bool) (*Link, error)
	UpdateLink(id uint, link Link, refresh bool) (*Link, error)
	DeleteLink(id uint) (uint, error)
	RefreshLink(link Link) (*Link, error)
}
