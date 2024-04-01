package rulemanagement

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadRules(ruleFile string) ([]string, error) {
	file, err := os.Open(ruleFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open rule file: %w", err)
	}
	defer file.Close()

	var rules []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rule := strings.TrimSpace(scanner.Text())
		if rule != "" && !strings.HasPrefix(rule, "#") {
			rules = append(rules, rule)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read rules from file: %w", err)
	}

	return rules, nil
}

func WriteRules(ruleFile string, rules []string) error {
	data := strings.Join(rules, "\n") + "\n"
	err := ioutil.WriteFile(ruleFile, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("failed to write rules to file: %w", err)
	}

	return nil
}

func ValidateRule(rule string) error {
	// Implement rule validation logic here
	// ...
	return nil
}
