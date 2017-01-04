package exec_test

import (
	"testing"

	"github.com/flyaways/storage/agent/storage/posix/exec"
)

func TestExecRight(t *testing.T) {
	exec.Exec("/bin/sh", "-c", " ping 127.0.0.1 -c 1")
	t.Log("TestExecRight!")
}

func TestExecInvalidIP(t *testing.T) {
	exec.Exec("/bin/sh", "-c", " ping 8.8.8.8.8 -c 1")
	t.Log("TestExecValidIP!")
}

func TestExecNoneC(t *testing.T) {
	exec.Exec("/bin/sh", "ping 127.0.0.1 -c 1")
	t.Log("TestExecNoneC!")
}

func TestExecNoneCMD(t *testing.T) {
	exec.Exec("/bin/sh")
	t.Log("TestExecNoneCMD!")
}

func TestExecNoneName(t *testing.T) {
	exec.Exec("ping 127.0.0.1 -c 1")
	t.Log("TestExecNoneName!")
}

func TestExecInvalidName(t *testing.T) {
	exec.Exec("/bin/bash  ping 127.0.0.1 -c 1")
	t.Log("TestExecInvalidName!")
}

func TestExecNone(t *testing.T) {
	exec.Exec("")
	t.Log("TestExecNone!")
}
