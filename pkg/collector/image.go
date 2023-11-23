package collector

import (
	"fmt"
	"regexp"
)

var (
	imagePattern string         = `^(?:(?P<registry>.*[.:].*?)/)?(?:(?P<namespace>[^/]+)/)?(?P<repository>[^:/]+)(?::(?P<tag>[^/@]+))?(?:@(?P<digest>[^/]+))?$`
	imageRe      *regexp.Regexp = regexp.MustCompile(imagePattern)
)

type ImageGroup struct {
	Output map[string]string `yaml:"output,omitempty"`
	Images []Image           `yaml:"images"`
}

func (ig *ImageGroup) Copy(targetRegistry string, platforms []string) {
	fmt.Println("copy")
}

type Image struct {
	Source string `yaml:"source"`
}

func (i *Image) Parse() (map[string]string, error) {
	match := imageRe.FindStringSubmatch(i.Source)
	if match == nil {
		return nil, fmt.Errorf("unable to parse image: %v", i.Source)
	}

	result := make(map[string]string)
	for i, name := range imageRe.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}
	return result, nil
}
