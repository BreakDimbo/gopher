package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose process message")
var n sync.WaitGroup
var done = make(chan struct{})

func main() {
	// Determine the initial directories
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Wait for cancel
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// Traverse the dir
	filesize := make(chan int64)
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, filesize)
	}

	go func() {
		n.Wait()
		close(filesize)
	}()

	// Print the disk usage
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range filesize {
				// Do nothing
			}
			return
		case size, ok := <-filesize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}

// if a channel is closed, will receive zero value all the time
func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, filesize chan<- int64) {
	defer n.Done()
	if canceled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, filesize)
		} else {
			filesize <- entry.Size()
		}
	}
}

var semo = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case semo <- struct{}{}:
	case <-done:
		return nil
	}
	// acquire a token
	defer func() { <-semo }() // release a token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}
