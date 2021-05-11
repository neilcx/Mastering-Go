package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var maxSyscalls = 0

const SYSCALLFILE = "SYSCALLS"

func main() {
	var SYSTEMCALLS []string
	f, err := os.Open(SYSCALLFILE)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1) // remove spaces
		line = strings.Replace(line, "SYS_", "", -1) // remove prefix
		temp := strings.ToLower(strings.Split(line, "=")[0]) //split by = into x parts , get the first  part, to lower

		SYSTEMCALLS = append(SYSTEMCALLS, temp)
		maxSyscalls++
	}

	COUNTER := make([]int, maxSyscalls)
	var regs syscall.PtraceRegs
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err = cmd.Start()
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Wait:", err)
	}

	pid := cmd.Process.Pid
	fmt.Println("Process ID:", pid)

	before := true
	forCount := 0
	for {
		if before {
			err := syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}
			if regs.Orig_rax > uint64(maxSyscalls) {
				fmt.Println("Unknown:", regs.Orig_rax)
				return
			}

			COUNTER[regs.Orig_rax]++
			forCount++
		}

		//continue the execution of the program that is being traced
		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			fmt.Println("PtraceSyscall:", err)
			return
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			fmt.Println("Wait4:", err)
			return
		}
		before = !before // ? only once , so for for each syscall twice ?

	}

	for i, x := range COUNTER {
		if x != 0 {
			fmt.Println(SYSTEMCALLS[i], "->", x)
		}
	}
	fmt.Println("Total System Calls:", forCount)
}
