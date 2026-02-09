package main

import (
	"os"
	"syscall"

	"github.com/11notes/docker-util"
)

const ENV_CONFIG = "LOKI_CONFIG"
const ROOT_CONFIG = "/loki/etc/default.yml"

func main() {
	if _, ok := os.LookupEnv(ENV_CONFIG); ok {
		eleven.Container.EnvToFile(ENV_CONFIG, ROOT_CONFIG)
	}
	if err := syscall.Exec("/usr/local/bin/loki", []string{"loki", "--config.file", ROOT_CONFIG}, os.Environ()); err != nil {
		os.Exit(1)
	}
}