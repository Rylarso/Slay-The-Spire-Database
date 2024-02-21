package fileHandling

import (
	"io/fs"
	"log"
	"os"
)

// Reads the given directory and returns an array of an interface than can call .Name(), .IsDir(), and .Type()
func GetFiles(dirPath string) []fs.DirEntry {
	runs, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	return runs
}

// Extracts the data from the filePath and turns it into a string
func ReadRun(filePath string) string {

	runData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(runData)
}
