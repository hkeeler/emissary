// Ambassador combines the various Golang binaries used in the Ambassador
// container, dispatching on os.Args[0] like BusyBox. Note that the
// capabilities_wrapper binary is _not_ included here. That one has special
// permissions magic applied to it that is not appropriate for these other
// binaries.
package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/datawire/ambassador/v2/pkg/busy"
	"github.com/datawire/ambassador/v2/pkg/environment"

	amb_agent "github.com/datawire/ambassador/v2/cmd/agent"
	"github.com/datawire/ambassador/v2/cmd/ambex"
	"github.com/datawire/ambassador/v2/cmd/apiext"
	"github.com/datawire/ambassador/v2/cmd/entrypoint"
	"github.com/datawire/ambassador/v2/cmd/kubestatus"
	"github.com/datawire/ambassador/v2/cmd/reproducer"
)

// Version is set at run-time by reading the 'ambassador.version' file.  We do this instead of
// compiling in a version so that we can promote RC images to GA without recompiling anything.
var Version = "MISSING(FILE)"

func noop(_ context.Context) {}

// Builtin for showing this image's version.
func showVersion(ctx context.Context, Version string, args ...string) error {
	fmt.Printf("Version %s\n", Version)

	return nil
}

func main() {
	// Keep this in-sync with VERSION.py.
	if verBytes, err := os.ReadFile("/buildroot/ambassador/python/ambassador.version"); err == nil {
		verLines := strings.Split(string(verBytes), "\n")
		for len(verLines) < 2 {
			verLines = append(verLines, "MISSING(VAL)")
		}
		Version = verLines[0]
	}

	busy.Main("busyambassador", "Ambassador", Version, map[string]busy.Command{
		"ambex":      {Setup: environment.EnvironmentSetupEntrypoint, Run: ambex.Main},
		"kubestatus": {Setup: environment.EnvironmentSetupEntrypoint, Run: kubestatus.Main},
		"entrypoint": {Setup: noop, Run: entrypoint.Main},
		"reproducer": {Setup: noop, Run: reproducer.Main},
		"agent":      {Setup: environment.EnvironmentSetupEntrypoint, Run: amb_agent.Main},
		"version":    {Setup: noop, Run: showVersion},
		"apiext":     {Setup: noop, Run: apiext.Main},
	})
}
