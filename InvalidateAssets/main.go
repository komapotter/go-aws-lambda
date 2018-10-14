package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awsctr "github.com/komapotter/go-awsctr"
)

type awsSvc struct {
	cf awsctr.CloudFront
	cp awsctr.CodePipeline
}

func newAWSSvc() awsSvc {
	sess := awsctr.NewSession("ap-southeast-1")
	return awsSvc{
		cf: awsctr.NewCloudFront(sess),
		cp: awsctr.NewCodePipeline(sess),
	}
}

// HandlerRequest -
func HandlerRequest(codePipelineEvent events.CodePipelineEvent) error {
	path := os.Getenv("OBJECT_PATH")

	awssvc := newAWSSvc()

	var jobID string
	if jobID = codePipelineEvent.CodePipelineJob.ID; jobID == "" {
		return errors.New("job-ID is missing")
	}
	fmt.Printf("job-ID: %s\n", jobID)

	var distID string
	if distID = codePipelineEvent.CodePipelineJob.Data.ActionConfiguration.Configuration.UserParameters; distID == "" {
		return errors.New("distribution-ID is missing")
	}

	err := awssvc.cf.Invalidate(awsctr.InvalidateInfo{
		DistID: distID,
		Path:   path,
	})
	if err != nil {
		if serr := awssvc.cp.SendJobFailure(awsctr.JobInfo{ID: jobID}); err != nil {
			return serr
		}
		return err
	}

	if err := awssvc.cp.SendJobSuccess(awsctr.JobInfo{ID: jobID}); err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(HandlerRequest)
}
