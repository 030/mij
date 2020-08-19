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
	URI           string
}

func bash(cmd string) error {
	log.Info(cmd)
	b, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	log.Info(string(b))
	if err != nil {
		return err
	}
	return nil
}

func (d *DockerImage) Run() {
	if err := bash("docker run -d --rm -p " + strconv.Itoa(d.PortExternal) + ":" + strconv.Itoa(d.PortInternal) + " --name " + d.ContainerName + " --health-interval 5s --health-retries 10 --health-cmd='curl --fail http://localhost:" + strconv.Itoa(d.PortInternal) + d.URI + "' " + d.Name + ":" + d.Version); err != nil {
		log.Fatal(err)
	}

	for bash("docker ps -f name="+d.ContainerName+" -f health=healthy | grep "+d.ContainerName) != nil {
		log.Warn("Docker container unhealthy")
	}
	log.Info("Docker container healthy")
}

func (d *DockerImage) Stop() {
	if err := bash("docker stop " + d.ContainerName); err != nil {
		log.Fatal(err)
	}
}
