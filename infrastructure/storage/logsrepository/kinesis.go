package logsrepository

import (
	"dacast-log-agent/core/domain"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/mmatagrin/ctxerror"
	"github.com/pkg/errors"
	"os"
)

var (
	awsSession *session.Session = nil
)

func NewKinesisClient() *kinesisClient {
	return &kinesisClient{}
}

func createOrGetSession() (*session.Session, error) {
	if awsSession == nil {
		sess, errSession := session.NewSession(
			&aws.Config{
				Region: aws.String(os.Getenv("AWS_REGION")),
			},
		)
		if errSession != nil {
			return nil, errors.Wrap(errSession, "Error while creating AWS session")
		}
		awsSession = sess
	}

	return awsSession, nil
}

type kinesisClient struct {
}

func (repo *kinesisClient) Save(log domain.Log) error {
	ctxErr := ctxerror.SetContext(map[string]interface{}{
		"log":     log,
	})

	bytes, err := json.Marshal(log)
	if err != nil {
		return ctxErr.Wrap(err, "log fails at marshal into json string")
	}

	sess, err := createOrGetSession()
	if err != nil {
		return err
	}

	kinesisClient := kinesis.New(sess)

	_, err = kinesisClient.PutRecord(&kinesis.PutRecordInput{
		Data:         bytes,
		StreamName:   aws.String(os.Getenv("STREAM_NAME")),
		PartitionKey: aws.String(log.Id),
	})

	if nil != err {
		return ctxErr.Wrap(err, "Error while enqueuing message")
	}

	return nil
}
