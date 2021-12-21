// Code generated by goa v3.5.3, DO NOT EDIT.
//
// customers endpoints
//
// Command:
// $ goa gen github.com/jolinGalal/jumia/internal/customer/design

package customers

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "customers" service endpoints.
type Endpoints struct {
	List goa.Endpoint
}

// NewEndpoints wraps the methods of the "customers" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List: NewListEndpoint(s),
	}
}

// Use applies the given middleware to all the "customers" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "customers".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		return s.List(ctx, p)
	}
}
