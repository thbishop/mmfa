package commands

import (
	"strings"
	"testing"
)

func Test_getSecretCommand(t *testing.T) {
	cmd_name, args := getSecretCommand("foo")

	expect := "security find-generic-password -a foo -s foo -w"
	output := cmd_name + " " + strings.Join(args, " ")
	if expect != cmd_name+" "+strings.Join(args, " ") {
		t.Errorf("getSecretCommand:\n\texpected: %+v\n\t     got: %+v", expect, output)
	}
}
