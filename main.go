package main

import (
	"fmt"

	"github.com/Rylarso/Slay-The-Spire-Database/file/fileHandling"
)

const runDir = "./internal/file/runs/"

func main() {
	runs := fileHandling.GetFiles(runDir)
	for _, run := range runs {
		filepath := runDir + run.Name()
		fmt.Println(fileHandling.ReadRun(filepath))
	}
}
