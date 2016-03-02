package commands

import (
	"strings"
	"testing"
)

func Test_addItemCommand(t *testing.T) {
	cmd_name, args := addItemCommand("foo", "comments", "secret")

	expect := "security add-generic-password -a foo -s foo -j comments -w secret"
	output := cmd_name + " " + strings.Join(args, " ")
	if expect != output {
		t.Errorf("addItemCommand:\n\texpected: %+v\n\t     got: %+v", expect, output)
	}
}
