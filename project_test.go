package aireqlog

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProject(t *testing.T) {
	proj, err := projectDetails()
	require.NoError(t, err)

	if testing.Verbose() {
		t.Logf("%#v", proj)
	}

	require.Equal(t, "aireqlog", proj.Name)
	require.NotEmpty(t, proj.GitCommit)
}
