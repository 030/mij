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
	Version      string
}

func (d *DockerImage) Setup() {
	cmdString := "docker run -d --rm -p " + strconv.Itoa(d.PortExternal) + ":" + strconv.Itoa(d.PortInternal) + " --name " + Name + " " + Name + ":" + Version
	cmd := exec.Command("bash", "-c", cmdString)
	stdoutStderr, err := cmd.CombinedOutput()
	log.Info(stdoutStderr)
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}

func (d *DockerImage) Shutdown() {
	cmd := exec.Command("bash", "-c", "docker stop "+Name)
	stdoutStderr, err := cmd.CombinedOutput()
	log.Info(stdoutStderr)
	if err != nil {
		log.Fatal(err, string(stdoutStderr))
	}
}
