// Code generated by go-swagger; DO NOT EDIT.

package sample_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetV1BulkSampleSchemaNameParams creates a new GetV1BulkSampleSchemaNameParams object
// no default values defined in spec.
func NewGetV1BulkSampleSchemaNameParams() GetV1BulkSampleSchemaNameParams {

	return GetV1BulkSampleSchemaNameParams{}
}

// GetV1BulkSampleSchemaNameParams contains all the bound params for the get v1 bulk sample schema name operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetV1BulkSampleSchemaName
type GetV1BulkSampleSchemaNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*schema name
	  Required: true
	  In: path
	*/
	SchemaName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetV1BulkSampleSchemaNameParams() beforehand.
func (o *GetV1BulkSampleSchemaNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSchemaName, rhkSchemaName, _ := route.Params.GetOK("schemaName")
	if err := o.bindSchemaName(rSchemaName, rhkSchemaName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSchemaName binds and validates parameter SchemaName from path.
func (o *GetV1BulkSampleSchemaNameParams) bindSchemaName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.SchemaName = raw

	return nil
}
