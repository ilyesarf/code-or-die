package main

import (
	"archive/zip"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

func checkRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[isRoot] Unable to get current user: %s", err)
	}

	return currentUser.Username == "root"
}

func getflags() (int, string) {
	var interval = flag.Int("interval", 30, "set interval time in minutes > 1 minute")

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var dir_path = flag.String("dir", pwd, "set direcotry")

	flag.Parse()

	return *interval * 60, *dir_path
}

func zipper(dir_path, zip_path string) error { //yes, i copied and pasted this.. zipping in Go is too complicated
	f, err := os.Create(zip_path)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	return filepath.Walk(dir_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Deflate

		header.Name, err = filepath.Rel(filepath.Dir(dir_path), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

func md5sum(path string) string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		panic(err)
	}

	return string(hash.Sum(nil))
}

func cod(interval int, prevhash, dir_path, zip_path string) string {
	zipper(dir_path, zip_path)
	hash := md5sum(zip_path)
	if hash == prevhash {
		os.RemoveAll(dir_path)
		log.Fatal("Your project is deleted, you should've been more productive")
	} else {
		return hash
	}

	return ""
}

func main() {
	if checkRoot() {
		interval, dir_path := getflags()
		zip_path := fmt.Sprintf("/tmp/%s.zip", filepath.Base(dir_path))
		log.Print(zip_path)
		hash := cod(interval, "", dir_path, zip_path)

		for {
			time.Sleep(time.Duration(interval) * time.Second)

			hash = cod(interval, hash, dir_path, zip_path)
		}
	} else {
		log.Fatal("SHOULD BE RUNNED AS ROOT (sudo ./cod)")
	}

}
