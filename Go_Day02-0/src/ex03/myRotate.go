package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		fmt.Println("Argumens is empty")
		return
	}
	//  поиск архив папки
	archiveFileName := ""
	info, err := os.Stat(fileNames[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	if info.IsDir() {
		archiveFileName = fileNames[0]
		fileNames = fileNames[1:]
	}
	var wg sync.WaitGroup
	for _, fn := range fileNames {
		wg.Add(1)
		go func(name string) {
			err := createArchive(name, archiveFileName)
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(fn)
	}
	wg.Wait()
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

	if err = writeToArchive(file, archiveFile); err != nil {
		return err
	}

	return nil
}

func writeToArchive(file *os.File, archiveFile *os.File) error {
	gw := gzip.NewWriter(archiveFile)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	if err := addToArchive(tw, file); err != nil {
		return err
	}
	return nil
}

func addToArchive(tw *tar.Writer, file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Создаем заголовок tar из данных FileInfo
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	header.Name = file.Name()

	// Запись заголовока в tar-архив
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	// Копирование содержимое файла в tar-архив
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}
