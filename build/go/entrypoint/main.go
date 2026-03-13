package main

import (
	"os"
	"syscall"

	"github.com/11notes/go-eleven"
)

const APP_BIN = "loki"
const APP_CONFIG_ENV string = "LOKI_CONFIG"
const APP_CONFIG string = "/loki/etc/default.yml"

func main() {
	// write env to file if set
	eleven.Container.EnvToFile(APP_CONFIG_ENV, APP_CONFIG)

	// start app
	if err := syscall.Exec("/usr/local/bin/" + APP_BIN, []string{APP_BIN, "--config.file", APP_CONFIG}, os.Environ()); err != nil {
		os.Exit(1)
	}
}