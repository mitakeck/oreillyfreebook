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
		fmt.Println("-d (save directory is required)")
		usage := `
$ oreillyfreebook -d directory [-f format] [-c category]

-d directory	Specify the directory to save
-f format	Specify the ebook format to download; the default is "pdf"
		"pdf", "mobi", "epub"
-c category	Specify the ebook category to download;
		if not specified, all categories will be download
		"business", "design", "iot", "data", "programming", "security", "web-platform", "webops"`
		fmt.Println(usage)
		os.Exit(1)
	}

	if err := os.MkdirAll(directory, 0777); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d := Downloader{}
	d.Download(category, format, directory)
}
