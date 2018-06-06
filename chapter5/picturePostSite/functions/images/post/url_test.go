package main

import (
	"fmt"
	"testing"
)

func TestGetPresignedURL(t *testing.T) {
	result, err := getPresignedURL("picture-post-photos", "f940d05d-ebe4-4ab6-87f5-bcb83afc03a3", "image/jpeg")
	if err != nil {
		t.Fatalf("raise error.\n")
	}
	if result == "" {
		t.Fatalf("result error.\n")
	}
	fmt.Printf("url = %s\n", result)
}
