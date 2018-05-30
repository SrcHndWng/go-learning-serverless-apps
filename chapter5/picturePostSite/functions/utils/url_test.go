package utils

import (
	"fmt"
	"testing"
)

func TestGetPresignedURL(t *testing.T) {
	result, err := GetPresignedURL("picture-post-photos", "f940d05d-ebe4-4ab6-87f5-bcb83afc03a3", "aaa/bbb")
	if err != nil {
		t.Fatalf("raise error.\n")
	}
	if result == "" {
		t.Fatalf("result error.\n")
	}
	fmt.Printf("url = %s\n", result)
}
