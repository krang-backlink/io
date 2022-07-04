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
		Err     *errors.Error `json:"error"`
		Service string        `json:"service"` // Currently running function, for example "scrape"
		Meta    Meta          `json:"meta"`
	}
	// Meta represents the attributes of a failed task.
	Meta struct {
		GroupSlug  string         `json:"group_slug"`
		TaskID     int64          `json:"task_id"`
		ScrapeID   string         `json:"scrape_id"`
		URL        string         `json:"url"`
		SearchTerm string         `json:"search_term"`
		Data       map[string]any `json:"data"`
	}
	// marshalError is the error sent when a function
	// or service failed.
	marshalError struct {
		Error   wrappingError `json:"error"`
		Service string        `json:"service"`
		Meta    Meta          `json:"meta"`
	}
	// wrappingError is the wrapping error features the error
	// and file line in strings suitable for json.Marshal
	wrappingError struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Operation string `json:"operation"`
		Err       string `json:"error"`
		FileLine  string `json:"file_line"`
	}
)

// NewError returns a new Lambda error.
func NewError(err error, service string, meta Meta) *Error {
	return &Error{
		Err:     errors.ToError(err),
		Service: service,
		Meta:    meta,
	}
}

// MarshalJSON implements json encode.MarshalJSON to transform
// the error into a formatted string.
//func (e *Error) MarshalJSON() ([]byte, error) {
//	var err wrappingError
//	if e.Err != nil {
//		err = wrappingError{
//			Code:      e.Err.Code,
//			Message:   e.Err.Message,
//			Operation: e.Err.Message,
//			Err:       e.Err.Error(),
//			FileLine:  e.Err.FileLine(),
//		}
//	}
//	me := marshalError{
//		Error:   err,
//		Service: e.Service,
//		Meta:    e.Meta,
//	}
//	return json.Marshal(me)
//}

// Error returns the JSON representation of the error
// message by implementing the error interface.
func (e *Error) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "error marshalling lambda error"
	}
	return string(b[:])
}
