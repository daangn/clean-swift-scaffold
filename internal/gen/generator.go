package gen

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/converter"
	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
	"gopkg.in/yaml.v2"
)

type GeneratorConfig struct {
	Name           string
	UsecasesString string
	SourcePath     string
	TestPath       string
	ConfigFilePath string
}

type Generator struct {
	name       string
	usecases   []string
	sourcePath string
	testPath   string
	config     *model.Config
}

func NewGenerator(config GeneratorConfig) *Generator {

	return &Generator{
		name:       config.Name,
		usecases:   strings.Split(config.UsecasesString, ","),
		sourcePath: config.SourcePath,
		testPath:   config.TestPath,
		config:     readConfig(config.ConfigFilePath),
	}
}

func (gen *Generator) Run() {
	today := time.Now()

	header := converter.NewHeaderConverter(
		gen.config,
		today,
	)

	source := converter.NewSourceConverter(
		gen.name,
		gen.usecases,
		gen.sourcePath,
		gen.testPath,
		today,
		gen.config.Intentation,
		header,
	)

	// TODO: RUN, sources, error & save to destinations :]
	source.RenderAll()

	// TODO: save to destination. you are so lucy :]
}

func readConfig(path string) *model.Config {

	content, err := ioutil.ReadFile(path)

	if err != nil {
		return &model.Config{
			Org:          "Unknown",
			Copyright:    "Geektree0101",
			TemplatePath: "./templates",
		}
	}

	config := &model.Config{}

	err = yaml.Unmarshal(content, &config)

	if err != nil {
		return &model.Config{
			Org:          "Unknown",
			Copyright:    "Geektree0101",
			TemplatePath: "./templates",
		}
	}

	return config
}
