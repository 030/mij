package mij

import (
	"log"
	"os/exec"
	"strconv"
)

type DockerImage struct {
	Name         string
	PortExternal int
	PortInternal int
}

func (d *DockerImage) setup() {
	cmdString := "docker run -d --rm -p " + strconv.Itoa(d.PortExternal) + ":" + strconv.Itoa(d.PortInternal) + " --name nexus sonatype/nexus3:3.16.1"
	cmd := exec.Command("bash", "-c", cmdString)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}

func (d *DockerImage) shutdown() {
	cmd := exec.Command("bash", "-c", "docker stop nexus")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}
