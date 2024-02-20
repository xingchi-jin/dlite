package handlers


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

func NewHandlersMapper() *HandlersMapper {
	mapper := &HandlersMapper {
		routes: make(map[string]Handler),
	}
	return mapper
}