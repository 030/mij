package mij

import (
	"os/exec"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type DockerImage struct {
	ContainerName            string
	Name                     string
	PortInternal             int
	PortExternal             int
	Version                  string
	LogFile                  string
	LogFileStringHealthCheck string
	HealthcheckURL           string
	EnvironmentVariables     []string
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
	var healthcheck string
	if d.HealthcheckURL != "" {
	  healthcheck = "curl --fail "+d.HealthcheckURL
	} else {
	  healthcheck = "cat " + d.LogFile + " | grep \"" + d.LogFileStringHealthCheck + "\""
	}
	
	var envVars string
	var str strings.Builder
	if len(d.EnvironmentVariables)>0{
		for _,e:=range envVars{
		  str.WriteString(" -e "+e)
		}
		envVars = str.String()
	}
	
	if err := bash("docker run -d --rm " +envVars+ " -p " + strconv.Itoa(d.PortExternal) + ":" + strconv.Itoa(d.PortInternal) + " --name " + d.ContainerName + " --health-interval 5s --health-retries 10 --health-cmd='"+healthcheck+"' " + d.Name + ":" + d.Version); err != nil {
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
