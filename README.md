# Serialized Go client
![Build Status](https://github.com/serialized-io/serialized-go/workflows/Go/badge.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/serialized-io/serialized-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/serialized-io/serialized-go)](https://goreportcard.com/report/github.com/serialized-io/serialized-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

Provides access to the [Serialized](https://serialized.io) API.

## Installation

```
go get github.com/serialized-io/serialized-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/serialized-io/serialized-go"
)

func main() {
    var (
        accessKey       = os.Getenv("SERIALIZED_ACCESS_KEY")
        secretAccessKey = os.Getenv("SERIALIZED_SECRET_ACCESS_KEY")
    )

    client := serialized.NewClient(
        serialized.WithAccessKey(accessKey),
        serialized.WithSecretAccessKey(secretAccessKey),
    )

    err := client.Feed(context.Background(), "order", 0, func(entry *serialized.FeedEntry) {
        for _, event := range entry.Events {
            if event.Type == "OrderPaidEvent" {
                fmt.Printf("The order with ID %s was paid\n", entry.AggregateID)
            }
        }
    })
    if err != nil {
        log.Fatal(err)
    }
}
```
