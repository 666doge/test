package utils

import (
	"os"
	"go/build"
)

func GetGopath() (gopath string) {
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return
}