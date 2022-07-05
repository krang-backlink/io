// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krangio

//func TestPage_ToProto(t *testing.T) {
//	scrapeID := primitive.NewObjectID()
//	want := Page{
//		ID:         primitive.NewObjectID(),
//		ScrapeID:   &scrapeID,
//		URL:        "url",
//		GroupSlug:  "slug",
//		TaskID:     1,
//		SearchTerm: "term",
//	}
//	got := want.ToProto()
//	if !reflect.DeepEqual(want.ID.Hex(), got.Id) {
//		t.Fatalf("expecting %s, got %s", want.ID.Hex(), got.Id)
//	}
//	if !reflect.DeepEqual(want.ScrapeID.Hex(), got.ScrapeId) {
//		t.Fatalf("expecting %s, got %s", want.ScrapeID.Hex(), got.ScrapeId)
//	}
//	if !reflect.DeepEqual(want.URL, got.Url) {
//		t.Fatalf("expecting %s, got %s", want.URL, got.Url)
//	}
//	if !reflect.DeepEqual(want.GroupSlug, got.GroupSlug) {
//		t.Fatalf("expecting %s, got %s", want.GroupSlug, got.GroupSlug)
//	}
//	if !reflect.DeepEqual(want.TaskID, got.TaskId) {
//		t.Fatalf("expecting %d, got %d", want.TaskID, got.TaskId)
//	}
//	if !reflect.DeepEqual(want.SearchTerm, got.SearchTerm) {
//		t.Fatalf("expecting %s, got %s", want.SearchTerm, got.SearchTerm)
//	}
//}
