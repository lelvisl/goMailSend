package main

import (
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

func configToml(fname string) Config {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var config Config
	if err := toml.Unmarshal(buf, &config); err != nil {
		// fmt.Println("Файл конфигурации некорректен")
		panic(err)

	}
	return config
}

type Config struct {
	Mail struct {
		Host    string
		Port    int
		From    string
		To      []string
		Subject string
		Body    string
		Attach  string
	}
}

var cnf = configToml("config.toml")
