package aireqlog

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSystem(t *testing.T) {
	system, err := getSystemIdentity()
	require.NoError(t, err)

	if testing.Verbose() {
		t.Logf("%#v", system)
	}

	require.NotEmpty(t, system.Hostname)
	require.Contains(t, system.OperatingSystem, runtime.GOOS)
	require.Contains(t, system.OperatingSystem, runtime.GOARCH)
	require.NotEmpty(t, system.MACAddress)
	require.NotEmpty(t, system.Username)
}
