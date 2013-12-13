package commands

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/craigmj/gototp"
)

func GetCode(name string) {
	secret, err := secret(name)
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
	print(strings.Repeat("0", 6-len(code)) + code)
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
