// Code generated by go-swagger; DO NOT EDIT.

package ark_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ark-network/ark/pkg/client-sdk/client/rest/service/models"
)

// NewArkServiceFinalizePaymentParams creates a new ArkServiceFinalizePaymentParams object
// with the default values initialized.
func NewArkServiceFinalizePaymentParams() *ArkServiceFinalizePaymentParams {
	var ()
	return &ArkServiceFinalizePaymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewArkServiceFinalizePaymentParamsWithTimeout creates a new ArkServiceFinalizePaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewArkServiceFinalizePaymentParamsWithTimeout(timeout time.Duration) *ArkServiceFinalizePaymentParams {
	var ()
	return &ArkServiceFinalizePaymentParams{

		timeout: timeout,
	}
}

// NewArkServiceFinalizePaymentParamsWithContext creates a new ArkServiceFinalizePaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewArkServiceFinalizePaymentParamsWithContext(ctx context.Context) *ArkServiceFinalizePaymentParams {
	var ()
	return &ArkServiceFinalizePaymentParams{

		Context: ctx,
	}
}

// NewArkServiceFinalizePaymentParamsWithHTTPClient creates a new ArkServiceFinalizePaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewArkServiceFinalizePaymentParamsWithHTTPClient(client *http.Client) *ArkServiceFinalizePaymentParams {
	var ()
	return &ArkServiceFinalizePaymentParams{
		HTTPClient: client,
	}
}

/*ArkServiceFinalizePaymentParams contains all the parameters to send to the API endpoint
for the ark service finalize payment operation typically these are written to a http.Request
*/
type ArkServiceFinalizePaymentParams struct {

	/*Body*/
	Body *models.V1FinalizePaymentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ark service finalize payment params
func (o *ArkServiceFinalizePaymentParams) WithTimeout(timeout time.Duration) *ArkServiceFinalizePaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ark service finalize payment params
func (o *ArkServiceFinalizePaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ark service finalize payment params
func (o *ArkServiceFinalizePaymentParams) WithContext(ctx context.Context) *ArkServiceFinalizePaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ark service finalize payment params
func (o *ArkServiceFinalizePaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ark service finalize payment params
func (o *ArkServiceFinalizePaymentParams) WithHTTPClient(client *http.Client) *ArkServiceFinalizePaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ark service finalize payment params
func (o *ArkServiceFinalizePaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ark service finalize payment params
func (o *ArkServiceFinalizePaymentParams) WithBody(body *models.V1FinalizePaymentRequest) *ArkServiceFinalizePaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ark service finalize payment params
func (o *ArkServiceFinalizePaymentParams) SetBody(body *models.V1FinalizePaymentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ArkServiceFinalizePaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
