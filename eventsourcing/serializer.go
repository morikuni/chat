package eventsourcing

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
)

type Type struct {
	Package string
	Name    string
}

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
		return nil, errors.WithStack(RaiseUnknownEventError(typ))
	}
	data, err := json.Marshal(event)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal event")
	}
	return data, nil
}

func (s *jsonSerializer) Deserialize(typ Type, data []byte) (Event, error) {
	t, ok := s.eventMap[typ]
	if !ok {
		return nil, errors.WithStack(RaiseUnknownEventError(typ))
	}
	event := reflect.New(t).Interface()
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal data")
	}
	return reflect.ValueOf(event).Elem().Interface().(Event), nil
}

func TypeOf(event Event) Type {
	t := reflect.TypeOf(event)
	return Type{t.PkgPath(), t.Name()}
}
