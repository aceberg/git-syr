package logfile

import (
	"io"
	"log"
	"os"
)

// Output - copy stdout and stderr to file
func Output(path string) {
	if path == "" {
		return
	}

	// open file read/write | create if not exist | clear file at open if exists
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	// save existing stdout | MultiWriter writes to saved stdout and file
	out := os.Stdout
	mw := io.MultiWriter(out, f)

	// get pipe reader and writer | writes to pipe writer come out pipe reader
	r, w, _ := os.Pipe()

	// replace stdout,stderr with pipe writer | all writes to stdout, stderr will go through pipe instead (fmt.print, log)
	os.Stdout = w
	os.Stderr = w

	log.SetOutput(mw)

	_, _ = io.Copy(mw, r)
}

// Code from
// https://gist.github.com/jerblack/4b98ba48ed3fb1d9f7544d2b1a1be287
