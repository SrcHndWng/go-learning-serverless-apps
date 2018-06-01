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
	tbl := table()
	err = tbl.Get("ID", photoID).One(&item)
	return
}

// Count gets count by id.
func Count(photoID string) (count int64, err error) {
	tbl := table()
	count, err = tbl.Get("ID", photoID).Count()
	return
}

// UpdateItem updates data.
func UpdateItem(photoID string, timestamp int64, status string) error {
	tbl := table()
	return tbl.Update("ID", photoID).Set("Timestamp", timestamp).Set("Status", status).Run()
}

// GetItemsByStatus gets items by status.
func GetItemsByStatus(status string) (items []Item, err error) {
	tbl := table()
	err = tbl.Scan().Filter("'Status' = ?", status).All(&items)
	return
}

// DeleteItem removes item.
func DeleteItem(photoID string) error {
	tbl := table()
	return tbl.Delete("ID", photoID).Run()
}

func table() dynamo.Table {
	return Table(os.Getenv("TABLE_NAME"))
}
