package kinesis_repository

import (
	"dacast-log-agent/config"
	"dacast-log-agent/lib/core/domain"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/mmatagrin/ctxerror"
	"github.com/pkg/errors"
	"sync"
)

var (
	awsSession *session.Session = nil
	mutex = &sync.Mutex{}
    client *kinesis.Kinesis
	streamName string
)

func NewKinesisClient(config config.ClientConfig) *kinesisClient {
	streamName = config.Config["streamName"]
	return &kinesisClient{}
}

func GetClient() (*kinesis.Kinesis, error) {
	if client == nil{
		mutex.Lock()
		defer mutex.Unlock()

		sess, err := createOrGetSession()
		if err != nil {
			return nil, err
		}

		client = kinesis.New(sess)
	}

	return client, nil
}

func createOrGetSession() (*session.Session, error) {
	if awsSession == nil {
		sess, errSession := session.NewSession()
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

	kinesisClient, err := GetClient()
	if nil != err {
		return err
	}

	_, err = kinesisClient.PutRecord(&kinesis.PutRecordInput{
		Data:         bytes,
		StreamName:   aws.String(streamName),
		PartitionKey: aws.String(log.Id),
	})

	if nil != err {
		return ctxErr.Wrap(err, "Error while enqueuing message")
	}

	return nil
}
