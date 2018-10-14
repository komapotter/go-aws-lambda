package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	awsctr "github.com/komapotter/go-awsctr"
)

type awsSvc struct {
	cw awsctr.CloudWatch
}

func newAWSSvc() awsSvc {
	sess := awsctr.NewSession("ap-southeast-1")
	return awsSvc{
		cw: awsctr.NewCloudWatch(sess),
	}
}

// HandlerRequest -
func HandlerRequest() error {
	alarm := os.Getenv("CW_ALARM_NAME")
	flag := os.Getenv("FLAG")

	awssvc := newAWSSvc()

	switch flag {
	case "on":
		err := awssvc.cw.AlarmOn(awsctr.AlarmInfo{
			Name: alarm,
		})
		if err != nil {
			return err
		}
	case "off":
		err := awssvc.cw.AlarmOff(awsctr.AlarmInfo{
			Name: alarm,
		})
		if err != nil {
			return err
		}
	default:
		log.Fatal(fmt.Sprintf("invalid flag: %s", flag))
	}
	return nil
}

func main() {
	lambda.Start(HandlerRequest)
}
