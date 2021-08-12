package main

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
)

//go:embed skel
var skel embed.FS

func rootTo(file string) string {
	return path.Join(os.Getenv("HOME"), "patent-server", file)
}

func skelTo(file string) string {
	return path.Join("skel", file)
}

func initialize() (err error) {

	if isFirstTime() {
		err = initFiles()
		if err != nil {
			return
		}
		err = exec.Command(rootTo("initialize.sh")).Run()
	}

	if err != nil {
		return
	}

	return

}

func isFirstTime() bool {
	_, err := os.Open(rootTo(""))
	return err != nil
}

func initFiles() (err error) {

	err = os.Mkdir(rootTo(""), 0700)
	if err != nil {
		return
	}

	files, err := skel.ReadDir(skelTo(""))
	if err != nil {
		return
	}

	for _, file := range files {

		fileName := file.Name()

		err = copyFile(fileName)
		if err != nil {
			return
		}

	}

	return

}

func copyFile(fileName string) (err error) {

	var skelFile fs.File
	var created *os.File

	skelFile, err = skel.Open(skelTo(fileName))
	if err != nil {
		return
	}

	created, err = os.Create(rootTo(fileName))
	if err != nil {
		return
	}

	err = os.Chmod(rootTo(fileName), 0755)
	if err != nil {
		return
	}

	_, err = io.Copy(created, skelFile)
	if err != nil && err != io.EOF {
		return
	}

	err = nil
	return

}
