package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/fswebcoder/awsgo/constants"
	"github.com/fswebcoder/shared/constant"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func StarAws() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion(constants.US_EST_1))

	if err != nil {
		panic(constant.ERROR_AWS_SDK_CONFIG + err.Error())
	}

}
