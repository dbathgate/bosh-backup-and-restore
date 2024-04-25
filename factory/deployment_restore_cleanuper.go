package factory

import (
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/bosh"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/executor"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orderer"
)

func BuildDeploymentRestoreCleanuper(target,
	usename,
	password,
	caCert,
	bbrVersion string,
	withManifest,
	isDebug bool,
	maxConnectionsPerMinute int) (*orchestrator.RestoreCleaner, error) {

	logger := BuildLogger(isDebug)

	boshClient, err := BuildBoshClient(
		target,
		usename,
		password,
		caCert,
		bbrVersion,
		maxConnectionsPerMinute,
		logger,
	)

	if err != nil {
		return nil, err
	}

	return orchestrator.NewRestoreCleaner(logger,
		bosh.NewDeploymentManager(boshClient, logger, withManifest), orderer.NewKahnRestoreLockOrderer(), executor.NewSerialExecutor()), nil
}
