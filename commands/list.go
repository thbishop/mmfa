package commands

import (
	"os"
	"os/exec"
)

func ListKeychainItems() {
	security := exec.Command("security", "dump-keychain")
	grep := exec.Command("grep", "-E", "acct.*mmfa_")
	cut := exec.Command("cut", "-d", "=", "-f2")
	sed := exec.Command("sed", "s/\"//g;s/mmfa_//")
	sort := exec.Command("sort")
	output, err := pipedCommands(security, grep, cut, sed, sort)
	if err != nil {
		print(err)
	}

	os.Stdout.Write([]byte(output))
}
