package mij

import (
	"log"
	"os/exec"
	"strconv"
)

type dockerImage struct {
	name         string
	portExternal int
	portInternal int
}

func (d *dockerImage) setup() {
	cmdString := "docker run -d --rm -p " + strconv.Itoa(d.portExternal) + ":" + strconv.Itoa(d.portInternal) + " --name nexus sonatype/nexus3:3.16.1"
	cmd := exec.Command("bash", "-c", cmdString)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}

func (d *dockerImage) shutdown() {
	cmd := exec.Command("bash", "-c", "docker stop nexus")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}
