package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	st := strings.Split(getShellOutput(), "\n")
	var x int64
	for i := range st[:len(st)-1] {
		fmt.Println(st[i])
		value := getFileInfo(st[i])
		x += value

	}
	fmt.Println(x)
	switch check := x; {
	case check < 1000:
		fmt.Println("Total size: ", x, "bytes")
	case check > 1000000000:
		fmt.Println("Total size: ", x/1024000000, "Gb")

	}
}

func getFileInfo(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Fatalf("stat: %v", err)
	}
	log.Printf("size: %d bytes", stat.Size())
	// log.Printf("file %q: size: %d, mod. time: %q", stat.Name(), stat.Size(), stat.ModTime())
	// log.Printf("file sys: %T", stat.Sys())
	return stat.Size()

}

func getShellOutput() string {
	// cmd, err := exec.Command("/bin/bash", "-c", "find /home/user/GO/src/github.com/VitaliyKhatrus/golang/files3 -type d -name 'snapshot' -exec find {} -type f \\;").Output()
	cmd, err := exec.Command("/bin/bash", "-c", "find /mnt/csdrive/cassandra/data -type d -name 'snapshots' -exec find {} -type f -links 1 \\;").Output()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return string(cmd)
}
