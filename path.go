package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"
)

//go:embed skel
var skel embed.FS

func rootTo(file string) string {
	return path.Join(os.Getenv("HOME"), "patent-server", file)
}

func skelTo(file string) string {
	return path.Join("skel", file)
}

func Initialize() (err error) {

	if isFirstTime() {
		fmt.Println("It is first time to run server.")
		fmt.Println("It will take few minutes, so BE PATIENT.")
		if err = InstallEssentials(); err != nil {
			return
		}
	}

	if !IsLatest() {
		err = Update()
	}

	return

}

func InstallEssentials() (err error) {

	if err = InitFiles(); err != nil {
		return
	}

	fmt.Println("You should put password for sudo command to install/upgrade essential environment.")
	err = exec.Command(rootTo("initialize.sh")).Run()
	fmt.Println("Installing/Upgrading process has been done! Good luck!")

	return

}

func isFirstTime() bool {
	_, err := os.Stat(rootTo(""))
	return !os.IsNotExist(err)
}

func InitFiles() (err error) {

	_ = os.Mkdir(rootTo(""), 0700)

	files, err := skel.ReadDir(skelTo(""))
	if err != nil {
		return
	}

	for _, file := range files {
		fileName := file.Name()
		if err = copyFile(fileName); err != nil {
			return
		}
	}

	return

}

func copyFile(fileName string) (err error) {

	var skelFile fs.File
	var created *os.File

	if _, err = os.Stat(rootTo(fileName)); err == nil {
		err = nil
		return
	}

	if skelFile, err = skel.Open(skelTo(fileName)); err != nil {
		return
	}
	defer skelFile.Close()

	if created, err = os.Create(rootTo(fileName)); err != nil {
		return
	}
	defer created.Close()

	if err = os.Chmod(rootTo(fileName), 0755); err != nil {
		return
	}

	if _, err = io.Copy(created, skelFile); err != nil && err != io.EOF {
		return
	}

	return

}

func Update() (err error) {

	if err = UpdateVersion(); err != nil {
		return
	}

	err = updateFiles()
	return

}

func updateFiles() (err error) {

	var list []string

	if list, err = getUpdateList(); err != nil {
		return
	}

	for _, file := range list {
		if err = exec.Command("rm", "-rf", file).Run(); err != nil {
			return
		}
	}

	err = InstallEssentials()
	return

}

func getUpdateList() (result []string, err error) {

	var data []byte

	findCmd := exec.Command("find", "/Users/simp7/patent-server", "-type", "directory", "-name", "venv", "-prune", "-and", "!", "-name", "venv", "-o", "-type", "file", "-and", "!", "-name", "*.log", "-and", "!", "-name", "config.*")
	if data, err = findCmd.Output(); err == nil {
		result = strings.Split(string(data), "\n")
	}

	return

}
