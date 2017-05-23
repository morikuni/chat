package eventsourcing

import (
	"encoding/json"
	"reflect"
)

type Type string

type Serializer interface {
	Serialize(event Event) ([]byte, error)
	Deserialize(typ Type, data []byte) (Event, error)
}

type jsonSerializer struct {
	eventMap map[Type]reflect.Type
}

func NewJSONSerializer(events ...Event) Serializer {
	eventMap := make(map[Type]reflect.Type, len(events))
	for _, event := range events {
		eventMap[TypeOf(event)] = reflect.TypeOf(event)
	}
	return &jsonSerializer{eventMap}
}

func (s *jsonSerializer) Serialize(event Event) ([]byte, error) {
	typ := TypeOf(event)
	if _, ok := s.eventMap[typ]; !ok {
		return nil, RaiseUnknownEventError(typ)
	}
	return json.Marshal(event)
}

func (s *jsonSerializer) Deserialize(typ Type, data []byte) (Event, error) {
	t, ok := s.eventMap[typ]
	if !ok {
		return nil, RaiseUnknownEventError(typ)
	}
	event := reflect.New(t).Interface()
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, err
	}
	return reflect.ValueOf(event).Elem().Interface().(Event), nil
}

func TypeOf(event Event) Type {
	t := reflect.TypeOf(event)
	return Type(t.PkgPath() + "." + t.Name())
}
