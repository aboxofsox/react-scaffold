package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
)

//go:embed templates/html/index.html
var markup string

//go:embed templates/css/styles.css
var css string

//go:embed templates/ts/index.tsx
var indexTsx string

//go:embed templates/esbuild/build.js
var esbuildBuild string

//go:embed templates/esbuild/dev.js
var esbuildDev string

func createHtml() {
	file, err := os.OpenFile("index.html", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := file.WriteString(markup); err != nil {
		log.Fatal(err.Error())
	}
}

func createIndexTsx() {
	file, err := os.OpenFile(filepath.Join("src", "index.tsx"), os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := file.WriteString(indexTsx); err != nil {
		log.Fatal(err.Error())
	}
}

func createCss() {
	file, err := os.OpenFile(filepath.Join("src", "styles", "styles.css"), os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := file.WriteString(css); err != nil {
		log.Fatal(err.Error())
	}
}

func createEsbuild() {
	fileA, err := os.OpenFile(filepath.Join("scripts", "build.js"), os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := fileA.WriteString(esbuildBuild); err != nil {
		log.Fatal(err.Error())
	}

	fileB, err := os.OpenFile(filepath.Join("scripts", "dev.js"), os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := fileB.WriteString(esbuildDev); err != nil {
		log.Fatal(err.Error())
	}
}

func doCreation() {
	createHtml()
	createIndexTsx()
	createCss()
	createEsbuild()
}
