package aireqlog

import (
	"fmt"
	"net/url"
	"os/exec"
	"path"
	"strings"
)

func projectDetails() (Project, error) {
	var out Project

	name, err := projectName()
	if err != nil {
		return out, err
	}
	out.Name = name

	head, err := headCommit()
	if err != nil {
		return out, err
	}
	out.GitCommit = head

	return out, nil
}

func projectName() (string, error) {
	out, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("getting repository name: %w", err)
	}

	input := strings.TrimSpace(string(out))
	input = strings.ReplaceAll(input, ":", "/")

	u, err := url.Parse(input)
	if err != nil {
		return "", fmt.Errorf("parsing repository name: %w", err)
	}

	return strings.TrimSuffix(path.Base(u.Path), ".git"), nil
}

func headCommit() (string, error) {
	out, err := exec.Command("git", "rev-parse", "HEAD").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("getting HEAD: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}
