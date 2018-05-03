package main

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

const region = "ap-northeast-1"
const successMessage = "cloudwatch alarm function success."

type Response struct {
	Message string `json:"message"`
}

func readMessage(message string) (alarmName string, streamName string, err error) {
	var msg map[string]interface{}
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		return "", "", err
	}
	alarmName = msg["AlarmName"].(string)
	streamName = msg["Trigger"].(map[string]interface{})["Dimensions"].([]interface{})[0].(map[string]interface{})["value"].(string)
	return
}

func reshard(alarmName string, streamName string) error {
	s := session.New(&aws.Config{Region: aws.String(region)})
	k := kinesis.New(s)
	w := cloudwatch.New(s)

	summaryOut, err := k.DescribeStreamSummary(&kinesis.DescribeStreamSummaryInput{
		StreamName: aws.String(streamName),
	})
	if err != nil {
		return err
	}

	currentCnt := aws.Int64Value(summaryOut.StreamDescriptionSummary.OpenShardCount)
	targetCnt := currentCnt * 2
	_, err = k.UpdateShardCount(&kinesis.UpdateShardCountInput{
		StreamName:       aws.String(streamName),
		TargetShardCount: aws.Int64(targetCnt),
		ScalingType:      aws.String("UNIFORM_SCALING"),
	})
	if err != nil {
		return err
	}

	threshold := float64(targetCnt*1000) * 0.8
	_, err = w.PutMetricAlarm(&cloudwatch.PutMetricAlarmInput{
		AlarmName:          aws.String(alarmName),
		MetricName:         aws.String("IncomingRecords"),
		Namespace:          aws.String("AWS/Kinesis"),
		Period:             aws.Int64(60),
		EvaluationPeriods:  aws.Int64(1),
		ComparisonOperator: aws.String(cloudwatch.ComparisonOperatorGreaterThanThreshold),
		Threshold:          aws.Float64(threshold),
		Statistic:          aws.String(cloudwatch.StatisticSum),
	})
	if err != nil {
		return err
	}

	return nil
}

func Handler(request events.SNSEvent) (Response, error) {
	alarmName, streamName, err := readMessage(request.Records[0].SNS.Message)
	if err != nil {
		return Response{Message: ""}, err
	}

	if alarmName != "cloudwatch-alarm-sample-kinesis-mon" {
		return Response{Message: ""}, errors.New("alarmName is not valid")
	}

	err = reshard(alarmName, streamName)
	if err != nil {
		return Response{Message: ""}, err
	}

	return Response{Message: successMessage}, nil
}

func main() {
	lambda.Start(Handler)
}
