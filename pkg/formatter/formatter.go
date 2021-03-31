package formatter

import (
	"fmt"
	"os"

	"github.com/willywill/ladda/pkg/url"
)

func GetFilePath(extension string) (path string, pathWithFile string) {
	// Write path is a static path that we define to write to the container.
	// The container user is responsible for binding the volume to the host machine,
	// to write the file to the desired location
	writePath := "files"
	currentDir, _ := os.Getwd()

	// TODO: Use another method for the file name to prevent collisions?
	urlString := url.GenerateSafeString(12)

	path = fmt.Sprintf("%s/%s", currentDir, writePath)
	pathWithFile = fmt.Sprintf("%s/%s%s", path, urlString, extension)

	return path, pathWithFile
}
