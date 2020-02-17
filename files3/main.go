package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

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
	cmd, err := exec.Command("/bin/bash", "-c", "find /home/solid/GO/work/src/files3/new -type d -name 'snapshot'").Output()
	if err != nil {
		// fmt.Fprintln(w, "Error: ", err)
		log.Fatal("Error: ", err)
	}

	return string(cmd)
}
