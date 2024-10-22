package bnfuture

import (
	"context"
	repomodel "tradething/app/bn/bn_future/repository_model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamodbConfig struct{}

func (d *dynamodbConfig) LoadConfig() aws.Config {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return cfg
}
func NewDynamodbConfig() *dynamodbConfig {
	return &dynamodbConfig{}
}

type credential struct {
	AccessKeyID     string
	SecretAccessKey string
}

func (c *credential) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     c.AccessKeyID,
		SecretAccessKey: c.SecretAccessKey,
	}, nil
}

func NewCredential(accessKeyID, secretAccessKey string) *credential {
	return &credential{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
}

type IRepository interface {
	GetAllOpenOrders(ctx context.Context) ([]repomodel.BinanceFutureOpeningPosition, error)
	GetOpenOrderBySymbol(ctx context.Context, symbol string) (*repomodel.BinanceFutureOpeningPosition, error)
	GetOpenOrderByClientID(ctx context.Context, clientId string) (*repomodel.BinanceFutureOpeningPosition, error)
	NewOpenOrder(ctx context.Context, openOrder *repomodel.BinanceFutureOpeningPosition) error
	DeleteOpenOrderBySymbol(ctx context.Context, symbol string) error
}

type dynamoDBRepository struct {
	dynamodb *dynamodb.Client
}

func NewDynamoDBRepository(
	dynamodb *dynamodb.Client,
) IRepository {
	return &dynamoDBRepository{
		dynamodb: dynamodb,
	}
}

type newDynamodb struct {
	_endPoint   dynamodb.EndpointResolverV2
	_credential aws.CredentialsProvider
	_awsconfig  aws.Config
}

func (c newDynamodb) New() *dynamodb.Client {
	svc := dynamodb.NewFromConfig(c._awsconfig, func(o *dynamodb.Options) {
		o.Credentials = c._credential
		o.EndpointResolverV2 = c._endPoint
	})
	return svc
}

func DynamoDB(
	_endPoint dynamodb.EndpointResolverV2,
	_credential aws.CredentialsProvider,
	_awsconfig aws.Config,
) *newDynamodb {
	return &newDynamodb{
		_endPoint,
		_credential,
		_awsconfig,
	}
}
