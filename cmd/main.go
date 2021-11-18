package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	configPath := os.Args[1]
	if configPath == "" {
		panic("empty config path")
	}
	fi, err := os.Stat(configPath)
	if err != nil {
		panic(err)
	}
	var dir string
	if fi.IsDir() {
		dir = configPath
	} else {
		dir = filepath.Dir(configPath)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(dir)
	for _, file := range files {
		fileName := file.Name()
		ext := filepath.Ext(fileName)
		if len(ext) < 2 {
			continue
		}
		if !stringInSlice(ext[1:], viper.SupportedExts) {
			continue
		}

		viper.SetConfigFile(dir + "/" + fileName)
		err := viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
	}

	generate(dir)

}
