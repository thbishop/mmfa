package commands

import (
	"os"
	"os/exec"
)

func ListKeychainItems() {
	security := exec.Command("security", "dump-keychain")
	seda := exec.Command("sed", "-n", "s/acct.*mmfa_//p;s/icmt.*mmfa_//p")
	sedb := exec.Command("sed", "s/\"//g")
	paste := exec.Command("paste", "-d", " ", "-", "-")
	sort := exec.Command("sort")

	output, err := pipedCommands(security, seda, sedb, paste, sort)
	if err != nil {
		print(err)
	}

	os.Stdout.Write([]byte(output))
}
