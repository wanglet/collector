package collector

import (
	"fmt"
	"regexp"
)

/*
定义一组镜像
imageGroups:
  - name: cluster-proportional-autoscaler
    output:
    type: docker
    images:
  - source: 172.18.60.199:5000/cpa/cluster-proportional-autoscaler-arm64:1.8.5
  - source: 172.18.60.199:5000/cpa/cluster-proportional-autoscaler-amd64:1.8.5
  - name: httpd
    images:
  - source: httpd:2.4
*/
type ImageGroup struct {
	Name   string            `yaml:"name"`
	Output map[string]string `yaml:"output,omitempty"`
	Images []Image           `yaml:"images"`
}

func (ig *ImageGroup) Copy(targetRegistry string, platforms []string) {
	fmt.Println("copy")
}

type Image struct {
	Source string `yaml:"source"`
}

/*
解析镜像名称

	{
		input: "k8s.gcr.io/pause:3.3",
		expected: map[string]string{
			"registry":   "k8s.gcr.io",
			"namespace":  "",
			"repository": "pause",
			"tag":        "3.3",
			"digest":     "",
		},
	},
*/
func (i *Image) Parse() (map[string]string, error) {
	pattern := `^(?:(?P<registry>.*[.:].*?)/)?(?:(?P<namespace>[^/]+)/)?(?P<repository>[^:/]+)(?::(?P<tag>[^/@]+))?(?:@(?P<digest>[^/]+))?$`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(i.Source)
	if match == nil {
		return nil, fmt.Errorf("unable to parse image: %v", i.Source)
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}
	return result, nil
}
