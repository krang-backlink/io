// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krangio

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/errors"
	"reflect"
	"testing"
)

var (
	err = Error{
		Err:     errors.NewInternal(errors.New("error"), "message", "op"),
		Service: "service",
		Meta: Meta{
			GroupSlug: "slug",
		},
	}
)

func TestNewLambdaError(t *testing.T) {
	want := &Error{
		Err:     err.Err,
		Service: err.Service,
		Meta:    err.Meta,
	}
	got := NewError(err.Err, err.Service, err.Meta)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expecting %s, got %s", want, got)
	}
}

func TestLambdaError_Error(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		want := err
		got := marshalError{}
		err := json.Unmarshal([]byte(err.Error()), &got)
		if err != nil {
			t.Fatalf("expecting no error got %s", err.Error())
		}
		if !reflect.DeepEqual(want.Service, got.Service) {
			t.Fatalf("expecting %s, got %s", want.Service, got.Service)
		}
		if !reflect.DeepEqual(want.Meta, got.Meta) {
			t.Fatalf("expecting %s, got %s", want.Service, got.Service)
		}
		if !reflect.DeepEqual(want.Err.Message, got.Error.Message) {
			t.Fatalf("expecting %s, got %s", want.Err.Message, got.Error.Message)
		}
	})

	t.Run("Error", func(t *testing.T) {
		e := &Error{Meta: Meta{Data: map[string]any{"wrong": make(chan int)}}}
		got := e.Error()
		want := "error marshalling lambda error"
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("expecting %s, got %s", want, got)
		}
	})
}

func TestError_ToMap(t *testing.T) {
	t.Run("With", func(t *testing.T) {
		got := err.ToMap()
		want := map[string]any{
			"code":      err.Err.Code,
			"message":   err.Err.Message,
			"operation": err.Err.Operation,
			"error":     err.Err.Error(),
			"file_line": err.Err.FileLine(),
		}
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("expecting %s, got %s", want, got)
		}
	})

	t.Run("Without", func(t *testing.T) {
		tmp := Error{}
		got := tmp.ToMap()
		var want map[string]any
		if !reflect.DeepEqual(want, got) {
			fmt.Println(got)
			t.Fatalf("expecting %s, got %s", want, got)
		}
	})
}
