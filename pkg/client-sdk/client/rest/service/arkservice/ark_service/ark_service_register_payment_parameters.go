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

// NewArkServiceRegisterPaymentParams creates a new ArkServiceRegisterPaymentParams object
// with the default values initialized.
func NewArkServiceRegisterPaymentParams() *ArkServiceRegisterPaymentParams {
	var ()
	return &ArkServiceRegisterPaymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewArkServiceRegisterPaymentParamsWithTimeout creates a new ArkServiceRegisterPaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewArkServiceRegisterPaymentParamsWithTimeout(timeout time.Duration) *ArkServiceRegisterPaymentParams {
	var ()
	return &ArkServiceRegisterPaymentParams{

		timeout: timeout,
	}
}

// NewArkServiceRegisterPaymentParamsWithContext creates a new ArkServiceRegisterPaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewArkServiceRegisterPaymentParamsWithContext(ctx context.Context) *ArkServiceRegisterPaymentParams {
	var ()
	return &ArkServiceRegisterPaymentParams{

		Context: ctx,
	}
}

// NewArkServiceRegisterPaymentParamsWithHTTPClient creates a new ArkServiceRegisterPaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewArkServiceRegisterPaymentParamsWithHTTPClient(client *http.Client) *ArkServiceRegisterPaymentParams {
	var ()
	return &ArkServiceRegisterPaymentParams{
		HTTPClient: client,
	}
}

/*ArkServiceRegisterPaymentParams contains all the parameters to send to the API endpoint
for the ark service register payment operation typically these are written to a http.Request
*/
type ArkServiceRegisterPaymentParams struct {

	/*Body*/
	Body *models.V1RegisterPaymentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ark service register payment params
func (o *ArkServiceRegisterPaymentParams) WithTimeout(timeout time.Duration) *ArkServiceRegisterPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ark service register payment params
func (o *ArkServiceRegisterPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ark service register payment params
func (o *ArkServiceRegisterPaymentParams) WithContext(ctx context.Context) *ArkServiceRegisterPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ark service register payment params
func (o *ArkServiceRegisterPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ark service register payment params
func (o *ArkServiceRegisterPaymentParams) WithHTTPClient(client *http.Client) *ArkServiceRegisterPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ark service register payment params
func (o *ArkServiceRegisterPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ark service register payment params
func (o *ArkServiceRegisterPaymentParams) WithBody(body *models.V1RegisterPaymentRequest) *ArkServiceRegisterPaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ark service register payment params
func (o *ArkServiceRegisterPaymentParams) SetBody(body *models.V1RegisterPaymentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ArkServiceRegisterPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
