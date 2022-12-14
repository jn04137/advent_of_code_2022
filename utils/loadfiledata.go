package loadfiledata

import (
	"log"
	"os"
)

func LoadData(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
