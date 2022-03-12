package aws

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"mime/multipart"
	"net/http"
)

type Aws struct {
	*conf
	Uploader *s3manager.Uploader
}

func New() *Aws {
	a := Aws{}
	a.conf = newConf()
	sess, err := session.NewSession(a.conf.getAwsConfig())
	if err != nil {
		log.Panicf("error while connecting aws: %v", err)
	}
	uploader := s3manager.NewUploader(sess)
	return &Aws{
		Uploader: uploader,
	}
}

func (a *Aws) UploadFile(file multipart.File, header *multipart.FileHeader) (*s3manager.UploadOutput, error) {
	filename := fmt.Sprintf("photos/%s", header.Filename)
	buff := make([]byte, header.Size)
	_, err := file.Read(buff)
	if err != nil {
		return nil, err
	}
	up, err := a.Uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(a.conf.Bucket),
		ACL:                a.conf.ACL,
		CacheControl:       a.conf.CacheControl,
		ContentType:        aws.String(http.DetectContentType(buff)),
		Key:                aws.String(filename),
		Body:               bytes.NewBuffer(buff),
		ContentDisposition: a.conf.ContentDisposition,
	})
	if err != nil {
		return nil, fmt.Errorf("error in UploadFile while uploading file: %w", err)
	}
	return up, nil
}

func (a *Aws) GetFilePath(filename string) string {
	return fmt.Sprintf("https://%s/%s", a.conf.getDomain(), filename)
}
