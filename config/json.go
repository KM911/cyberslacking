package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/KM911/demo/format"

	"github.com/KM911/demo/util"
)

const (
	JsonConfigurationFile = "config.json"
)

type JsonConfiguration struct {
	Source       string
	Target       string
	BuildCommand string
	BuildOutput  string
	ExecuteBoxs  []ExecuteBox
}

type ExecuteBox struct {
	Args     []string
	Env      []string
	Dir      string
	ExitCode int
}

var (
	DefaultJsonConfiguration = JsonConfiguration{
		Source:       util.WorkingDirectory,
		Target:       TargetPath(),
		BuildCommand: "go build -o sapkin",
		BuildOutput:  "sapkin",
		ExecuteBoxs: []ExecuteBox{
			{
				Args:     []string{},
				Env:      []string{},
				Dir:      "",
				ExitCode: 0,
			},
			{
				Args:     []string{"-f", "config.toml"},
				Env:      []string{},
				Dir:      "",
				ExitCode: 0,
			},
		},
	}
	UserJsonConfiguration = JsonConfiguration{}
)

func TargetPath() string {
	return filepath.Join(filepath.Dir(util.TempDirectory), "paper")
}

func CreateDefaultJson() {
	byte_json, err := json.Marshal(DefaultJsonConfiguration)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(JsonConfigurationFile, byte_json, 0777)
	if err != nil {
		panic(err)
	}
}

func LoadJsonConfiguration() {
	if !util.IsExist(JsonConfigurationFile) {
		CreateDefaultJson()
		format.InfoMessage("Create", "default json configuration file")
		os.Exit(1)
	}

	fileByte, err := os.ReadFile(JsonConfigurationFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileByte, &UserJsonConfiguration)
	if err != nil {
		panic(err)
	}
}
