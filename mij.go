package mij

import (
	"os/exec"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type DockerImage struct {
	Name         string
	PortExternal int
	PortInternal int
	Version      string
}

func (d *DockerImage) Setup() {
	cmdString := "docker run -d --rm -p " + strconv.Itoa(d.PortExternal) + ":" + strconv.Itoa(d.PortInternal) + " --name " + d.Name + " " + d.Name + ":" + d.Version
	cmd := exec.Command("bash", "-c", cmdString)
	stdoutStderr, err := cmd.CombinedOutput()
	log.Info(stdoutStderr)
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}

func (d *DockerImage) Shutdown() {
	cmd := exec.Command("bash", "-c", "docker stop "+d.Name)
	stdoutStderr, err := cmd.CombinedOutput()
	log.Info(stdoutStderr)
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}
