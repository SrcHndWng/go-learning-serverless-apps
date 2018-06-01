package models

import (
	"fmt"
	"os"
	"testing"
)

func TestGetItem(t *testing.T) {
	before()
	item, err := GetItem("e618fa31-eb30-4b8e-9a43-a8f7a3bd7272")
	if err != nil {
		t.Fatalf("raise error.\n%v", err)
	}
	fmt.Printf("item = \n%v\n", item)
}

func TestGetItemsByStatus(t *testing.T) {
	before()
	items, err := GetItemsByStatus("Uploaded")
	if err != nil {
		t.Fatalf("raise error.\n%v", err)
	}
	fmt.Printf("items count = %v\n", len(items))
	fmt.Printf("items = \n%v\n", items)
}

func before() {
	os.Setenv("TABLE_NAME", "picture_post_photos")
}
