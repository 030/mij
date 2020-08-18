package mij

import (
	"os"
	"os/exec"
	"strconv"
	"testing"

	log "github.com/sirupsen/logrus"
)

type dockerImage struct {
	name         string
	portExternal int
	portInternal int
}

// https://stackoverflow.com/a/34102842/2777965
func (d *dockerImage) TestMain(m *testing.M) {
	d.setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func (d *dockerImage) setup() {
	cmdString := "docker run -d --rm -p " + strconv.Itoa(d.portExternal) + ":" + strconv.Itoa(d.portInternal) + " --name nexus sonatype/nexus3:3.16.1"
	cmd := exec.Command("bash", "-c", cmdString)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}

func shutdown() {
	cmd := exec.Command("bash", "-c", "docker stop nexus")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}
