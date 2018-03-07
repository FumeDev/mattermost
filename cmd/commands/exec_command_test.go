// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package commands

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mattermost/mattermost-server/cmd"
)

func TestExecCommand(t *testing.T) {
	if filter := flag.Lookup("test.run").Value.String(); filter != "ExecCommand" {
		t.Skip("use -run ExecCommand to execute a command via the test executable")
	}
	cmd.RootCmd.SetArgs(flag.Args())
	require.NoError(t, cmd.RootCmd.Execute())
}
