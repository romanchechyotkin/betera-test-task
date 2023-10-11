package utils

import (
	"os"
)

func CleanTmp() error {
	return os.Remove("tmp/img.jpg")
}
