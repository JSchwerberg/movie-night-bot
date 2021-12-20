// Code generated by go-swagger; DO NOT EDIT.

package id_parameter

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

// NewGetIDParams creates a new GetIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIDParams() *GetIDParams {
	return &GetIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIDParamsWithTimeout creates a new GetIDParams object
// with the ability to set a timeout on a request.
func NewGetIDParamsWithTimeout(timeout time.Duration) *GetIDParams {
	return &GetIDParams{
		timeout: timeout,
	}
}

// NewGetIDParamsWithContext creates a new GetIDParams object
// with the ability to set a context for a request.
func NewGetIDParamsWithContext(ctx context.Context) *GetIDParams {
	return &GetIDParams{
		Context: ctx,
	}
}

// NewGetIDParamsWithHTTPClient creates a new GetIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIDParamsWithHTTPClient(client *http.Client) *GetIDParams {
	return &GetIDParams{
		HTTPClient: client,
	}
}

/* GetIDParams contains all the parameters to send to the API endpoint
   for the get Id operation.

   Typically these are written to a http.Request.
*/
type GetIDParams struct {

	/* Callback.

	   JSONP callback name
	*/
	Callback *string

	/* I.

	   A valid IMDb ID (e.g. tt0000001)
	*/
	I string

	/* Plot.

	   Return short or full plot
	*/
	Plot *string

	/* R.

	   The response type to return
	*/
	R *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIDParams) WithDefaults() *GetIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Id params
func (o *GetIDParams) WithTimeout(timeout time.Duration) *GetIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Id params
func (o *GetIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Id params
func (o *GetIDParams) WithContext(ctx context.Context) *GetIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Id params
func (o *GetIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Id params
func (o *GetIDParams) WithHTTPClient(client *http.Client) *GetIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Id params
func (o *GetIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallback adds the callback to the get Id params
func (o *GetIDParams) WithCallback(callback *string) *GetIDParams {
	o.SetCallback(callback)
	return o
}

// SetCallback adds the callback to the get Id params
func (o *GetIDParams) SetCallback(callback *string) {
	o.Callback = callback
}

// WithI adds the i to the get Id params
func (o *GetIDParams) WithI(i string) *GetIDParams {
	o.SetI(i)
	return o
}

// SetI adds the i to the get Id params
func (o *GetIDParams) SetI(i string) {
	o.I = i
}

// WithPlot adds the plot to the get Id params
func (o *GetIDParams) WithPlot(plot *string) *GetIDParams {
	o.SetPlot(plot)
	return o
}

// SetPlot adds the plot to the get Id params
func (o *GetIDParams) SetPlot(plot *string) {
	o.Plot = plot
}

// WithR adds the r to the get Id params
func (o *GetIDParams) WithR(r *string) *GetIDParams {
	o.SetR(r)
	return o
}

// SetR adds the r to the get Id params
func (o *GetIDParams) SetR(r *string) {
	o.R = r
}

// WriteToRequest writes these params to a swagger request
func (o *GetIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Callback != nil {

		// query param callback
		var qrCallback string

		if o.Callback != nil {
			qrCallback = *o.Callback
		}
		qCallback := qrCallback
		if qCallback != "" {

			if err := r.SetQueryParam("callback", qCallback); err != nil {
				return err
			}
		}
	}

	// query param i
	qrI := o.I
	qI := qrI
	if qI != "" {

		if err := r.SetQueryParam("i", qI); err != nil {
			return err
		}
	}

	if o.Plot != nil {

		// query param plot
		var qrPlot string

		if o.Plot != nil {
			qrPlot = *o.Plot
		}
		qPlot := qrPlot
		if qPlot != "" {

			if err := r.SetQueryParam("plot", qPlot); err != nil {
				return err
			}
		}
	}

	if o.R != nil {

		// query param r
		var qrR string

		if o.R != nil {
			qrR = *o.R
		}
		qR := qrR
		if qR != "" {

			if err := r.SetQueryParam("r", qR); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
