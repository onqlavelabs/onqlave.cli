package utils

import (
	"encoding/json"
	"fmt"
)

type eventPayloadHandler[T any] struct {
}

func (s *eventPayloadHandler[T]) Extract(data interface{}) (*T, error) {
	var result T
	var jsonString []byte
	switch data.(type) {
	case string:
		jsonString = []byte(fmt.Sprintf("%v", data))
		// convert json to struct
		err := json.Unmarshal(jsonString, &result)
		if err != nil {
			return nil, err
		}
	default:
		jsonString, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		// convert json to struct
		err = json.Unmarshal(jsonString, &result)
		if err != nil {
			return nil, err
		}
	}
	return &result, nil
}

func NewEventPayloadHandler[T any]() *eventPayloadHandler[T] {
	c := eventPayloadHandler[T]{}
	return &c
}
