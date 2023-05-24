package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Data struct {
	d, sl, f  bool
	ext, path string
}

func main() {
	data, err := Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := FlagCompatibility(&data); err != nil {
		fmt.Println(err)
	} else {
		MyFind(data)
	}
}

func Parse() (Data, error) {
	d := flag.Bool("d", false, "d flag")
	f := flag.Bool("f", false, "f flag")
	sl := flag.Bool("sl", false, "sl flag")
	ext := flag.String("ext", "", "ext flag")

	flag.Parse()
	data := Data{d: *d, f: *f, sl: *sl, ext: *ext, path: os.Args[len(os.Args)-1]}
	if len(os.Args) == 1 {
		return data, errors.New("Missing file path")
	}
	return data, nil
}

func FlagCompatibility(data *Data) error {
	if data.f == false && data.ext != "" {
		return errors.New("The \"-ext\" flag can be used when using the \"-f\" flag")
	}
	if !data.f && !data.d && !data.sl {
		data.f, data.d = true, true
	}
	return nil
}

func MyFind(data Data) ([]string, error) {
	findPaths := make([]string, 0)
	err := filepath.Walk(data.path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if data.f && !info.IsDir() && info.Mode().Type() != os.ModeSymlink {
				if data.ext != "" {
					if strings.HasSuffix(path, "."+data.ext) {
						fmt.Println(path)
					}
				} else {
					fmt.Println(path)
				}
			}
			if data.d && info.IsDir() && info.Mode().Type() != os.ModeSymlink {
				fmt.Println(path)
			}
			if data.sl && info.Mode().Type() == os.ModeSymlink {
				if str, err := filepath.EvalSymlinks(path); err != nil {
					fmt.Println("[broken]")
				} else {
					fmt.Println(path, "->", str)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return findPaths, nil
}
