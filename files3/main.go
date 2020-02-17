package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

//find /mnt/csdrive/cassandra/data/ms3/contact_personalisation3-a632ade0bd3e11e6bf5d6331f56c768a/snapshots/1581130801506 -type f -links 1 -printf "%s\n"

func main() {
	getFileInfo("hello.txt")
	fmt.Println(getShellOutput())
}

func getFileInfo(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Fatalf("stat: %v", err)
	}
	log.Printf("file %q: size: %d", stat.Name(), stat.Size())
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
