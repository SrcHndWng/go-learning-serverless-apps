package main

import "testing"

func TestRegister(t *testing.T) {
	data := new(KinesisData)
	data.ID = 1
	data.Text = "test text."
	data.CreatedAt = "Mon May 14 13:22:58 +0000 2018"

	err := register(data)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}
