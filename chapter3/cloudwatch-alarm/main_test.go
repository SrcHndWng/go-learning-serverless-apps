package main

import (
	"io/ioutil"
	"testing"
)

const BUFSIZE = 1024

func TestReadMessage(t *testing.T) {
	data, err := ioutil.ReadFile("message.json")
	if err != nil {
		t.Error("message.json read error.")
	}
	message := string(data)
	alarmName, streamName, err := readMessage(message)
	if err != nil {
		t.Error(err)
	}
	if alarmName != "cloudwatch-alarm-sample-kinesis-mon" {
		t.Error("read alarmName error.")
	}
	if streamName != "cloudwatch-alarm-sample-stream" {
		t.Error("read streamName error.")
	}
}
func TestReshard(t *testing.T) {
	err := reshard("cloudwatch-alarm-sample-kinesis-mon", "cloudwatch-alarm-sample-stream")
	if err != nil {
		t.Error(err)
	}
}
