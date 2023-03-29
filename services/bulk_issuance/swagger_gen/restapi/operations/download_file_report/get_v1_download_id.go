// Code generated by go-swagger; DO NOT EDIT.

package download_file_report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetV1DownloadIDHandlerFunc turns a function with the right signature into a get v1 download ID handler
type GetV1DownloadIDHandlerFunc func(GetV1DownloadIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1DownloadIDHandlerFunc) Handle(params GetV1DownloadIDParams) middleware.Responder {
	return fn(params)
}

// GetV1DownloadIDHandler interface for that can handle valid get v1 download ID params
type GetV1DownloadIDHandler interface {
	Handle(GetV1DownloadIDParams) middleware.Responder
}

// NewGetV1DownloadID creates a new http.Handler for the get v1 download ID operation
func NewGetV1DownloadID(ctx *middleware.Context, handler GetV1DownloadIDHandler) *GetV1DownloadID {
	return &GetV1DownloadID{Context: ctx, Handler: handler}
}

/*GetV1DownloadID swagger:route GET /v1/download/{id} downloadFileReport getV1DownloadId

download the success and error report of file uploaded

*/
type GetV1DownloadID struct {
	Context *middleware.Context
	Handler GetV1DownloadIDHandler
}

func (o *GetV1DownloadID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1DownloadIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}