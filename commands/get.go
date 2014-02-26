package commands

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/craigmj/gototp"
)

func GetCode(name string) {
	keychainName := "mmfa_" + name
	secret, err := secret(keychainName)
	if err != nil {
		os.Stderr.Write([]byte("Error getting secret: " + err.Error() + "\n" + secret))
		os.Exit(1)
	}

	otp, err := gototp.New(secret)
	if err != nil {
		os.Stderr.Write([]byte("Error calculating code: " + err.Error()))
		os.Exit(1)
	}

	code := strconv.Itoa(int(otp.Now()))
	copyToClipboard(strings.Repeat("0", 6-len(code)) + code)
	os.Stdout.Write([]byte(strings.Repeat("0", 6-len(code)) + code + "\n"))
}

func secret(name string) (string, error) {
	command, args := getSecretCommand(name)
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return strings.TrimSpace(string(output)), nil
}

func getSecretCommand(name string) (string, []string) {
	args := []string{
		"find-generic-password",
		"-a",
		name,
		"-s",
		name,
		"-w",
	}
	return "security", args
}

func copyToClipboard(code string) {
	echo := exec.Command("echo", code)
	pbcopy := exec.Command("pbcopy")
	_, err := pipedCommands(echo, pbcopy)
	if err != nil {
		print(err)
	}
}
