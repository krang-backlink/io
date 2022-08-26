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
		Meta    Meta          `json:"meta" bson:"meta"`
	}
	// Meta represents the attributes of a failed task.
	Meta struct {
		GroupSlug  string         `json:"group_slug" bson:"group_slug"`
		ProjectID  int64          `json:"project_id" bson:"project_id"`
		TaskID     int64          `json:"task_id" bson:"task_id"`
		ScrapeID   string         `json:"scrape_id" bson:"scrape_id"`
		URL        string         `json:"url" bson:"url"`
		SearchTerm string         `json:"search_term" bson:"search_term"`
		Data       map[string]any `json:"data" bson:"data"`
	}
	// marshalError is the error sent when a function
	// or service failed.
	marshalError struct {
		Error   wrappingError `json:"error" bson:"error"`
		Service string        `json:"service" bson:"service"`
		Meta    Meta          `json:"meta" bson:"meta"`
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

// NewError returns a new Lambda error.
func NewError(err error, service string, meta Meta) *Error {
	return &Error{
		Err:     errors.ToError(err),
		Service: service,
		Meta:    meta,
	}
}

// Error returns the JSON representation of the error
// message by implementing the error interface.
func (e *Error) Error() string {
	b, err := json.Marshal(e)
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
