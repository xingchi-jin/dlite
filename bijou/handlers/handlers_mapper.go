package handlers

import (
	"github.com/wings-software/dlite/client"
)

const (
	SETUP   string = "SETUP"
	EXECUTE string = "EXECUTE"
	CLEANUP string = "CLEANUP"
)

type HandlersMapper struct {
	routes map[string]Handler
}

func (hm *HandlersMapper) Add(eventType string, h Handler) {
	if _, ok := hm.routes[eventType]; !ok {
		hm.routes[eventType] = h
	}
}

func (hm *HandlersMapper) Get(eventType string) Handler {
	if _, ok := hm.routes[eventType]; !ok {
		return nil
	}
	return hm.routes[eventType]
}

func NewHandlersMapper(client client.Client) *HandlersMapper {
	mapper := &HandlersMapper{
		routes: make(map[string]Handler),
	}
	mapper.Add(SETUP, &SetupHandler{client: client})
	mapper.Add(EXECUTE, &ExecuteHandler{client: client})
	mapper.Add(CLEANUP, &CleanupHandler{client: client})
	return mapper
}
