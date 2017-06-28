package director

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	. "github.com/pivotal-cf/bosh-backup-and-restore/system"

	"fmt"
	"testing"
	"time"
)

var workspaceDir string
var fixturesPath = "../../fixtures/director-backup/"

func TestDirector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Director Suite")
}

var _ = BeforeSuite(func() {
	SetDefaultEventuallyTimeout(4 * time.Minute)

	By("building bbr")
	commandPath, err := gexec.BuildWithEnvironment("github.com/pivotal-cf/bosh-backup-and-restore/cmd/bbr", []string{"GOOS=linux", "GOARCH=amd64"})
	Expect(err).NotTo(HaveOccurred())

	workspaceDir = fmt.Sprintf("/var/vcap/store/pre_backup_check_workspace-%d", time.Now().Unix())

	By("setting up the jump box")
	Eventually(JumpboxDeployment().RunCommand("jumpbox", "0",
		fmt.Sprintf("sudo mkdir %s && sudo chown vcap:vcap %s && sudo chmod 0777 %s", workspaceDir, workspaceDir, workspaceDir),
	)).Should(gexec.Exit(0))

	JumpboxDeployment().Copy("jumpbox", "0", commandPath, workspaceDir)
	JumpboxDeployment().Copy("jumpbox", "0", MustHaveEnv("SSH_KEY"), workspaceDir+"/key.pem")

	Eventually(JumpboxDeployment().RunCommand("jumpbox", "0",
		fmt.Sprintf("sudo chown -R vcap:vcap %s", workspaceDir),
	)).Should(gexec.Exit(0))
})

var _ = AfterSuite(func() {
	By("cleaning up the jump box")
	Eventually(JumpboxDeployment().RunCommand("jumpbox", "0",
		fmt.Sprintf("sudo rm -rf %s", workspaceDir),
	)).Should(gexec.Exit(0))
})
