// A script that wraps the run of the project test suite and starts stripe-mock
// with a custom OpenAPI + fixtures bundle if one was found in the appropriate
// spot (see `pathSpec` below).
//
// The script passes all its arguments to a Go's test command, and defaults to
// `./...`. For example, both of the following are valid invocations:
//
//	go run test_with_stripe_mock/main.go
//	go run test_with_stripe_mock/main.go ./charge
//
// The reason that we need a separate script for this is because Go's testing
// infrastructure doesn't provide any kind of global hook that we can use to do
// this work in only one place before any tests are run.
package main

import (
	"fmt"
	"os"
	"regexp"
)

const (
	defaultStripeMockPort = "12112"

	pathSpec     = "./testing/openapi/spec3.json"
	pathFixtures = "./testing/openapi/fixtures3.json"
)

var portMatch = regexp.MustCompile(` port: (\d+)`)

func main() {
	var stripeMockProcess *os.Process
	autostart := true

	port := os.Getenv("STRIPE_MOCK_PORT")
	if port == "" {
		port = defaultStripeMockPort
	}

	//
	// Maybe start stripe-mock
	//

	autostartEnv := os.Getenv("STRIPE_MOCK_AUTOSTART")
	if autostartEnv == "0" || autostartEnv == "false" {
		fmt.Printf("STRIPE_MOCK_AUTOSTART=%s, "+
			"assuming stripe-mock is already running on port %s\n", autostartEnv, port)
		autostart = false
	} else if _, err := os.Stat(pathSpec); os.IsNotExist(err) {
		fmt.Printf("No custom spec file found, "+
			"assuming stripe-mock is already running on port %s\n", port)
		autostart = false
	}

	if autostart {
		var err error
		port, stripeMockProcess, err = startStripeMock()
		if err != nil {
			exitWithError(err)
		}
	}

	//
	// Run tests
	//

	err := runTests(port)
	if err != nil {
		stopStripeMock(stripeMockProcess)
		exitWithError(err)
	}

	//
	// Stop stripe-mock
	//

	// Try to cleanly stop stripe-mock, but in case we don't, it'll die anyway
	// because it's executing as a subprocess.
	stopStripeMock(stripeMockProcess)
}

//
// Private functions
//

func exitWithError(err error) { _ = "STUB: not implemented"; return }

func runTests(port string) error { _ = "STUB: not implemented"; return nil }

// Defaults to `./...`, but also allows a specific package (or other CLI
// flags like `-test.v`) to be passed.

// Inherit this script's environment so that it's still possible to pass
// the test package flags like `GOCACHE=off`.

func startStripeMock() (string, *os.Process, error) { _ = "STUB: not implemented"; return "", nil, nil }

// stripe-mock will select a port

// We store the entire captured output because the string we're looking for
// may have appeared across a read boundary.

// Look for port in "Listening for HTTP on port: 50602"

func stopStripeMock(process *os.Process) { _ = "STUB: not implemented"; return }
