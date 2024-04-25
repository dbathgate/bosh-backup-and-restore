package factory

import (
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/bosh"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/executor"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orderer"
	"github.com/cloudfoundry/bosh-utils/logger"
)

func BuildDeploymentBackupCleanuper(
	target string,
	username string,
	password string,
	caCert string,
	bbrVersion string,
	maxInFlightThreads int,
	maxConnectionsPerMinute int,
	logger logger.Logger,
) (*orchestrator.BackupCleaner, error) {

	boshClient, err := BuildBoshClient(target, username, password, caCert, bbrVersion, maxConnectionsPerMinute, logger)

	if err != nil {
		return nil, err
	}

	return orchestrator.NewBackupCleaner(
		logger,
		bosh.NewDeploymentManager(boshClient, logger, false),
		orderer.NewKahnBackupLockOrderer(),
		executor.NewParallelExecutor(maxInFlightThreads),
	), nil
}
