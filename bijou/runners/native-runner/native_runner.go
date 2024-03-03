package nativerunner

import (
	"context"
	"github.com/wings-software/dlite/client/proto"
	pb "google.golang.org/protobuf/proto"
)

type NativeRunner struct {}

// Setup(ctx context.Context, infraId string, infraData *proto.InputData) error
// Execute(ctx context.Context, infraId string, infraData, taskData *proto.InputData)
// Cleanup(ctx context.Context, infraId string) error

func (runner *NativeRunner) Setup(ctx context.Context, infraId string, infraData *proto.InputData) error {
	data := infraData.GetBinaryData()
	nativeInfra := &proto.NativeInfra{}
	pb.Unmarshal(data, nativeInfra)
	
}

func (runner *NativeRunner) Execute(ctx context.Context, infraId string, infraData, taskData *proto.InputData) {

}

func (runner *NativeRunner) Cleanup(ctx context.Context, infraId string) error {

}