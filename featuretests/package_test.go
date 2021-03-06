// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package featuretests

import (
	"flag"
	stdtesting "testing"

	gc "gopkg.in/check.v1"

	coretesting "github.com/juju/juju/testing"
)

var runFeatureTests = flag.Bool("featuretests", true, "Run long-running feature tests.")

func init() {

	flag.Parse()

	if *runFeatureTests == false {
		return
	}
	// Initialize all suites here.
	gc.Suite(&cmdJujuSuite{})
	gc.Suite(&annotationsSuite{})
	gc.Suite(&apiEnvironmentSuite{})
	gc.Suite(&blockSuite{})
	gc.Suite(&cmdModelSuite{})
	gc.Suite(&cmdStorageSuite{})
	gc.Suite(&cmdControllerSuite{})
	gc.Suite(&dblogSuite{})
	gc.Suite(&cloudImageMetadataSuite{})
	gc.Suite(&cmdSpaceSuite{})
	gc.Suite(&cmdSubnetSuite{})
	gc.Suite(&undertakerSuite{})
	gc.Suite(&dumpLogsCommandSuite{})
	gc.Suite(&upgradeSuite{})
	gc.Suite(&cmdRegistrationSuite{})
	gc.Suite(&cmdLoginSuite{})
	gc.Suite(&BakeryStorageSuite{})
}

func TestPackage(t *stdtesting.T) {
	coretesting.MgoTestPackage(t)
}
