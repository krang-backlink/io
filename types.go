// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krangio

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Page represents an individual task scrape including
	// metadata from the Task.
	Page struct {
		ID             primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
		ScrapeID       *primitive.ObjectID `json:"scrape_id" bson:"scrape_id"`
		UUID           string              `json:"uuid,omitempty" bson:"-"` // Used for SQS dedupe.
		URL            string              `json:"url" bson:"url"`
		GroupSlug      string              `json:"group_slug" bson:"group_slug"`
		ProjectID      int64               `json:"project_id" bson:"project_id"`
		TaskID         int64               `json:"task_id" bson:"task_id"`
		SearchTerm     string              `json:"search_term" bson:"search_term"`
		RelevancyScore int                 `json:"relevancy_score" bson:"relevancy_score"`
		SiteScore      int                 `json:"site_score" bson:"site_score"`
		Scrape         Scrape              `json:"scrape" bson:"scrape,omitempty"`
		Status         ScrapeStatus        `json:"status" bson:"status"`
		Usage          PageUsage           `json:"usage" bson:"usage"`
		UpdatedAt      time.Time           `json:"updated_at" bson:"updated_at"`
		CreatedAt      time.Time           `json:"created_at" bson:"created_at"`
	} //@name Page
	// Scrape represents an individual scrape of a page and its
	// various metrics.
	Scrape struct {
		ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		URL        string             `json:"-" bson:"url" swagggerignore:"true"`
		HTTPStatus int                `json:"http_status" bson:"http_status"`
		Content    ScrapeContent      `json:"content" bson:"content"`
		Metrics    ScrapeMetrics      `json:"metrics" bson:"metrics"`
		Message    string             `json:"message" bson:"message"`
		Status     ScrapeStatus       `json:"status" bson:"status"`
		Error      any                `json:"error" bson:"error"`
		Service    string             `json:"service" bson:"service"` // Currently running function, for example "scrape"`
		UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
		CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	} //@name Scrape
	// ScrapeContent represents the HTML markup of a page including any
	// <body> content that's relevant for scoring.
	ScrapeContent struct {
		H1            string          `json:"h1" bson:"h1"`
		H2            string          `json:"h2" bson:"h2"`
		Title         string          `json:"title" bson:"title"`
		ExternalLinks int             `json:"external_links" bson:"external_links"`
		Keywords      []ScrapeKeyword `json:"keywords" bson:"keywords"`
	} //@name ScrapeContent
	// ScrapeKeyword represents a singular entity extracted from
	// a given piece of text.
	ScrapeKeyword struct {
		Term     string  `json:"term" bson:"term"`
		Salience float64 `json:"salience" bson:"salience"`
	} //@name ScrapeKeyword
	// ScrapeMetrics represents the scores and metrics retrieved from
	// Ahrefs, Moz and Majestic.
	ScrapeMetrics struct {
		Ahrefs ScrapeAhrefs `json:"ahrefs" bson:"ahrefs"`
	} //@name ScrapeMetrics
	// ScrapeAhrefs represents the metrics retrieved from the
	// Ahrefs API including cost, rows and if it was cached.
	ScrapeAhrefs struct {
		DR   float64  `json:"dr" bson:"dr"`     // Domain Ranking
		Rank *float64 `json:"rank" bson:"rank"` // Ahrefs Rank
	} //@name ScrapeAhrefs
	// BackLinkCheck represents the data sent to the Lambda function
	// for checking if a backlink appears on the page.
	BackLinkCheck struct {
		GroupSlug string `json:"group_slug" bson:"group_slug"`
		LinkID    int64  `json:"link_id" bson:"link_id"`
		URL       string `json:"url" bson:"url"`
		Link      string `json:"link" bson:"link"`
	} //@name BackLinkCheck
	// ScrapeStatus status represents the status of a page task.
	ScrapeStatus string
	// PageUsage represents any costs that have been associated
	// with the page.
	PageUsage struct {
		Ahrefs PageUsageAhrefs `json:"ahrefs" bson:"ahrefs"`
	} //@name PageUsage
	// PageUsageAhrefs represents the total amount of cost a
	// singular call to Ahrefs cost.
	PageUsageAhrefs struct {
		Rows         int  `json:"rows_used" bson:"rows"`
		UnitCostRows int  `json:"unit_cost_rows" bson:"unit_cost_rows"`
		Cached       bool `json:"cached" bson:"cached"`
	} //@name PageUsageAhrefs
)

const (
	// ScrapeStatusProcessing is the status that defines
	// a processing page.
	ScrapeStatusProcessing ScrapeStatus = "processing"
	// ScrapeStatusFailed is the status that defines
	// a failed page task.
	ScrapeStatusFailed ScrapeStatus = "failed"
	// ScrapeStatusTimedOut is the status that defines
	// a timed out page task.
	ScrapeStatusTimedOut ScrapeStatus = "timed-out"
	// ScrapeStatusSuccess is the status that defines
	// a successful page task.
	ScrapeStatusSuccess ScrapeStatus = "success"
)

// HasScrape determines if a page has a Scrape ID
// attached to it.
func (p *Page) HasScrape() bool {
	return p.ScrapeID != nil
}

// LogMessage returns a formatted message for processing
// Lambda functions.
func (p *Page) LogMessage(service string) string {
	return fmt.Sprintf("Starting Lambda Function: %s, URL: %s", service, p.URL)
}

// LoggerFields returns logrus Fields to log the Page
// meta data.
func (p *Page) LoggerFields(service string) map[string]any {
	scrapeID := ""
	if p.ScrapeID != nil {
		scrapeID = p.ScrapeID.String()
	}
	return map[string]any{
		"service":     service,
		"id":          p.ID.String(),
		"scrape_id":   scrapeID,
		"uuid":        p.UUID,
		"group_slug":  p.GroupSlug,
		"project_id":  p.ProjectID,
		"task_id":     p.TaskID,
		"url":         p.URL,
		"search_term": p.SearchTerm,
	}
}

// GetObjectID returns the primitive.ObjectID if there
// is one set, otherwise it returns nil.
func GetObjectID(hex string) *primitive.ObjectID {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil
	}
	return &id
}
