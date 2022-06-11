package main

import (
	"fmt" //for printing
	"os"
	"os/exec"
	"syscall" //for namespaces, ...
	//for cgroups
	//"path/filepath"
	//"strconv"
	//"io/ioutil"
)

// docker run [container name] cmd args
// go run My_cli.go run cmd args
// sudo go run My_cli.go run /bin/bash hostname RAM
func main() {
	switch os.Args[1]{
	case "run":
		run()
	case "child":
		child()
	default:
		panic("first arg should be run")
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"},os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
		syscall.CLONE_NEWPID |
		syscall.CLONE_NEWNS |
		syscall.CLONE_NEWNET,
		//unshare new mount namespace
		Unshareflags: syscall.CLONE_NEWNS,
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
	var path string
	path = os.Args[4]
	fmt.Printf("container "+path+"\n")
	must(syscall.Chroot("/home/rootfs"+path))
	//also set the default directory inside container
	must(os.Chdir("/"))
	//mount /proc 
	must(syscall.Mount("proc","proc","proc",0,""))
	var hostname string
	hostname = os.Args[3]
	//set hostname of new uts namespace
	must(syscall.Sethostname([]byte(hostname)))

	cgroup()
	must(cmd.Run())
	//unmount /proc after exiting
	must(syscall.Unmount("/proc",0))
	//os.RemoveAll("/")
}

func cgroup() { //just makes directory hi for checking isolation
	//cgroups := "sys/fs/cgroup"
	//pids := filepath.Join(cgroups,"pids")
	//os.Mkdir(filepath.Join(pids,"ourContainer"), 0755)
	os.Mkdir("/hii_for_testing_isolation!", 0755)
}

//panic if anything went wrong
func must(err error) {
	if err != nil {
		panic(err)
	}
}