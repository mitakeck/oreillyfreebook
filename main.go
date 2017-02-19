package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		directory string
		format    string
		category  string
	)

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&directory, "d", "", "specify a directory to save the file")
	f.StringVar(&format, "f", "pdf", "specify the ebook format")
	f.StringVar(&category, "c", "all", "specify the ebook category")

	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	if directory == "" {
		fmt.Errorf("-d (save directory is required)")
		os.Exit(1)
	}

	if err := os.MkdirAll(directory, 0777); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d := Downloader{}
	d.Download(category, format, directory)
}
