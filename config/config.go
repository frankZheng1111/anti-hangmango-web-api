package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const PROJECT_NAME string = "anti-hangmango-web-api"

type JSONConfig struct {
	ENV          string
	UserCount    int
	HangmanCount int
	APIUrl       string
	Hangman      Hangman
}

type Hangman struct {
	Hp             int
	Dictionary     []string
	DictionaryName string
}

var Config JSONConfig

func init() {
	InitConfig(&Config)
}

func (config *JSONConfig) ConfigFolderPath() string {
	var projectGoPath string
	// 在GOPATH 寻找项目存在的那一条路径
	//
	goPaths := strings.Split(os.Getenv("GOPATH"), ":")
	for _, goPath := range goPaths {
		if _, err := os.Stat(filepath.Join(goPath, "src", PROJECT_NAME)); err == nil {
			projectGoPath = goPath
			break
		}
	}
	return filepath.Join(projectGoPath, "src", PROJECT_NAME, "config")
}

func (config *JSONConfig) ConfigFilePath(env string) string {
	return filepath.Join(config.ConfigFolderPath(), fmt.Sprintf("%s.json", env))
}

func (config *JSONConfig) InitDictionary() error {
	dictionaryPath := filepath.Join(config.ConfigFolderPath(), config.Hangman.DictionaryName)
	content, err := ioutil.ReadFile(dictionaryPath)
	if err != nil {
		return err
	}
	for _, letter := range strings.Split(string(content), "\n") {
		config.Hangman.Dictionary = append(config.Hangman.Dictionary, strings.ToLower(letter))
	}
	return nil
}

func InitConfig(config *JSONConfig) {
	env := os.Getenv("GOENV")
	if env == "" {
		env = "dev"
	}
	config.ENV = env

	filePath := config.ConfigFilePath(env)
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		panic(err)
	}
	if err = config.InitDictionary(); err != nil {
		panic(err)
	}
}
