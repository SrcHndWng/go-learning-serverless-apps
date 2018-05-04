package main

import (
	"testing"
)

func TestValiate(t *testing.T) {
	const site = "https://www.google.co.jp/"
	const contained = "google"
	const uncontained = "yahoo"

	test := func(expected string, result bool) {
		r, err := validate(site, expected)
		if err != nil {
			t.Error("validate error.")
		}
		if r != result {
			t.Errorf("%s is not found.", expected)
		}
	}

	test(contained, true)
	test(uncontained, false)
}
