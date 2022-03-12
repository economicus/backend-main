package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
)

const (
	region = "ap-northeast-2"
	bucket = "economicus"
)

type conf struct {
	ACL                *string
	CacheControl       *string
	ContentDisposition *string
	Region             string
	Bucket             string
}

func newConf() *conf {
	return &conf{
		ACL:                aws.String("public-read"),
		CacheControl:       aws.String("max-age=86400"),
		ContentDisposition: aws.String("attachment"),
		Region:             region,
		Bucket:             bucket,
	}
}

func (c *conf) getDomain() string {
	return fmt.Sprintf("%s.s3.%s.amazonaws.com", c.Bucket, c.Region)
}

func (c *conf) getAwsConfig() *aws.Config {
	return &aws.Config{
		Region: aws.String(c.Region),
	}
}

func (c *conf) Info() {
	fmt.Println("========== AWS ==========")
	fmt.Println("ACL: ", *c.ACL)
	fmt.Println("CacheControl", *c.CacheControl)
	fmt.Println("ContentDisposition: ", *c.ContentDisposition)
	fmt.Println("Region", c.Region)
	fmt.Println("Bucket: ", c.Bucket)
}
