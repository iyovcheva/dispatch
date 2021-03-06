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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFunctionRunsParams creates a new GetFunctionRunsParams object
// with the default values initialized.
func NewGetFunctionRunsParams() GetFunctionRunsParams {
	var ()
	return GetFunctionRunsParams{}
}

// GetFunctionRunsParams contains all the bound params for the get function runs operation
// typically these are obtained from a http.Request
//
// swagger:parameters getFunctionRuns
type GetFunctionRunsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Name of function to run
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	FunctionName string
	/*Filter based on tags
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetFunctionRunsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rFunctionName, rhkFunctionName, _ := route.Params.GetOK("functionName")
	if err := o.bindFunctionName(rFunctionName, rhkFunctionName, route.Formats); err != nil {
		res = append(res, err)
	}

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFunctionRunsParams) bindFunctionName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.FunctionName = raw

	if err := o.validateFunctionName(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetFunctionRunsParams) validateFunctionName(formats strfmt.Registry) error {

	if err := validate.Pattern("functionName", "path", o.FunctionName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (o *GetFunctionRunsParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {

	tagsIC := rawData

	if len(tagsIC) == 0 {
		return nil
	}

	var tagsIR []string
	for _, tagsIV := range tagsIC {
		tagsI := tagsIV

		tagsIR = append(tagsIR, tagsI)
	}

	o.Tags = tagsIR

	return nil
}
