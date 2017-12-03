// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticache

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/internal/protocol/query"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
)

// Amazon ElastiCache is a web service that makes it easier to set up, operate,
// and scale a distributed cache in the cloud.
//
// With ElastiCache, customers gain all of the benefits of a high-performance,
// in-memory cache with far less of the administrative burden of launching and
// managing a distributed cache. The service makes setup, scaling, and cluster
// failure handling much simpler than in a self-managed cache deployment.
//
// In addition, through integration with Amazon CloudWatch, customers get enhanced
// visibility into the key performance statistics associated with their cache
// and can receive alarms if a part of their cache runs hot.
type ElastiCache struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new ElastiCache client.
func New(config *aws.Config) *ElastiCache {
	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config),
		ServiceName: "elasticache",
		APIVersion:  "2015-02-02",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &ElastiCache{service}
}

// newRequest creates a new request for a ElastiCache operation and runs any
// custom request initialization.
func (c *ElastiCache) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}