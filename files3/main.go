package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

//find /mnt/csdrive/cassandra/data/ms3/contact_personalisation3-a632ade0bd3e11e6bf5d6331f56c768a/snapshots/1581130801506 -type f -links 1 -printf "%s\n"

func main() {

	// fmt.Println(strings.Trim(getShellOutput(), "\n"))
	st := strings.Split(getShellOutput(), "\n")
	// fmt.Println(len(st))
	// fmt.Println(st[4])
	// sl := strings.Split(getShellOutput(), " ")
	for i := range st[:len(st)-1] {
		fmt.Println(st[i])
		// fmt.Println(i)
		getFileInfo(st[i])
		// fmt.Printf("%T\n", st[i])
	}
	// files := "/home/solid/GO/work/src/github.com/VitaliyKhatrus/files3/new/3/snapshot/158001/terra.txt"
	// getFileInfo(files)

}

func getFileInfo(filename string) {
	// fmt.Println(filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Fatalf("stat: %v", err)
	}
	log.Printf("file %q: size: %d bytes", stat.Name(), stat.Size())
	// log.Printf("file %q: size: %d, mod. time: %q", stat.Name(), stat.Size(), stat.ModTime())
	// log.Printf("file sys: %T", stat.Sys())

	// sysstat := stat.Sys().(*syscall.Stat_t)
	// log.Printf("file %q: access time: %d, mod. time: %d, change time: %d", stat.Name(), sysstat.Atim.Nano(), sysstat.Mtim.Nano(), sysstat.Ctim.Nano())
}

func getShellOutput() string {
	cmd, err := exec.Command("/bin/bash", "-c", "find /home/solid/GO/work/src/github.com/VitaliyKhatrus/files3/new -type d -name 'snapshot' -exec find {} -type f \\;").Output()
	if err != nil {
		// fmt.Fprintln(w, "Error: ", err)
		log.Fatal("Error: ", err)
	}

	return string(cmd)
}
