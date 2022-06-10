package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run [container name] cmd args
// go run My_cli.go run cmd args
// sudo go run My_cli.go run /bin/bash host name
func main() {
	switch os.Args[1]{
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what??")
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"},os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[3], os.Getpid())

	cmd := exec.Command(os.Args[2]) /////,os.Args[3:]...
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//set root file system for container
	must(syscall.Chroot("/home/rootfs")) 
	//also set the default directory too
	must(os.Chdir("/"))
	//mount /proc 
	must(syscall.Mount("proc","proc","proc",0,""))
	var hostname string
	hostname = os.Args[3]
	must(syscall.Sethostname([]byte(hostname)))
	must(cmd.Run())
}

//panic if anything went wrong
func must(err error) {
	if err != nil {
		panic(err)
	}
}