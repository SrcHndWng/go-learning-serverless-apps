package models

import (
	"os"

	"github.com/guregu/dynamo"
)

// Item contains data to register to dynamoDB.
type Item struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Size      int    `json:"size"`
	SignedURL string `json:"signed_url"`
}

// SaveItem inserts data to dynamoDB.
func SaveItem(item Item) error {
	tbl := table()
	return tbl.Put(item).Run()
}

// GetItem gets table data and return.
func GetItem(photoID string) (item Item, err error) {
	tbl := Table("picture_post_photos")
	err = tbl.Get("ID", photoID).One(&item)
	return
}

func table() dynamo.Table {
	return Table(os.Getenv("TABLE_NAME"))
}
