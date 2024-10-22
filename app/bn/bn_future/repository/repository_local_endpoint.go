package bnfuture

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

type endPointResolver struct {
	Region   string
	EndPoint string
}

func (e *endPointResolver) ResolveEndpoint(ctx context.Context, params dynamodb.EndpointParameters) (smithyendpoints.Endpoint, error) {
	uri, err := url.Parse(e.EndPoint)
	if err != nil {
		return smithyendpoints.Endpoint{}, err
	}
	return smithyendpoints.Endpoint{
		URI: *uri,
	}, nil
}

func NewEndPointResolver(region, endPoint string) dynamodb.EndpointResolverV2 {
	return &endPointResolver{
		Region:   region,
		EndPoint: endPoint,
	}
}
