package helpers

import (
	"os"
	"io/ioutil"
	"log"
	fmt "fmt"
	_ "path"
	host "../host"
)

func GetConfigs(folder string) []string {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	var u []string

	for _, file := range files {
		u = append(u, file.Name())
	}

	return u
}

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		log.Print(err)
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}


func GetHost(config_dir string, config string) *host.Host {
	configs := GetConfigs(config_dir)

	for _, file := range configs {
		if file == config {
			return host.LoadFromFile(fmt.Sprintf("%s/%s", config_dir, file))
		}
	}

	return nil
}

