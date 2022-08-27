// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krangio

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	err = Error{
		Err:     errors.NewInternal(errors.New("error"), "message", "op"),
		Service: "service",
	}
)

func TestNewLambdaError(t *testing.T) {
	want := &Error{
		Err:     err.Err,
		Service: err.Service,
	}
	got := NewError(err.Err, err.Service)
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
		if !reflect.DeepEqual(want.Err.Message, got.Error.Message) {
			t.Fatalf("expecting %s, got %s", want.Err.Message, got.Error.Message)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		orig := marshal
		defer func() {
			marshal = orig
		}()
		marshal = func(v any) ([]byte, error) {
			return nil, errors.New("marshal error")
		}
		e := &Error{}
		got := e.Error()
		assert.Contains(t, got, "error marshalling lambda error")
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
