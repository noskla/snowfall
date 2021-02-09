//package main

// package not being compiled as it's being replaced by just executing Sass
// executable so it can be compiled for every platform

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/wellington/go-libsass"
)

// CompileSASS loops through the files in the sass directory, then runs the compiler from go-libsass on them
// and outputs the compressed CSS files inside the static directory. Note that this could be
// pain when dealing with compiling on other platforms than native. Workaround is using Docker, WSL, Qemu or
// other virtualization that allows you to compile C binaries for that system.
func CompileSASS() {
	sassPath := getDirectoryPath("sass/")
	outputPath := getDirectoryPath("static/css/")

	files, err := ioutil.ReadDir(sassPath)
	errorOccurred(err, true)

	p := libsass.IncludePaths([]string{sassPath})
	s := libsass.OutputStyle(libsass.COMPRESSED_STYLE)

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(sassPath, file.Name())
			sass, err := os.Open(filePath)
			if errorOccurred(err, false) {
				continue
			}

			cssFilename := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + ".css"
			css, err := os.Create(outputPath + cssFilename)
			errorOccurred(err, true)

			comp, err := libsass.New(sass, css, p, s)
			errorOccurred(err, true)

			errorOccurred(comp.Run(), true)
			sass.Close()
			css.Close()
		}
	}

}
