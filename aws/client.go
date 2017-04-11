package aws

import (
	"fmt"
	"io"
	//"log"

	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ClientAWS struct {
	session *s3.S3
}

func NewClientAWS() *ClientAWS {
	creds, err := getCredentials()
	if err != nil {
		//log.Panicf("Ocorreu um erro ao tentar pegar as credenciais na AWS. %v", err)
	}
	return &ClientAWS{newSessionAWS(creds)}
}

func getCredentials() (*credentials.Credentials, error) {
	token := ""
	creds := credentials.NewStaticCredentials(config.Conf.Credentials.Id, config.Conf.Credentials.SecretKey, token)
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}
	return creds, nil
}

func newSessionAWS(creds *credentials.Credentials) *s3.S3 {
	cfg := aws.NewConfig().WithRegion(config.Conf.Bucket.Region).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)
	return svc
}
func (c *ClientAWS) SendToS3(params *s3.PutObjectInput) {
	_, err := c.session.PutObject(params)
	if err != nil {
		fmt.Printf("bad response: %s", err)
	}
	//fmt.Printf("response %s", awsutil.StringValue(resp))
}

func NewParams(bucketName string, path string, sizeFile int64, fileBytes io.ReadSeeker, fileType string) *s3.PutObjectInput {
	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(sizeFile),
		ContentType:   aws.String(fileType),
	}
	return params
}
