package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/serialized-io/serialized-go"
)

func eventsStoreHandler(c *serialized.Client, aggType, aggID, eventType, eventID, data string, version int64) error {
	if eventID == "" {
		eventID = uuid.New().String()
	}

	if aggID == "" {
		aggID = uuid.New().String()
	}

	event := &serialized.Event{
		Type: eventType,
		ID:   eventID,
		Data: []byte(data),
	}

	return c.Store(context.Background(), aggType, aggID, version, event)
}
