package config

import (
	"context"
	"io"
	"time"
)

type StateId string

type State interface {
	Id() StateId
	NotBefore() time.Time
	NotAfter() time.Time
	PreviousId() StateId
	SuccessorId() StateId
	Dump(out io.Writer) error
}

func StateById(ctx context.Context, statedId StateId) (State, error) {
	return nil, nil
}

func LatestState(ctx context.Context) (State, error) {
	return nil, nil
}

type Serializable interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type JsonSerializable interface {
	Serializable
	MarshalToJson() ([]byte, error)
	UnmarshalFromJson([]byte) error
}
