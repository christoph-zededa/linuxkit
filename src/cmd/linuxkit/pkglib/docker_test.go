package pkglib

import (
	"fmt"
	"testing"
)

func compareArgs(t *testing.T, args, argsExpected []string) {
	if len(args) != len(argsExpected) {
		t.Fatalf("args are: %+v, expected: %+v", args, argsExpected)
	}

	for i := 0; i < len(args); i++ {
		if args[i] != argsExpected[i] {
			t.Fatalf("args are: %+v, expected: %+v", args, argsExpected)
		}
	}
}

func TestCreateBuildkitContainerArgsDefault(t *testing.T) {
	argsDefaultAddr := []string{
		"container",
		"run",
		"-d",
		"--name",
		"name",
		"--privileged",
		"image",
		"--allow-insecure-entitlement",
		"network.host",
		"--debug",
		"--addr",
		fmt.Sprintf("unix://%s", buildkitSocketPath),
	}

	args, err := createBuildkitContainerArgs("name", "image", "")
	if err != nil {
		t.Fatalf("createBuildkitContainerArgs failed: %s", err)
	}

	compareArgs(t, args, argsDefaultAddr)
}

func TestCreateBuildkitContainerArgsTcpSocket(t *testing.T) {
	addr := "tcp://0.0.0.0:1234"

	argsWithTcpAddr := []string{
		"container",
		"run",
		"-d",
		"-p",
		"1234:1234",
		"--name",
		"name",
		"--privileged",
		"image",
		"--allow-insecure-entitlement",
		"network.host",
		"--debug",
		"--addr",
		addr,
	}

	args, err := createBuildkitContainerArgs("name", "image", addr)
	if err != nil {
		t.Fatalf("createBuildkitContainerArgs failed: %s", err)
	}

	compareArgs(t, args, argsWithTcpAddr)
}

func TestCreateBuildkitContainerArgsInvalidTcpSocket(t *testing.T) {
	addr := "tcp://0.0.0.0"
	_, err := createBuildkitContainerArgs("name", "image", addr)
	if err == nil {
		t.Fatalf("expected error, but got nil for invalid addr %s", addr)
	}
}
