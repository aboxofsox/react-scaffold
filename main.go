package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
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
	_, err := npm.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func npmInstall(libs ...string) {
	if len(libs) == 0 {
		log.Fatal("no libraries to install")
	}

	flags := "i -s"
	flagSplit := strings.Split(flags, " ")
	flagSplit = append(flagSplit, libs...)

	npm := exec.Command("npm", flagSplit...)
	_, err := npm.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func npmInstallDev(libs ...string) {
	if len(libs) == 0 {
		log.Fatal("no libraries to install")
	}

	flags := "i -D"
	flagSplit := strings.Split(flags, " ")
	flagSplit = append(flagSplit, libs...)

	npm := exec.Command("npm", flagSplit...)
	_, err := npm.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func tailwindInit() {
	npx := exec.Command("npx", "tailwindcss", "init")
	_, err := npx.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func tsInit() {
	tsc := exec.Command("tsc", "--init")
	_, err := tsc.Output()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func do() {
	start := time.Now()

	log.Println("creating directory structure")
	scaffold()

	log.Println("initializing npm")
	npmInit()

	log.Println("installing npm packages")
	npmInstall("react", "react-dom", "esbuild")

	log.Println("installing npm dev dependencies")
	npmInstallDev("tailwindcss@latest", "@iconify/react", "typescript")

	log.Println("initializing tailwindcss")
	tailwindInit()

	log.Println("initializing typescript")
	tsInit()

	log.Println("creating necessary files")
	doCreation()
	since := time.Since(start)

	log.Printf("done. %.2f seconds\n", since.Seconds())
}

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "scaffold":
		do()
	case "serve":
		serve()
	case "both":
		do()
		serve()
	default:
		log.Println("unsupported argument")
	}
}
