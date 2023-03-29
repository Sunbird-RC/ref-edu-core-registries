// Code generated by go-swagger; DO NOT EDIT.

package uploaded_files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetV1BulkUploadedFilesHandlerFunc turns a function with the right signature into a get v1 bulk uploaded files handler
type GetV1BulkUploadedFilesHandlerFunc func(GetV1BulkUploadedFilesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1BulkUploadedFilesHandlerFunc) Handle(params GetV1BulkUploadedFilesParams) middleware.Responder {
	return fn(params)
}

// GetV1BulkUploadedFilesHandler interface for that can handle valid get v1 bulk uploaded files params
type GetV1BulkUploadedFilesHandler interface {
	Handle(GetV1BulkUploadedFilesParams) middleware.Responder
}

// NewGetV1BulkUploadedFiles creates a new http.Handler for the get v1 bulk uploaded files operation
func NewGetV1BulkUploadedFiles(ctx *middleware.Context, handler GetV1BulkUploadedFilesHandler) *GetV1BulkUploadedFiles {
	return &GetV1BulkUploadedFiles{Context: ctx, Handler: handler}
}

/*GetV1BulkUploadedFiles swagger:route GET /v1/bulk/uploadedFiles uploadedFiles getV1BulkUploadedFiles

get uploaded files

*/
type GetV1BulkUploadedFiles struct {
	Context *middleware.Context
	Handler GetV1BulkUploadedFilesHandler
}

func (o *GetV1BulkUploadedFiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1BulkUploadedFilesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}