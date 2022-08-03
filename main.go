package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	dirs = []string{
		"src",
		"dist",
		"scripts",
		filepath.Join("src", "components"),
		filepath.Join("src", "styles"),
	}
	files = []string{
		filepath.Join("src", "index.tsx"),
		filepath.Join("src", "styles", "styles.css"),
		filepath.Join(".", "index.html"),
	}
)

func fpj(paths ...string) string { return filepath.Join(paths...) }

func scaffold() {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal(err.Error())
		}
	}

	for _, file := range files {
		if _, err := os.Create(file); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func npmInit() {
	npm := exec.Command("npm", "init", "-y")
	out, err := npm.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(out))
}

func npmInstall(libs ...string) {
	if len(libs) == 0 {
		log.Fatal("no libraries to install")
	}

	flags := "i -s"
	flagSplit := strings.Split(flags, " ")
	flagSplit = append(flagSplit, libs...)

	for _, fs := range flagSplit {
		fmt.Println(fs)
	}

	npm := exec.Command("npm", flagSplit...)
	out, err := npm.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(out))
}

func npmInstallDev(libs ...string) {
	if len(libs) == 0 {
		log.Fatal("no libraries to install")
	}

	flags := "i -D"
	flagSplit := strings.Split(flags, " ")
	flagSplit = append(flagSplit, libs...)

	for _, fs := range flagSplit {
		fmt.Println(fs)
	}

	npm := exec.Command("npm", flagSplit...)
	out, err := npm.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(out))
}

func tailwindInit() {
	npx := exec.Command("npx", "tailwindcss", "init")
	out, err := npx.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(out))
}

func tsInit() {
	tsc := exec.Command("tsc", "--init")
	out, err := tsc.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(out))
}

func main() {
	scaffold()
	npmInit()
	npmInstall("react", "react-dom", "esbuild")
	npmInstallDev("tailwindcss@latest", "@iconify/react", "typescript")
	tailwindInit()
	tsInit()

}
