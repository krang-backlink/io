// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krangio

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
)

func TestPage_HasScrape(t *testing.T) {
	hex := primitive.NewObjectID()

	tt := map[string]struct {
		input Page
		want  bool
	}{
		"True": {
			Page{ScrapeID: &hex},
			true,
		},
		"False": {
			Page{},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.HasScrape()
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("expecting %v, got %v", test.want, got)
			}
		})
	}
}

func TestGetObjectID(t *testing.T) {
	hex := primitive.NewObjectID()

	tt := map[string]struct {
		input string
		want  *primitive.ObjectID
	}{
		"Nil": {
			"",
			nil,
		},
		"Error": {
			hex.Hex(),
			&hex,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := GetObjectID(test.input)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("expecting %s, got %s", test.want, got)
			}
		})
	}
}
