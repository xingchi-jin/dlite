// Copyright 2022 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package pipeline

import (
	"github.com/wings-software/dlite/bijou/runners/native-runner/api"
	"github.com/wings-software/dlite/bijou/runners/native-runner/engine/spec"
	"github.com/wings-software/dlite/bijou/runners/native-runner/logstream"
	"github.com/wings-software/dlite/bijou/runners/native-runner/logstream/filestore"
	"github.com/wings-software/dlite/bijou/runners/native-runner/logstream/remote"
	tiCfg "github.com/harness/lite-engine/ti/config"
)

const (
	SharedVolPath = "/tmp/engine"
	SharedVolName = "_engine"
)

// State stores the pipeline state.
type State struct {
	volumes   []*spec.Volume
	logConfig api.LogConfig
	tiConfig  tiCfg.Cfg
	secrets   []string
	logClient logstream.Client
	network   string
}

func NewState() *State {
	return &State{
		volumes:   make([]*spec.Volume, 0),
		logConfig: api.LogConfig{},
		tiConfig:  tiCfg.Cfg{},
		secrets:   make([]string, 0),
		logClient: nil,
	}
}

func (s *State) Set(volumes []*spec.Volume, secrets []string, logConfig api.LogConfig, tiConfig tiCfg.Cfg, network string) { // nolint:gocritic
	s.volumes = volumes
	s.secrets = secrets
	s.logConfig = logConfig
	s.tiConfig = tiConfig
	s.network = network
}

func (s *State) GetSecrets() []string {
	return s.secrets
}

func (s *State) GetVolumes() []*spec.Volume {
	return s.volumes
}
func (s *State) AppendSecrets(secrets []string) {
	s.secrets = append(s.secrets, secrets...)
}

func (s *State) GetLogStreamClient() logstream.Client {
	if s.logClient == nil {
		if s.logConfig.URL != "" {
			s.logClient = remote.NewHTTPClient(s.logConfig.URL, s.logConfig.AccountID,
				s.logConfig.Token, s.logConfig.IndirectUpload, false)
		} else {
			s.logClient = filestore.New(SharedVolPath)
		}
	}
	return s.logClient
}

func (s *State) GetTIConfig() *tiCfg.Cfg {
	return &s.tiConfig
}

func (s *State) GetNetwork() string {
	return s.network
}
