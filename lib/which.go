package lib

import "os/exec"

func which(bin string) string {
	path, _ := exec.LookPath(bin)
	return path
}
