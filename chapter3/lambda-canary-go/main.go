package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func validate(site string, expected string) (bool, error) {
	resp, err := http.Get(site)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	return strings.Contains(string(byteArray), expected), nil
}

func nowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Handler() (Response, error) {
	site := os.Getenv("site")
	expected := os.Getenv("expected")

	fmt.Printf("Checking %s at %s...\n", site, nowDateTime())

	result, err := validate(site, expected)
	if err != nil {
		panic(err)
	}

	var msg string
	if result {
		msg = "Check passed!"
	} else {
		msg = "Check failed!"
	}
	fmt.Println(msg)

	fmt.Printf("Check complete at %s\n", nowDateTime())

	return Response{
		Message: msg,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
