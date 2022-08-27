// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krangio

import (
	"encoding/json"
	"github.com/ainsleyclark/errors"
)

type (
	// Error represents an error that occurred during the
	// processing of a Krang Lambda function.
	Error struct {
		Err     *errors.Error `json:"error" bson:"error"`
		Service string        `json:"service" bson:"service"` // Currently running function, for example "scrape"
	}
	// marshalError is the error sent when a function
	// or service failed.
	marshalError struct {
		Error   wrappingError `json:"error" bson:"error"`
		Service string        `json:"service" bson:"service"`
	}
	// wrappingError is the wrapping error features the error
	// and file line in strings suitable for json.Marshal
	wrappingError struct {
		Code      string `json:"code" bson:"code"`
		Message   string `json:"message" bson:"message"`
		Operation string `json:"operation" bson:"operation"`
		Err       string `json:"error" bson:"error"`
		FileLine  string `json:"file_line" bson:"file_line"`
	}
)

// marshal is an alias of json.Marshal
var marshal = json.Marshal

// NewError returns a new Lambda error.
func NewError(err error, service string) *Error {
	return &Error{
		Err:     errors.ToError(err),
		Service: service,
	}
}

// Error returns the JSON representation of the error
// message by implementing the error interface.
func (e *Error) Error() string {
	b, err := marshal(e)
	if err != nil {
		return "error marshalling lambda error"
	}
	return string(b[:])
}

// ToMap returns a map of the error if there is one.
func (e *Error) ToMap() map[string]any {
	if e == nil || e.Err == nil {
		return nil
	}
	return map[string]any{
		"code":      e.Err.Code,
		"message":   e.Err.Message,
		"operation": e.Err.Operation,
		"error":     e.Err.Error(),
		"file_line": e.Err.FileLine(),
	}
}
