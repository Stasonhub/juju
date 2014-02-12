// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package upgrades

import (
	"fmt"
	"path"

	"launchpad.net/juju-core/utils/exec"
)

// stepsFor118 returns upgrade steps to upgrade to a Juju 1.18 deployment.
func stepsFor118(context Context) []UpgradeStep {
	return []UpgradeStep{
		// Nothing yet.
		&upgradeStep{
			description: `
				Previously the lock directory was created when the uniter started.
				This allows serialization of all of the hook execution across units
				running on a single machine.  This lock directory is now also used
				but the juju-run command on the host machine.  juju-run also gets
				a lock on the hook execution fslock prior to execution.  However,
				the lock directory was owned by root, and the juju-run process was
				being executed by the ubuntu user, so we need to change the ownership
				of the lock directory to ubuntu:ubuntu. Also we need to make sure that
				this directory exists on machines with no units.
			`,
			targets: []UpgradeTarget{HostMachine},
			run:     ensureLockDirExistsAndUbuntuWritable,
		},
	}
}

// Ensure that the lock dir exists and change the ownership of the lock dir
// itself to ubuntu:ubuntu from root:root so the juju-run command run as the
// ubuntu user is able to get access to the hook execution lock (like the
// uniter itself does.)
func ensureLockDirExistsAndUbuntuWritable(context Context) error {
	lockDir := path.Join(context.AgentConfig().DataDir(), "locks")
	// We only try to change ownership if there is an ubuntu user
	// defined, and we determine this by the existance of the home dir.
	command := fmt.Sprintf(""+
		"mkdir -p %s\n"+
		"[ -e /home/ubuntu ] && chown ubuntu:ubuntu %s",
		lockDir, lockDir)
	result, err := exec.RunCommands(exec.RunParams{
		Commands: command,
	})
	if err != nil {
		return err
	}
	if result.Code != 0 {
		return fmt.Errorf("failed to create lock dir or ownership change: \n%s\n%s", result.Stdout, result.Stderr)
	}
	return nil
}
