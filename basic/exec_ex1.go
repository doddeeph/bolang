package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var out1, _ = exec.Command("ls", "-l").Output()
	fmt.Printf(" --> ls -l\n%s\n", string(out1))

	var out2, _ = exec.Command("pwd").Output()
	fmt.Printf(" --> pwd\n%s\n", string(out2))

	//var out3, _ = exec.Command("git", "config", "user.name").Output()
	//fmt.Printf(" -> git config user.name\n%s\n", string(out3))

	var out3 []byte
	if runtime.GOOS == "windows" {
		out3, _ = exec.Command("cmd", "-C", "git", "config", "user.name").Output()
	} else {
		out3, _ = exec.Command("bash", "-c", "git config user.name").Output()
	}
	fmt.Printf(" --> git config user.name\n%s\n", string(out3))
}
