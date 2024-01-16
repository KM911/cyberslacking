package system

import (
	"os"
	"os/exec"
	"runtime"
)

var (
	CreateCommand func(_command string) *exec.Cmd
)

func init() {
	if runtime.GOOS == "windows" {
		CreateCommand = createCmd
	} else {
		CreateCommand = createBash
	}
}

func redriectStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
}

func createCmd(_command string) (cmd *exec.Cmd) {

	cmd = exec.Command("cmd", "/C", _command)

	return
}

func createBash(_command string) *exec.Cmd {
	return exec.Command("bash", "-c", _command)
}

func ExecuteCommand(_command string) int {
	cmdExecutor := CreateCommand(_command)
	redriectStd(cmdExecutor)
	cmdExecutor.Run()
	return cmdExecutor.ProcessState.ExitCode()
}

func ExecuteCommandSilent(_command string) int {
	cmdExecutor := CreateCommand(_command)
	cmdExecutor.Run()
	return cmdExecutor.ProcessState.ExitCode()
}

func ExecuteCommandResult(_command string) string {
	cmdExecutor := CreateCommand(_command)
	redriectStd(cmdExecutor)
	output, _ := cmdExecutor.Output()
	return string(output)
}

func ExecuteCommandSilentResult(_command string) string {
	cmdExecutor := CreateCommand(_command)
	output, _ := cmdExecutor.Output()
	return string(output)
}
