package command

import (
	"fmt"
	"time"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/executor/deployment"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/factory"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	"github.com/urfave/cli"
)

type DeploymentBackupCleanupCommand struct {
}

func NewDeploymentBackupCleanupCommand() DeploymentBackupCleanupCommand {
	return DeploymentBackupCleanupCommand{}
}

func (d DeploymentBackupCleanupCommand) Cli() cli.Command {
	return cli.Command{
		Name:   "backup-cleanup",
		Usage:  "Cleanup a deployment after a backup was interrupted",
		Action: d.Action,
	}
}

func (d DeploymentBackupCleanupCommand) Action(c *cli.Context) error {
	trapSigint(true)

	username, password, target, caCert, bbrVersion, debug, deployment, allDeployments, maxInFlightThreads := getDeploymentParams(c)

	if !allDeployments {
		logger := factory.BuildBoshLogger(debug)

		cleaner, err := factory.BuildDeploymentBackupCleanuper(
			target,
			username,
			password,
			caCert,
			c.App.Version,
			maxInFlightThreads,
			logger,
		)
		if err != nil {
			return processError(orchestrator.NewError(err))
		}

		cleanupErr := cleaner.Cleanup(deployment)
		return processError(cleanupErr)
	}

	return cleanupAllDeployments(target, username, password, caCert, bbrVersion, debug, maxInFlightThreads)
}

func cleanupAllDeployments(target, username, password, caCert, bbrVersion string, debug bool, maxInFlightThreads int) error {
	cleanupAction := func(deploymentName string) orchestrator.Error {
		timestamp := time.Now().UTC().Format(artifactTimeStampFormat)
		logFilePath, buffer, logger, logErr := createLogger(timestamp, "", deploymentName, debug)
		if logErr != nil {
			return orchestrator.NewError(logErr)
		}

		cleaner, factoryError := factory.BuildDeploymentBackupCleanuper(
			target,
			username,
			password,
			caCert,
			bbrVersion,
			maxInFlightThreads,
			logger,
		)

		if factoryError != nil {
			return orchestrator.NewError(factoryError)
		}

		printlnWithTimestamp(fmt.Sprintf("Starting cleanup of %s, log file: %s", deploymentName, logFilePath))
		err := cleanup(cleaner, deploymentName)

		if err != nil {
			printlnWithTimestamp(fmt.Sprintf("ERROR: failed to cleanup %s", deploymentName))
			fmt.Println(buffer.String())
		} else {
			printlnWithTimestamp(fmt.Sprintf("Finished cleanup of %s", deploymentName))
		}

		return err
	}

	errorHandler := func(deploymentError deployment.AllDeploymentsError) error {
		return deploymentError.Process()
	}

	logger, _ := factory.BuildBoshLoggerWithCustomBuffer(debug)

	boshClient, err := factory.BuildBoshClient(target, username, password, caCert, bbrVersion, logger)
	if err != nil {
		return err
	}

	fmt.Println("Starting cleanup...")

	return runForAllDeployments(
		cleanupAction,
		boshClient,
		"could not be cleaned up",
		"cleaned up",
		errorHandler,
		deployment.NewParallelExecutor(maxInFlightThreads))
}

func cleanup(cleaner *orchestrator.BackupCleaner, deployment string) orchestrator.Error {
	err := cleaner.Cleanup(deployment)
	if err != nil {
		fmt.Printf("Failed to cleanup deployment '%s'\n", deployment)
		return err
	}
	return nil
}
