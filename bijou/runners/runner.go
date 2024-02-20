package runners

import (
	"context"

	"github.com/wings-software/dlite/client/proto"
)

type Runner interface {
	Setup(ctx context.Context, infraId string, infraData *proto.InputData) error
	Execute(ctx context.Context, infraId string, infraData, taskData *proto.InputData)
	Cleanup(ctx context.Context, infraId string) error
}