package envs

import (
	"os"
	"strings"
)

const (
	PATH_NAME = "PATH"
)

type EnvVar = string

var (
	PATH_ENV EnvVar
)

func init() {
	PATH_ENV = os.Getenv(PATH_NAME)
}


func PathEnvProcess(path string) []string {
	return strings.Split(path, ":")
}
