package runners

import (
	"fmt"
)

// ------- Singleton ------
var runnersMapper *RunnersMapper

func GetRunnersMapper() *RunnersMapper {
	if (runnersMapper == nil) {
		runnersMapper = NewRunnersMapper()
	}
	return runnersMapper
}

func GetRunner(runnerType string) (Runner, error) {
	return GetRunnersMapper().Get(runnerType)
}

func AddRunner(runnerType string, runner Runner) {
	GetRunnersMapper().Add(runnerType, runner)
}

// ------- Factory ------
type RunnersMapper struct {
	runners map[string]Runner
}

func NewRunnersMapper() *RunnersMapper {
	mapper := &RunnersMapper{
		runners: make(map[string]Runner),
	}
	return mapper
}

// TODO: we need to design&think about should we initialize runner at the start of delegate or at the first time of get this runner
// They are big difference for customers. One will be 'delegate cannot start' <==> big incident. Second will be partial function not
// available while other tasks are ok.
func (rm *RunnersMapper) Add(runnerType string, runner Runner) {
	rm.runners[runnerType] = runner
}

func (rm *RunnersMapper) Get(runnerType string) (Runner,error) {
	if _, ok := rm.runners[runnerType]; !ok {
		return nil, fmt.Errorf("Runner %s not found", runnerType)
	}
	return rm.runners[runnerType], nil
}