package system

import "runtime"

var (
	KillByname func(_name string)
	KillByPid  func(_pid string)
)

func killByname(_name string) {
	ExecuteCommand("killall " + _name)
}

func taskKillByName(_name string) {
	ExecuteCommand("taskkill /F /IM " + _name)
}

func killByPid(_pid string) {
	ExecuteCommand("kill -9 " + _pid)
}

func taskKillByPid(_pid string) {
	ExecuteCommand("taskkill /F /PID " + _pid)
}

func init() {
	if runtime.GOOS == "windows" {
		KillByname = taskKillByName
		KillByPid = taskKillByPid
	} else {
		KillByname = killByname
		KillByPid = killByPid
	}
}
