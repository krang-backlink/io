// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krangio

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Page represents an individual task scrape including
	// metadata from the Task.
	Page struct {
		ID             primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
		ScrapeID       *primitive.ObjectID `json:"scrape_id" bson:"scrape_id"`
		URL            string              `json:"url" bson:"url"`
		GroupSlug      string              `json:"group_slug" bson:"group_slug,omitempty"`
		ProjectID      int64               `json:"project_id" bson:"project_id,omitempty"`
		TaskID         int64               `json:"task_id" bson:"task_id,omitempty" `
		TemplateID     *int64              `json:"template_id" bson:"template_id,omitempty" `
		SearchTerm     string              `json:"search_term" bson:"search_term"`
		RelevancyScore int                 `json:"relevancy_score" bson:"relevancy_score"`
		SiteScore      int                 `json:"site_score" bson:"site_score"`
		Scrape         Scrape              `json:"scrape" bson:"scrape"`
		Message        string              `json:"message" bson:"message"`
		Status         PageStatus          `json:"status" bson:"status"`
		UpdatedAt      time.Time           `json:"updated_at" bson:"updated_at"`
		CreatedAt      time.Time           `json:"created_at" bson:"created_at"`
	}
	// Scrape represents an individual scrape of a page and its
	// various metrics.
	Scrape struct {
		ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		URL        string             `json:"url" bson:"url"`
		HTTPStatus int                `json:"http_status" bson:"http_status,omitempty"`
		Content    ScrapeContent      `json:"content" bson:"content,omitempty"`
		Metrics    ScrapeMetrics      `json:"metrics" bson:"metrics,omitempty"`
	}
	// ScrapeContent represents the HTML markup of a page including any
	// <body> content that's relevant for scoring.
	ScrapeContent struct {
		H1            string         `json:"h1" bson:"h1"`
		H2            string         `json:"h2" bson:"h2"`
		Title         string         `json:"title" bson:"title"`
		ExternalLinks int            `json:"external_links" bson:"external_links"`
		Keywords      ScrapeKeywords `json:"keywords" bson:"keywords"`
	}
	// ScrapeKeywords represents a singular entity extracted from
	// a given piece of text.
	ScrapeKeywords struct {
		Term     string  `json:"term"`
		Salience float64 `json:"salience"`
	}
	// ScrapeMetrics represents the scores and metrics retrieved from
	// Ahrefs, Moz and Majestic.
	ScrapeMetrics struct {
		Backlinks    int           `json:"backlinks" bson:"backlinks"`
		LoadingTime  time.Duration `json:"loading_time" bson:"loading_time"`
		AhrefsDR     int           `json:"ahrefs_dr" bson:"ahrefs_dr"` // Domain Ranking
		MozPA        int           `json:"moz_pa" bson:"moz_pa"`       // Page Authority
		MozDA        int           `json:"moz_da" bson:"moz_da"`       // Domain Authority
		MozSpamScore int           `json:"moz_spam_score" bson:"moz_spam_score"`
		MajesticCF   int           `json:"majestic_cf" bson:"majestic_cf"` // Citation Flow
		MajesticTF   int           `json:"majestic_tf" bson:"majestic_tf"` // Trust Flow
	}
	// PageStatus status represents the status of a page task.
	PageStatus string
)

const (
	// PageStatusFailed is the status that defines
	// a failed page task.
	PageStatusFailed PageStatus = "failed"
	// PageStatusTimedOut is the status that defines
	// a timed out page task.
	PageStatusTimedOut PageStatus = "timed-out"
	// PageStatusSuccess is the status that defines
	// a successful page task.
	PageStatusSuccess PageStatus = "success"
)
