package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files := os.Args[1:]

	//  поиск архив папки
	archiveFileName := ""
	info, err := os.Stat(files[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	if info.IsDir() {
		archiveFileName = files[0]
		files = files[1:]
	}

	for _, file := range files {
		if err := createArchive(file, archiveFileName); err != nil {
			fmt.Println(err)
			return
		}

	}
}
func createArchive(fileName, archiveFileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	// избавляемся от раширения файла и создаем архивированный файл
	ext := filepath.Ext(fileName)
	fileName = strings.TrimSuffix(fileName, ext)
	archiveFile, err := os.Create(fmt.Sprintf("%s/%s_%d.tar.gz",
		archiveFileName, fileName, info.ModTime().Unix()))
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	if err = writeToArchive(fileName, archiveFile); err != nil {
		return err
	}

	return nil
}

func writeToArchive(fileName string, archiveFile *os.File) error {
	gw := gzip.NewWriter(archiveFile)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	err := addToArchive(tw, fileName)
	return err
}

func addToArchive(tw *tar.Writer, filename string) error {
	return nil
}

// func ArchiveFile(fileName string) {
// 	if err := targz.Compress(fileName, "./"+archiveFileName); err != nil {
// 		fmt.Println(err)
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	archiveFileName, err := FindAtchiveFile(args[0])
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {

// 	}
// 	fmt.Println(archiveFileName)
// }

// func FindAtchiveFile(path string) (string, error) {
// 	info, err := os.Stat(path)
// 	if err != nil {
// 		return "", err
// 	}
// 	if info.IsDir() {
// 		return path, nil
// 	}
// 	return "", nil
// }

// # Will create file /path/to/logs/some_application_1600785299.tag.gz
// # where 1600785299 is a UNIX timestamp made from `some_application.log`'s [MTIME](https://linuxize.com/post/linux-touch-command/)
// ~$ ./myRotate /path/to/logs/some_application.log

// # Will create two tar.gz files with timestamps (one for every log)
// # and put them into /data/archive directory
// ~$ ./myRotate -a /data/archive /path/to/logs/some_application.log /path/to/logs/other_application.log
