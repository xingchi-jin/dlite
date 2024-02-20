package handlers

import (
	"context"

    "github.com/wings-software/dlite/bijou/runners"
	"github.com/wings-software/dlite/client/proto"
)

type Handler interface {
	Handle(ctx context.Context, runnerType string, payload *proto.TaskPayload)
}

type SetupHandler struct {}

func (sh *SetupHandler) Handle(ctx context.Context, runnerType string, payload *proto.TaskPayload) {
	var runner runners.Runner 
	runner, err := runners.GetRunner(runnerType)
	if err != nil {
		// TODO: send error
		return
	}
	if err := runner.Setup(ctx, payload.ExecutionInfraId, payload.InfraData); err != nil {
		// TODO: send error
		return
	}
}

type ExecuteHandler struct {}

func (sh *ExecuteHandler) Handle(ctx context.Context, runnerType string, payload *proto.TaskPayload) {
	var runner runners.Runner 
	runner, err := runners.GetRunner(runnerType)
	if err != nil {
		// TODO: send error
		return
	}
	runner.Execute(ctx, payload.ExecutionInfraId, payload.InfraData, payload.TaskData)
}

type CleanupHandler struct {}

func (sh *CleanupHandler) Handle(ctx context.Context, runnerType string, payload *proto.TaskPayload) {
	var runner runners.Runner 
	runner, err := runners.GetRunner(runnerType)
	if err != nil {
		// TODO: send error
		return
	}
	if err := runner.Cleanup(ctx, payload.ExecutionInfraId); err != nil {
		// TODO: send error
		return
	}
}