package app

import "time"

type Link struct {
	// gorm fields
	ID        uint      `json:"id"`
	UserID    uint      `json:"-"`
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

type LinkRepository interface {
	FetchLinks(userId uint) ([]Link, error)
	CreateLink(link Link) (*Link, error)
	FetchLink(id uint) (*Link, error)
	UpdateLink(link Link) (*Link, error)
	DeleteLink(id uint) (uint, error)
}

type LinkService interface {
	FetchLinks(userId uint) ([]Link, error)
	CreateLink(userId uint, url string) (*Link, error)
	FetchLink(userId uint, id uint, refresh bool) (*Link, error)
	UpdateLink(userId uint, id uint, link Link, refresh bool) (*Link, error)
	DeleteLink(userId uint, id uint) (uint, error)
	RefreshLink(userId uint, link Link) (*Link, error)
}
