# Krang IO

This module contains the Go bindings and Proto file definitions and types for Krang and it's microservices to be used by
GRPC servers.

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Test](https://github.com/krang-backlink/io/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/krang-backlink/io/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/krang-backlink/io/branch/master/graph/badge.svg?token=HXEZIJ2NUY)](https://codecov.io/gh/krang-backlink/io)

## Installation

```
go get -u github.com/krang-backlink/io
```

## Types

All tasks should accept and return a `krangio.Page`, see `types.go` for more detailed insight into it's field data.
Below is an example of the Page struct stored in Mongo.

```go
type (
	// Page represents an individual task scrape including
	// metadata from the Task.
	Page struct {
		ID             primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
		ScrapeID       *primitive.ObjectID `bson:"scrape_id" json:"scrape_id"`
		URL            string              `bson:"url" json:"url"`
		GroupSlug      string              `bson:"group_slug,omitempty" json:"group_slug"`
		TaskID         int64               `bson:"task_id,omitempty" json:"task_id"`
		SearchTerm     string              `json:"search_term" bson:"search_term"`
		RelevancyScore uint                `json:"relevancy_score" bson:"relevancy_score"`
		SiteScore      uint                `json:"site_score" bson:"site_score"`
		Scrape         Scrape              `bson:"scrape" json:"scrape"`
		UpdatedAt      time.Time           `bson:"updated_at" json:"updated_at"`
		CreatedAt      time.Time           `bson:"created_at" json:"created_at"`
	}
)
```

## Errors

Any error that is returned from a Lambda function should be of type `krangio.Error` for error reporting. See below on
how to create and consume one.

### Types

```go
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
		// Any additional data
		Data       map[string]any `json:"data"`
	}
)
```

## Create a new Lambda error

```go
meta := krangio.Meta{
	GroupSlug:  in.GroupSlug,
	TaskID:     in.TaskID,
	URL:        in.URL,
	SearchTerm: in.SearchTerm,
}

status, err := myThing()
if err != nil {
	return in, lambda.NewError(err, ServiceName, meta)
}
```

## Proto

Each service contains its own proto definition along with its implementation for Server/Client.

```protobuf
message CompleteRequest {
	string id = 1;
	string scrape_id = 2;
	string url = 3;
	string group_slug = 4;
	int64 task_id = 5;
	string search_term = 6;
}

message Response {
	bool error = 2;
	string message = 3;
}

service TasksService {
	rpc CompleteTask(CompleteRequest) returns(Response) {}
}
```

## Usage

```go
func Send() error {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials())
	if err != nil {
	return err
	}

	s := proto.NewTasksServiceClient(conn)

	response, err := s.Scrape(context.Background(), &proto.CompleteRequest{
		GroupSlug: "",
		TaskId:    0,
		PageId:    "",
		Url:       "",
	})
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}
```

## Development

To set up this repository, run:

```bash
make setup
```

To generate the proto files run:

```bash
make generate
```
