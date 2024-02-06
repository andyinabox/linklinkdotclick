package linkrepository

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
)

func (r *Repository) UpdateLink(link app.Link) (*app.Link, error) {
	if link.UserID == 0 {
		return nil, errors.New("zero-value user id")
	}
	// disallow upsert
	var testLink app.Link
	tx := r.db.First(&testLink)
	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = r.db.Save(&link)
	return &link, tx.Error
}
