package aireqlog

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromRequest(t *testing.T) {
	req := RequestDetails{
		AIEngine:  "openai/gpt-4o",
		TargetURL: "http://api.openai.com",
	}
	file := FileInfo{
		Filename: "api_users.go",
	}

	log, err := FromRequest(req, file)
	require.NoError(t, err)

	if testing.Verbose() {
		t.Logf("%#v", log)
	}

	require.Equal(t, BetaVersion, log.Version)
	require.False(t, log.Timestamp.IsZero())

	require.Equal(t, "aireqlog", log.Project.Name)
	require.NotEmpty(t, log.Project.GitCommit)

	require.Equal(t, file, log.File)
	require.Equal(t, req, log.RequestDetails)

	require.NotEmpty(t, log.SystemIdentity.Hostname)
	require.Contains(t, log.SystemIdentity.OperatingSystem, runtime.GOOS)
	require.Contains(t, log.SystemIdentity.OperatingSystem, runtime.GOARCH)
	require.NotEmpty(t, log.SystemIdentity.MACAddress)
	require.NotEmpty(t, log.SystemIdentity.Username)
}
