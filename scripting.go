/* Sundry Go scripting utilities - these wrap error handling, and panic on error. */
/* No warranty! */

package scripting

import (
	"io/ioutil"
	"log"
	"os/exec"
)

// Replace path extension `origExt` with new extension `newExt`.
// Panic if `path` does not have extension `origExt`.
func ReplaceExt(path, origExt, newExt string) string {
	if len(path) < len(origExt) || path[len(path)-len(origExt):] != origExt {
		log.Panicf("Path %s does not have extension %s", path, origExt)
	}
	return path[:len(path)-len(origExt)] + newExt
}

// Execute the `command` with the given `arguments`.  Panic if the
// command fails.
func Exec(command string, arguments []string) {
	cmd := exec.Command(command, arguments...)
	err := cmd.Run()
	Try(err)
}

// Read the contents of file named by `filename` as a byte slice.
// Panic if the command fails.
func ReadFile(filename string) []byte {
	input, err := ioutil.ReadFile(filename)
	Try(err)
	return input
}

// If `err` is not nil, panic with that error.
func Try(err error) {
	if err != nil {
		log.Panic(err)
	}
}
