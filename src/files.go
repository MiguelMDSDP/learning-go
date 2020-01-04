package main

import (
	"fmt"
	"os"
)

// CreateFiles - creates the files with the name passed on the "files" argument into the "rootDir" directory
func CreateFiles(rootDir string, files ...string) {
	for _, name := range files {
		filePath := fmt.Sprintf("%s/%s.%s", rootDir, name, "txt")

		file, err := os.Create(filePath)
		defer file.Close()

		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("File %s created successfully.\n", file.Name())
	}
}

func main() {
	tmp := os.TempDir()

	CreateFiles(tmp)
	CreateFiles(tmp, "first-test")
	CreateFiles(tmp, "second-test", "third-test", "fourth-test")
}
