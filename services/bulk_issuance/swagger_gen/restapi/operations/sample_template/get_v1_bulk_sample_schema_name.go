// Code generated by go-swagger; DO NOT EDIT.

package sample_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetV1BulkSampleSchemaNameHandlerFunc turns a function with the right signature into a get v1 bulk sample schema name handler
type GetV1BulkSampleSchemaNameHandlerFunc func(GetV1BulkSampleSchemaNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1BulkSampleSchemaNameHandlerFunc) Handle(params GetV1BulkSampleSchemaNameParams) middleware.Responder {
	return fn(params)
}

// GetV1BulkSampleSchemaNameHandler interface for that can handle valid get v1 bulk sample schema name params
type GetV1BulkSampleSchemaNameHandler interface {
	Handle(GetV1BulkSampleSchemaNameParams) middleware.Responder
}

// NewGetV1BulkSampleSchemaName creates a new http.Handler for the get v1 bulk sample schema name operation
func NewGetV1BulkSampleSchemaName(ctx *middleware.Context, handler GetV1BulkSampleSchemaNameHandler) *GetV1BulkSampleSchemaName {
	return &GetV1BulkSampleSchemaName{Context: ctx, Handler: handler}
}

/*GetV1BulkSampleSchemaName swagger:route GET /v1/bulk/sample/{schemaName} sampleTemplate getV1BulkSampleSchemaName

get sample template

*/
type GetV1BulkSampleSchemaName struct {
	Context *middleware.Context
	Handler GetV1BulkSampleSchemaNameHandler
}

func (o *GetV1BulkSampleSchemaName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1BulkSampleSchemaNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}