package tasks

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

	"github.com/go-swagger/go-swagger/examples/task-tracker/models"
)

// NewCreateTaskParams creates a new CreateTaskParams object
// with the default values initialized.
func NewCreateTaskParams() *CreateTaskParams {
	var ()
	return &CreateTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTaskParamsWithTimeout creates a new CreateTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTaskParamsWithTimeout(timeout time.Duration) *CreateTaskParams {
	var ()
	return &CreateTaskParams{

		timeout: timeout,
	}
}

// NewCreateTaskParamsWithContext creates a new CreateTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTaskParamsWithContext(ctx context.Context) *CreateTaskParams {
	var ()
	return &CreateTaskParams{

		Context: ctx,
	}
}

/*CreateTaskParams contains all the parameters to send to the API endpoint
for the create task operation typically these are written to a http.Request
*/
type CreateTaskParams struct {

	/*Body
	  The task to create

	*/
	Body *models.Task

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create task params
func (o *CreateTaskParams) WithTimeout(timeout time.Duration) *CreateTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create task params
func (o *CreateTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create task params
func (o *CreateTaskParams) WithContext(ctx context.Context) *CreateTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create task params
func (o *CreateTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the create task params
func (o *CreateTaskParams) WithBody(body *models.Task) *CreateTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create task params
func (o *CreateTaskParams) SetBody(body *models.Task) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.Task)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
