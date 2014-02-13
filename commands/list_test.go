package commands

import (
	"bytes"
	"os/exec"
	"testing"
)

func Test_pipedCommands(t *testing.T) {
	echo := exec.Command("echo", "foo", "bar")
	xargs := exec.Command("xargs")
	cmds, _ := pipedCommands(echo, xargs)

	expect := []byte("foo bar\n")
	output := cmds
	if !bytes.Equal(expect, output) {
		t.Errorf("pipedCommands:\n\texpected: %+v\n\t     got: %+v", expect, output)
	}
}
