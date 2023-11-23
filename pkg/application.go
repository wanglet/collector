package application

import (
	"context"
	"fmt"
	"os"
	"slices"

	"github.com/wanglet/collector/pkg/collector"
	"gopkg.in/yaml.v3"
)

func NewApplications(cfgFile string) ([]*Application, error) {
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
	Version     string                  `yaml:"version"`
	ImageGroups []*collector.ImageGroup `yaml:"imageGroups"`
	Binaries    []*collector.Binary     `yaml:"binaries"`
}

func (app *Application) String() string {
	return fmt.Sprintf("Application{Name: %s, Version: %s, ImageGroups: %v, Binaries: %v}", app.Name, app.Version, app.ImageGroups, app.Binaries)
}

func (app *Application) Validate() error {
	if app.Name == "" {
		return fmt.Errorf("application name is required")
	}
	if app.Version == "" {
		return fmt.Errorf("application version is required")
	}
	if len(app.Binaries) != 0 {
		for _, binary := range app.Binaries {
			if err := binary.Validate(); err != nil {
				return err
			}
		}

	}
	return nil
}

func (app *Application) Collect(ctx context.Context, flagArch []string, flagOS []string, flagType []string, outputPath string) error {
	fmt.Println("Collecting", app)

	if slices.Contains(flagType, "binary") {
		fmt.Println("Collecting binaries", app.Binaries)
		for _, binary := range app.Binaries {
			if !slices.Contains(flagArch, binary.Arch) {
				continue
			}
			fmt.Println("Collecting binary", binary)
			err := binary.Collect(ctx, outputPath)
			if err != nil {
				return err
			}
		}
	}
	return nil

}
