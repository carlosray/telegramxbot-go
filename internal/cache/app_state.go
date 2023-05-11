package cache

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type AppState struct {
	startedAt     time.Time
	handlersNames []string
}

var current *AppState

func Initialize(startedAt time.Time, handlersNames []string) (*AppState, error) {
	if current != nil {
		return current, errors.New("app state already initialized")
	}
	current = &AppState{startedAt, handlersNames}
	return current, nil
}

func GetAppState() (*AppState, error) {
	if current == nil {
		return nil, errors.New("app state is not initialized")
	}
	return current, nil
}

func (a *AppState) StartedAt() time.Time {
	return a.startedAt
}

func (a *AppState) Handlers() []string {
	return a.handlersNames
}

func (a *AppState) String() string {
	return fmt.Sprintf("{ \"startedAt\": \"%s\", \"handlersNames\": \"%v\" }",
		a.startedAt.Format(time.RFC1123),
		fmt.Sprintf("[%s]", strings.Join(a.handlersNames, ",")),
	)
}
