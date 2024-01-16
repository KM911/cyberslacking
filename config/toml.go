package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KM911/hotpot/util"

	"github.com/pelletier/go-toml/v2"
)

const (
	TomlFile = "config.toml"
)

// TODO : Create a struct to store the config
type Toml struct {
}

var (
	DefaultToml = Toml{}
	UserToml    = Toml{}
)

func CreateDefaultTomlConfiguration() {
	f, err := os.Create(TomlFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = toml.NewEncoder(f).Encode(DefaultToml)
	if err != nil {
		panic(err)
	}
}

func LoadTomlConfiguration() {
	if _, err := os.Stat(TomlFile); os.IsNotExist(err) {
		CreateDefaultTomlConfiguration()
		fmt.Println("Configurations file created, please edit it and restart.")
		os.Exit(0)
	}
	file, err := os.ReadFile(filepath.Join(util.CmdPath(), TomlFile))
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(file, &UserToml)
	if err != nil {
		panic(err)
	}
}
