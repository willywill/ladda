package shipper

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/willywill/ladda/pkg/formatter"
)

// Writes a file to a target directory
func WriteFile(file []byte, extension string) (ok bool, err error) {
	path, pathWithFile := formatter.GetFilePath(extension)

	err = os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return false, err
	}

	err = ioutil.WriteFile(pathWithFile, file, 0755)

	debugString, _ := os.LookupEnv("DEBUG")
	debug, _ := strconv.ParseBool(debugString)

	if err != nil {
		if debug {
			log.Println("Error occurred when writing file", err)
		}
		return false, err
	}

	if debug {
		log.Printf("File written to: %s", path)
	}

	return true, nil
}
