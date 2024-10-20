package bnfuture

import (
	"net/url"

	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

type endPointResolver struct {
	Region   string
	EndPoint string
}

func (e *endPointResolver) ResolveEndpoint() (*smithyendpoints.Endpoint, error) {
	uri, err := url.Parse(e.EndPoint)
	if err != nil {
		return nil, err
	}
	return &smithyendpoints.Endpoint{
		URI: *uri,
	}, nil
}

func NewEndPointResolver(region, endPoint string) *endPointResolver {
	return &endPointResolver{
		Region:   region,
		EndPoint: endPoint,
	}
}
