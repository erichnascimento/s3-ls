package main

import (
	"fmt"
	"log"

	"github.com/erichnascimento/s3-ls/pkg/ls"

	"github.com/dustin/go-humanize"
	"github.com/tj/docopt"
)

const version = "0.0.1"

const usage = `
  Usage:
    s3-ls --region <region> <bucket> <file>

  Options:
    -r, --region region   region of bucket
    -h, --help            output help information
    -v, --version         output version
`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.Fatal(err)
	}

	region := args["--region"].(string)
	bucket := args["<bucket>"].(string)
	file := args["<file>"].(string)

	cfg := &ls.Config{
		Bucket: bucket,
		Region: region,
	}

	files, err := ls.List(file, cfg)
	if err != nil {
		log.Fatal(err)
	}

	listFiles(files)
}

func listFiles(files []*ls.FileInfo) {
	fmt.Printf("total %d\n", len(files))

	for _, f := range files {
		fmt.Printf(" %10s %6s %15s %s\n", f.Owner, humanize.Bytes(f.Size), humanize.Time(f.LastModified), f.Name)
	}
}
