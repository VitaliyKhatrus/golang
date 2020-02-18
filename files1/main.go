package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"syscall"
)

func main() {
	var files []string
	var x int64

	root := "/mnt/csdrive/cassandra/data"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	for i, file := range files {
		re := regexp.MustCompile("snapshot")
		regExp := re.FindString(file)
		if regExp != "" {
			// fmt.Println(file)
			if checkHardLink(file) {
				// getFileInfo(file)
				value := getFileInfo(file, i)
				x += value
			}

		}

	}
	// fmt.Println(x)
	switch check := x; {
	case check < 1000000:
		fmt.Println("Total size: ", x, "bytes")
	case check > 1000000000:
		fmt.Println("Total size: ", x/1024000000, "Gb")

	}
}
func getFileInfo(filename string, count int) int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Fatalf("stat: %v", err)
	}
	// fmt.Printf("%s ==> %d bytes \n", filename, stat.Size())
	fmt.Printf("[%d: %s ==> %d bytes \n", count, filename, stat.Size())

	// log.Printf("file %q: size: %d, mod. time: %q", stat.Name(), stat.Size(), stat.ModTime())
	// log.Printf("file sys: %T", stat.Sys())
	return stat.Size()

}

func checkHardLink(filename string) bool {
	fi, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}

	// https://github.com/docker/docker/blob/master/pkg/archive/archive_unix.go
	// in 'func setHeaderForSpecialDevice()'
	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		err = errors.New("cannot convert stat value to syscall.Stat_t")
		log.Fatal(err)
	}

	// The index number of this file's inode:
	// inode := uint64(s.Ino)
	// Total number of files/hardlinks connected to this file's inode:
	nlink := uint32(s.Nlink)

	if nlink == 1 {
		return true
	} else {
		return false
	}
}
