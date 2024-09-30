package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"workspace/pkg/yaml"
)

func LoadConfig(v any) {
	for _, data := range configs {
		if err := yaml.Unmarshal(data, v); err != nil {
			log.Fatal(err)
		}
	}
}

func PatchConfig(key string, value any, path ...string) error {
	if ConfigPath == "" {
		return errors.New("config file disabled")
	}

	// empty config is OK
	b, _ := os.ReadFile(ConfigPath)

	b, err := yaml.Patch(b, key, value, path...)
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigPath, b, 0644)
}

type flagConfig []string

func (c *flagConfig) String() string {
	return strings.Join(*c, " ")
}

func (c *flagConfig) Set(value string) error {
	*c = append(*c, value)
	return nil
}

var configs [][]byte

func initConfig() {
	// config as file
	ConfigPath := "Dangsan.yaml"

	if !filepath.IsAbs(ConfigPath) {
		if cwd, err := os.Getwd(); err == nil {
			ConfigPath = filepath.Join(cwd, ConfigPath)
		}
	}
	Info["config_path"] = ConfigPath

}
