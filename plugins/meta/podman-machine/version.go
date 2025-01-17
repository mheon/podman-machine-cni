package main

import "fmt"

// overwritten at build time
var gitCommit = "unknown"

const machineVersion = "0.1.0-dev"

func getVersion() string {
	return fmt.Sprintf(`CNI podman-machine plugin
version: %s
commit: %s`, machineVersion, gitCommit)
}
