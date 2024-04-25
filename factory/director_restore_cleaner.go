package factory

import (
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/executor"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/instance"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orderer"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/ssh"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/standalone"
)

func BuildDirectorRestoreCleaner(host,
	username,
	privateKeyPath,
	bbrVersion string,
	hasDebug bool,
	maxConnectionsPerMinute int) *orchestrator.RestoreCleaner {

	logger := BuildLogger(hasDebug)

	deploymentManager := standalone.NewDeploymentManager(logger,
		host,
		username,
		privateKeyPath,
		instance.NewJobFinderOmitMetadataReleases(bbrVersion, logger),
		ssh.NewSshRemoteRunner,
		maxConnectionsPerMinute,
	)

	return orchestrator.NewRestoreCleaner(logger, deploymentManager, orderer.NewKahnRestoreLockOrderer(), executor.NewSerialExecutor())
}
