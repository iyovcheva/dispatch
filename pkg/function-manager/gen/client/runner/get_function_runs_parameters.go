///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFunctionRunsParams creates a new GetFunctionRunsParams object
// with the default values initialized.
func NewGetFunctionRunsParams() *GetFunctionRunsParams {
	var ()
	return &GetFunctionRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFunctionRunsParamsWithTimeout creates a new GetFunctionRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFunctionRunsParamsWithTimeout(timeout time.Duration) *GetFunctionRunsParams {
	var ()
	return &GetFunctionRunsParams{

		timeout: timeout,
	}
}

// NewGetFunctionRunsParamsWithContext creates a new GetFunctionRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFunctionRunsParamsWithContext(ctx context.Context) *GetFunctionRunsParams {
	var ()
	return &GetFunctionRunsParams{

		Context: ctx,
	}
}

// NewGetFunctionRunsParamsWithHTTPClient creates a new GetFunctionRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFunctionRunsParamsWithHTTPClient(client *http.Client) *GetFunctionRunsParams {
	var ()
	return &GetFunctionRunsParams{
		HTTPClient: client,
	}
}

/*GetFunctionRunsParams contains all the parameters to send to the API endpoint
for the get function runs operation typically these are written to a http.Request
*/
type GetFunctionRunsParams struct {

	/*FunctionName
	  Name of function to run

	*/
	FunctionName string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get function runs params
func (o *GetFunctionRunsParams) WithTimeout(timeout time.Duration) *GetFunctionRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get function runs params
func (o *GetFunctionRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get function runs params
func (o *GetFunctionRunsParams) WithContext(ctx context.Context) *GetFunctionRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get function runs params
func (o *GetFunctionRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get function runs params
func (o *GetFunctionRunsParams) WithHTTPClient(client *http.Client) *GetFunctionRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get function runs params
func (o *GetFunctionRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFunctionName adds the functionName to the get function runs params
func (o *GetFunctionRunsParams) WithFunctionName(functionName string) *GetFunctionRunsParams {
	o.SetFunctionName(functionName)
	return o
}

// SetFunctionName adds the functionName to the get function runs params
func (o *GetFunctionRunsParams) SetFunctionName(functionName string) {
	o.FunctionName = functionName
}

// WithTags adds the tags to the get function runs params
func (o *GetFunctionRunsParams) WithTags(tags []string) *GetFunctionRunsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get function runs params
func (o *GetFunctionRunsParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetFunctionRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param functionName
	if err := r.SetPathParam("functionName", o.FunctionName); err != nil {
		return err
	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "multi")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
