package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type md5sumInfo struct {
	path  string
	match bool
	err   error
}

func readMd5sum(fpath string) (map[string]string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read md5sum file")
	}

	md5Map := make(map[string]string)
	sc := bufio.NewScanner(file)
	for ln := 0; sc.Scan(); ln++ {
		feilds := strings.Fields(sc.Text())
		if len(feilds) != 2 {
			log.Fatalf("Failed to get filed", feilds)
		}
		log.Printf("File: %s, Md5sum: %s", feilds[0], feilds[1])
		md5Map[feilds[0]] = feilds[1]
	}
	return md5Map, nil
}

func md5sumWorker(path, sig string, ch chan *md5sumInfo) {
	m := &md5sumInfo{path: path}
	s, err := fileMD5(path)
	m.err = err
	m.match = s == sig
	ch <- m
	log.Print("Added to channel: ", m)
}

func fileMD5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Print("Failed to open file.")
		return "", errors.Wrap(err, "Failed to open file")
	}
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", errors.Wrap(err, "Failed to calculate md5sum.")
	}

	md5 := fmt.Sprintf("%x", hash.Sum(nil))
	log.Printf("Calculated Md5sum for file: %s, md5sum: %s", path, md5)
	return md5, nil
}

func main() {
	// Read Md5sum from md5sum.txt
	log.Print("Reading md5sum from file.")
	md5sumMap, err := readMd5sum("md5sum.txt")
	if err != nil {
		log.Fatal("Error reading md5sum file: ", err)
	}
	log.Print("Completed reading md5sum from file.")

	log.Print("Calculating md5sum...")
	ch := make(chan *md5sumInfo)
	for path, sig := range md5sumMap {
		go md5sumWorker(path, sig, ch)
	}

	ok := true
	for range md5sumMap {
		m := <-ch
		switch {
		case m.err != nil:
			log.Printf("FAILED: %s error - %s\n", m.path, m.err)
			ok = false
		case !m.match:
			log.Printf("FAILED: %s signature mismatch\n", m.path)
			ok = false
		case m.match:
			log.Print("Md5sum matched for ", m.path)
		}

	}

	if !ok {
		log.Fatal("Failed to complete.")
	}
}
