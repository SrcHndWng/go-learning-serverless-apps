package utils

import (
	"fmt"
	"testing"
)

func TestGenerateID(t *testing.T) {
	result := GenerateID()
	if result == "" {
		t.Fatalf("result error.\n")
	}
	fmt.Printf("id = %s\n", result)
}
