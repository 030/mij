package mij

import (
	"os/exec"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type DockerImage struct {
	ContainerName string
	Name          string
	PortInternal  int
	PortExternal  int
	Version       string
}

func bash(cmd string) {
	b, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	log.Info(string(b))
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DockerImage) Run() {
	bash("docker run -d --rm -p " + strconv.Itoa(d.PortExternal) + ":" + strconv.Itoa(d.PortInternal) + " --name " + d.ContainerName + " " + d.Name + ":" + d.Version)
}

func (d *DockerImage) Stop() {
	bash("docker stop " + d.ContainerName)
}
