package aireqlog

import (
	"fmt"
	"time"
)

const (
	BetaVersion = "v0.1"
)

type RequestLog struct {
	Version        string         `json:"version"`
	Project        Project        `json:"project"`
	File           FileInfo       `json:"file"`
	RequestDetails RequestDetails `json:"request_details"`
	SystemIdentity SystemIdentity `json:"system_identity"`
	Timestamp      time.Time      `json:"timestamp"`
}

type Project struct {
	// Name is the full name of the project where data was sent from
	Name string `json:"name"`

	// GitCommit is the latest commit of the project
	GitCommit string `json:"git_commit"`
}

type FileInfo struct {
	Filename string `json:"filename"`
}

type RequestDetails struct {
	AIEngine  string `json:"ai_engine"`
	TargetURL string `json:"target_url"`
}

type SystemIdentity struct {
	Username        string `json:"username"`
	Hostname        string `json:"hostname"`
	OperatingSystem string `json:"operating_system"`
	MACAddress      string `json:"mac_address"`
}

func FromRequest(req RequestDetails, file FileInfo) (RequestLog, error) {
	var out RequestLog

	var err error
	if out, err = baseRequest(); err != nil {
		return out, err
	}

	out.File = file
	out.RequestDetails = req

	return out, nil
}

func baseRequest() (RequestLog, error) {
	var out RequestLog
	out.Version = BetaVersion
	out.Timestamp = time.Now().In(time.UTC)

	var err error
	out.Project, err = projectDetails()
	if err != nil {
		return out, fmt.Errorf("project: %w", err)
	}

	out.SystemIdentity, err = getSystemIdentity()
	if err != nil {
		return out, fmt.Errorf("system: %w", err)
	}

	return out, nil
}
