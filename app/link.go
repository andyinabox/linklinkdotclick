package app

import "time"

type Link struct {
	// gorm fields
	ID        uint      `json:"id" form:"id"`
	UserID    uint      `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// domain fields
	SiteName        string    `json:"siteName" form:"site-name"`
	SiteUrl         string    `json:"siteUrl" form:"site-url" gorm:"index"`
	FeedUrl         string    `json:"feedUrl" form:"feed-url"`
	OriginalUrl     string    `json:"originalUrl"`
	UnreadCount     int16     `json:"unreadCount" form:"unread-count"`
	LastClicked     time.Time `json:"lastClicked" form:"last-clicked"`
	LastFetched     time.Time `json:"lastFetched"`
	HideUnreadCount bool      `json:"hideUnreadCount" form:"hide-unread-count"`
}

type LinkRepository interface {
	FetchLinks(userId uint) ([]Link, error)
	CreateLink(link Link) (*Link, error)
	FetchLink(userId uint, id uint) (*Link, error)
	UpdateLink(link Link) (*Link, error)
	DeleteLink(userId uint, id uint) (uint, error)
}

type LinkService interface {
	FetchLinks(userId uint) ([]Link, error)
	CreateLink(userId uint, url string) (*Link, error)
	CreateLinkFromFeed(userId uint, siteTitle string, feedUrl string, siteUrl string) (*Link, error)
	RegisterLinkClick(userId uint, id uint, t time.Time) (*Link, error)
	FetchLink(userId uint, id uint, refresh bool) (*Link, error)
	UpdateLink(userId uint, link Link, refresh bool) (*Link, error)
	DeleteLink(userId uint, id uint) (uint, error)
	RefreshAndUpdateLink(userId uint, link Link, force bool) (*Link, error)
}
