// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/internal/models"
)

// NewUpdateNetworkPolicyParams creates a new UpdateNetworkPolicyParams object
// with the default values initialized.
func NewUpdateNetworkPolicyParams() *UpdateNetworkPolicyParams {
	var ()
	return &UpdateNetworkPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkPolicyParamsWithTimeout creates a new UpdateNetworkPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkPolicyParamsWithTimeout(timeout time.Duration) *UpdateNetworkPolicyParams {
	var ()
	return &UpdateNetworkPolicyParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkPolicyParamsWithContext creates a new UpdateNetworkPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkPolicyParamsWithContext(ctx context.Context) *UpdateNetworkPolicyParams {
	var ()
	return &UpdateNetworkPolicyParams{

		Context: ctx,
	}
}

// NewUpdateNetworkPolicyParamsWithHTTPClient creates a new UpdateNetworkPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkPolicyParamsWithHTTPClient(client *http.Client) *UpdateNetworkPolicyParams {
	var ()
	return &UpdateNetworkPolicyParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkPolicyParams contains all the parameters to send to the API endpoint
for the update network policy operation typically these are written to a http.Request
*/
type UpdateNetworkPolicyParams struct {

	/*Body*/
	Body *models.V1UpdateNetworkPolicyRequest
	/*NetworkPolicyID
	  A network policy's unique identifier

	*/
	NetworkPolicyID string
	/*NetworkPolicyStackID
	  The ID of the stack that a network policy belongs to

	*/
	NetworkPolicyStackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network policy params
func (o *UpdateNetworkPolicyParams) WithTimeout(timeout time.Duration) *UpdateNetworkPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network policy params
func (o *UpdateNetworkPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network policy params
func (o *UpdateNetworkPolicyParams) WithContext(ctx context.Context) *UpdateNetworkPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network policy params
func (o *UpdateNetworkPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network policy params
func (o *UpdateNetworkPolicyParams) WithHTTPClient(client *http.Client) *UpdateNetworkPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network policy params
func (o *UpdateNetworkPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update network policy params
func (o *UpdateNetworkPolicyParams) WithBody(body *models.V1UpdateNetworkPolicyRequest) *UpdateNetworkPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update network policy params
func (o *UpdateNetworkPolicyParams) SetBody(body *models.V1UpdateNetworkPolicyRequest) {
	o.Body = body
}

// WithNetworkPolicyID adds the networkPolicyID to the update network policy params
func (o *UpdateNetworkPolicyParams) WithNetworkPolicyID(networkPolicyID string) *UpdateNetworkPolicyParams {
	o.SetNetworkPolicyID(networkPolicyID)
	return o
}

// SetNetworkPolicyID adds the networkPolicyId to the update network policy params
func (o *UpdateNetworkPolicyParams) SetNetworkPolicyID(networkPolicyID string) {
	o.NetworkPolicyID = networkPolicyID
}

// WithNetworkPolicyStackID adds the networkPolicyStackID to the update network policy params
func (o *UpdateNetworkPolicyParams) WithNetworkPolicyStackID(networkPolicyStackID string) *UpdateNetworkPolicyParams {
	o.SetNetworkPolicyStackID(networkPolicyStackID)
	return o
}

// SetNetworkPolicyStackID adds the networkPolicyStackId to the update network policy params
func (o *UpdateNetworkPolicyParams) SetNetworkPolicyStackID(networkPolicyStackID string) {
	o.NetworkPolicyStackID = networkPolicyStackID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param network_policy.id
	if err := r.SetPathParam("network_policy.id", o.NetworkPolicyID); err != nil {
		return err
	}

	// path param network_policy.stack_id
	if err := r.SetPathParam("network_policy.stack_id", o.NetworkPolicyStackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
