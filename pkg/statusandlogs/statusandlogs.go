package statusandlogs

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func GetStatus() (string, error) {
	cmd := exec.Command("mosdns", "version")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get MosDNS status: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}

func GetLogs(logFile string) (string, error) {
	data, err := ioutil.ReadFile(logFile)
	if err != nil {
		return "", fmt.Errorf("failed to read log file: %w", err)
	}

	return string(data), nil
}
