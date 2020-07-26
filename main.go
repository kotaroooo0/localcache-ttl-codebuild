package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

func main() {
	svc := codebuild.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	input := &codebuild.StartBuildInput{ProjectName: aws.String("localcache-ttl-test")}

	// TODO: 二分探索で調べる
	output, err := svc.StartBuild(input)
	if err != nil {
		fmt.Println("Got error building project: ", err)
		os.Exit(1)
	}

	fmt.Println("Started build")
	fmt.Println(output)
}
