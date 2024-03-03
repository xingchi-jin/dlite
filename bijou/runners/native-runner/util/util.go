package util

import (
	"context"
	"sync"

	"github.com/wings-software/dlite/bijou/runners/native-runner/config"
	"github.com/sirupsen/logrus"
	"github.com/wings-software/dlite/client"
	"github.com/wings-software/dlite/delegate"
)

var onlyOnce sync.Once

// Register capacity of docker runner node 
func RegisterDelegateCapacity(Id string) {
	onlyOnce.Do(func() {
		c := config.GetConfig()
		cl := delegate.New(c.DelegateCapacity.ManagerEndpoint, c.DelegateCapacity.AccountID, c.DelegateCapacity.Secret, false)
		err := cl.RegisterCapacity(context.Background(), Id, &client.DelegateCapacity{MaxBuilds: c.DelegateCapacity.MaxBuilds})
		if err != nil {
			logrus.WithError(err).Errorln("failed to register delegate capacity")
		}
	})
}
