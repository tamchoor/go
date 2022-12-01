package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

func checkDir(dirFlag *string) {
	fileInfo, err := os.Stat(*dirFlag)
	if err != nil || fileInfo.IsDir() == false && (fileInfo.Mode()&os.ModeSymlink) != os.ModeSymlink || fileInfo.Mode().IsDir() == false {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
}

func createNewFileName(file string, dirFlag string) string {

	var stat syscall.Stat_t
	if err := syscall.Stat(file, &stat); err != nil {
		fmt.Println(err)
		return ""
	}
	filename := strings.TrimSuffix(file, filepath.Ext(file))
	if len(dirFlag) != 0 {
		if dirFlag[len(dirFlag)-1] != '/' {
			dirFlag = dirFlag + "/"
		}
	}
	filename = dirFlag + filename + "_" + strconv.FormatInt(stat.Mtimespec.Sec, 10) + ".tar.gz"
	return filename
}

func makeArh(file string, oldfile string, wg *sync.WaitGroup) {
	arhive, err := os.Create(file)
	if err != nil {
		fmt.Println("4", err)
		return
	}
	defer arhive.Close()
	arh := tar.NewWriter(arhive)
	defer arh.Close()
	defer wg.Done()
	bytes, err := os.ReadFile(oldfile)
	if err != nil {
		fmt.Println("3", err)
		return
	}
	hdr := &tar.Header{
		Name: file,
		Mode: 0600,
		Size: int64(len(bytes)),
	}
	err = arh.WriteHeader(hdr)
	if err != nil {
		fmt.Println("2", err)
		return
	}
	if _, err := arh.Write(bytes); err != nil {
		fmt.Println("1", err)
		return
	}

}

func main() {
	var wg sync.WaitGroup
	dirFlag := flag.String("a", "", "flag to put file into directory")
	flag.Parse()
	if len(*dirFlag) != 0 {
		checkDir(dirFlag)
	}
	// fmt.Println("dirFlag -  ", *dirFlag)

	for i := range flag.Args() {
		wg.Add(1)
		// fmt.Println("f  ", flag.Args()[i])
		filename := createNewFileName(flag.Args()[i], *dirFlag)
		// fmt.Println("f  ", filename)
		makeArh(filename, flag.Args()[i], &wg)
	}
	wg.Wait()
}
