package handlers

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wings-software/dlite/bijou/runners"
	"github.com/wings-software/dlite/client"
	"github.com/wings-software/dlite/client/proto"
)

type Handler interface {
	Handle(ctx context.Context, runnerType string, payload *proto.TaskPayload)
}

type SetupHandler struct {
	client client.Client
}

func (sh *SetupHandler) Handle(ctx context.Context, runnerType string, payload *proto.TaskPayload) {
	var runner runners.Runner
	runner, err := runners.GetRunner(runnerType)
	delegateId := ctx.Value("delegateId").(string)
	delegateName := ctx.Value("delegateName").(string)
	if err != nil {
		if http_err := sh.client.SendSetupResponse(ctx, delegateId, payload.ExecutionInfraId, payload.Id, &proto.SetupInfraResponse{
			ResponseCode: proto.ResponseCode_RESPONSE_FAILED,
			ErrorMessage: err.Error(),
			Location: &proto.ExecutionInfraInfo{
				DelegateName: delegateName,
				RunnerType:   runnerType,
			},
		}); err != nil {
			logrus.WithError(http_err).Error("Send setup failure response failed")
		}
		return
	}
	if err := runner.Setup(ctx, payload.ExecutionInfraId, payload.InfraData); err != nil {
		if http_err := sh.client.SendSetupResponse(ctx, delegateId, payload.ExecutionInfraId, payload.Id, &proto.SetupInfraResponse{
			ResponseCode: proto.ResponseCode_RESPONSE_FAILED,
			ErrorMessage: err.Error(),
			Location: &proto.ExecutionInfraInfo{
				DelegateName: delegateName,
				RunnerType:   runnerType,
			},
		}); err != nil {
			logrus.WithError(http_err).Error("Send setup failure response failed")
		}
		return
	} else {
		if http_err := sh.client.SendSetupResponse(ctx, delegateId, payload.ExecutionInfraId, payload.Id, &proto.SetupInfraResponse{
			ResponseCode: proto.ResponseCode(proto.StatusCode_CODE_SUCCESS),
			Location: &proto.ExecutionInfraInfo{
				DelegateName: delegateName,
				RunnerType:   runnerType,
			},
		}); err != nil {
			logrus.WithError(http_err).Error("Send setup failure response failed")
		}
		return
	}
}

type ExecuteHandler struct {
	client client.Client
}

func (sh *ExecuteHandler) Handle(ctx context.Context, runnerType string, payload *proto.TaskPayload) {
	var runner runners.Runner
	runner, err := runners.GetRunner(runnerType)
	delegateId := ctx.Value("delegateId").(string)
	if err != nil {
		if http_err := sh.client.SendExecutionResponse(ctx, delegateId, payload.Id, &proto.ExecutionStatusResponse{
			Status: &proto.ExecutionStatus{
				Id:    payload.Id,
				Code:  proto.StatusCode_CODE_FAILED,
				Error: err.Error(),
			},
		}); err != nil {
			logrus.WithError(http_err).Error("Send execution failure response failed")
		}
		return
	}
	runner.Execute(ctx, payload.ExecutionInfraId, payload.InfraData, payload.TaskData)
}

type CleanupHandler struct {
	client client.Client
}

func (sh *CleanupHandler) Handle(ctx context.Context, runnerType string, payload *proto.TaskPayload) {
	var runner runners.Runner
	runner, err := runners.GetRunner(runnerType)
	delegateId := ctx.Value("delegateId").(string)
	if err != nil {
		if http_err := sh.client.SendCleanupResponse(ctx, delegateId, payload.ExecutionInfraId, payload.Id, &proto.CleanupInfraResponse{
			ResponseCode: proto.ResponseCode_RESPONSE_FAILED,
			ErrorMessage: err.Error(),
		}); err != nil {
			logrus.WithError(http_err).Error("Send cleanup failure response failed")
		}
		return
	}
	if err := runner.Cleanup(ctx, payload.ExecutionInfraId); err != nil {
		if http_err := sh.client.SendCleanupResponse(ctx, delegateId, payload.ExecutionInfraId, payload.Id, &proto.CleanupInfraResponse{
			ResponseCode: proto.ResponseCode_RESPONSE_FAILED,
			ErrorMessage: err.Error(),
		}); err != nil {
			logrus.WithError(http_err).Error("Send cleanup failure response failed")
		}
		return
	} else {
		if http_err := sh.client.SendCleanupResponse(ctx, delegateId, payload.ExecutionInfraId, payload.Id, &proto.CleanupInfraResponse{
			ResponseCode: proto.ResponseCode_RESPONSE_OK,
		}); err != nil {
			logrus.WithError(http_err).Error("Send cleanup success response failed")
		}
	}
}
