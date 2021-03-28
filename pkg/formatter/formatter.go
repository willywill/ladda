package formatter

import (
	"fmt"
	"os"

	"github.com/willywill/ladda/pkg/url"
)

func GetFilePath(extension string) (path string, pathWithFile string) {
	writePath, _ := os.LookupEnv("WRITE_PATH")
	currentDir, _ := os.Getwd()

	urlString := url.GenerateSafeString(12)

	path = fmt.Sprintf("%s/%s", currentDir, writePath)
	pathWithFile = fmt.Sprintf("%s/%s%s", path, urlString, extension)

	return path, pathWithFile
}
