package verbose

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"syscall"
	"testing"
)

func getExitCode(err error) (int, error) {
	exitCode := 0
	if exiterr, ok := err.(*exec.ExitError); ok {
		if procExit := exiterr.Sys().(syscall.WaitStatus); ok {
			return procExit.ExitStatus(), nil
		}
	}
	return exitCode, fmt.Errorf("failed to get exit code")
}

func runCommandWithOutput(cmd *exec.Cmd) (output string, exitCode int, err error) {
	exitCode = 0
	out, err := cmd.CombinedOutput()
	if err != nil {
		var exiterr error
		if exitCode, exiterr = getExitCode(err); exiterr != nil {
			// TODO: Fix this so we check the error's text.
			// we've failed to retrieve exit code, so we set it to 127
			exitCode = 127
		}
	}
	output = string(out)
	return
}


func runCommandWithStdoutStderr(cmd *exec.Cmd) (exitCode int, err error) {
	exitCode = 0
	var stderrBuffer,stdoutBuffer bytes.Buffer
	stderrPipe, err := cmd.StderrPipe()
	stdoutpipe, err := cmd.StdoutPipe()

	if err != nil {
		return "", "", -1, err
	}

	err = cmd.Start()
	if err != nil {
		var exiterr error
		if exitCode, exiterr = getExitCode(err); exiterr != nil {
			// TODO: Fix this so we check the error's text.
			// we've failed to retrieve exit code, so we set it to 127
			exitCode = 127
		}
	}


	go io.Copy(os.Stdout, stdoutpipe)
	go io.Copy(os.Stderr, stderrPipe)

	err = cmd2.Wait()
	if err != nil {
		var exiterr error
		if exitCode, exiterr = getExitCode(err); exiterr != nil {
			// TODO: Fix this so we check the error's text.
			// we've failed to retrieve exit code, so we set it to 127
			exitCode = 127
		}
	}
	return
}


func runCommand(cmd *exec.Cmd) (exitCode int, err error) {
	exitCode = 0
	err = cmd.Run()
	if err != nil {
		var exiterr error
		if exitCode, exiterr = getExitCode(err); exiterr != nil {
			// TODO: Fix this so we check the error's text.
			// we've failed to retrieve exit code, so we set it to 127
			exitCode = 127
		}
	}
	return
}

func startCommand(cmd *exec.Cmd) (exitCode int, err error) {
	exitCode = 0
	err = cmd.Start()
	if err != nil {
		var exiterr error
		if exitCode, exiterr = getExitCode(err); exiterr != nil {
			// TODO: Fix this so we check the error's text.
			// we've failed to retrieve exit code, so we set it to 127
			exitCode = 127
		}
	}
	return
}

func logDone(message string) {
	fmt.Printf("[PASSED]: %s\n", message)
}

func stripTrailingCharacters(target string) string {
	target = strings.Trim(target, "\n")
	target = strings.Trim(target, " ")
	return target
}

func errorOut(err error, t *testing.T, message string) {
	if err != nil {
		t.Fatal(message)
	}
}

func errorOutOnNonNilError(err error, t *testing.T, message string) {
	if err == nil {
		t.Fatalf(message)
	}
}

func nLines(s string) int {
	return strings.Count(s, "\n")
}

//func deis(bash string , arg string ,  cmd string )
