package application

import (
	"context"
	"os"

	"github.com/wanglet/collector/pkg/collector"
	"gopkg.in/yaml.v3"
)

func NewApplications(ctx context.Context, cfgFile string) ([]*Application, error) {
	var applications []*Application
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &applications); err != nil {
		return nil, err
	}

	return applications, nil
}

type Application struct {
	Name        string                  `yaml:"name"`
	ImageGroups []*collector.ImageGroup `yaml:"imageGroups"`
}

func (app *Application) collect(flagArch []string, flagOS []string, flagType []string) error {
	return nil

}
