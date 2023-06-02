// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewBuyCandyParams creates a new BuyCandyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBuyCandyParams() *BuyCandyParams {
	return &BuyCandyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBuyCandyParamsWithTimeout creates a new BuyCandyParams object
// with the ability to set a timeout on a request.
func NewBuyCandyParamsWithTimeout(timeout time.Duration) *BuyCandyParams {
	return &BuyCandyParams{
		timeout: timeout,
	}
}

// NewBuyCandyParamsWithContext creates a new BuyCandyParams object
// with the ability to set a context for a request.
func NewBuyCandyParamsWithContext(ctx context.Context) *BuyCandyParams {
	return &BuyCandyParams{
		Context: ctx,
	}
}

// NewBuyCandyParamsWithHTTPClient creates a new BuyCandyParams object
// with the ability to set a custom HTTPClient for a request.
func NewBuyCandyParamsWithHTTPClient(client *http.Client) *BuyCandyParams {
	return &BuyCandyParams{
		HTTPClient: client,
	}
}

/*
BuyCandyParams contains all the parameters to send to the API endpoint

	for the buy candy operation.

	Typically these are written to a http.Request.
*/
type BuyCandyParams struct {

	/* Order.

	   summary of the candy order
	*/
	Order BuyCandyBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the buy candy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuyCandyParams) WithDefaults() *BuyCandyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the buy candy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuyCandyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the buy candy params
func (o *BuyCandyParams) WithTimeout(timeout time.Duration) *BuyCandyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the buy candy params
func (o *BuyCandyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the buy candy params
func (o *BuyCandyParams) WithContext(ctx context.Context) *BuyCandyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the buy candy params
func (o *BuyCandyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the buy candy params
func (o *BuyCandyParams) WithHTTPClient(client *http.Client) *BuyCandyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the buy candy params
func (o *BuyCandyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrder adds the order to the buy candy params
func (o *BuyCandyParams) WithOrder(order BuyCandyBody) *BuyCandyParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the buy candy params
func (o *BuyCandyParams) SetOrder(order BuyCandyBody) {
	o.Order = order
}

// WriteToRequest writes these params to a swagger request
func (o *BuyCandyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Order); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
