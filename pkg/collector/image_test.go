package collector_test

import (
	"reflect"
	"testing"

	"github.com/wanglet/collector/pkg/collector"
)

func TestImageParse(t *testing.T) {
	testCases := []struct {
		input    string
		expected map[string]string
	}{
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
		{
			input: "quay.io/argoproj/argocd:v2.7.6",
			expected: map[string]string{
				"registry":   "quay.io",
				"namespace":  "argoproj",
				"repository": "argocd",
				"tag":        "v2.7.6",
				"digest":     "",
			},
		},
		{
			input: "nginx",
			expected: map[string]string{
				"registry":   "",
				"namespace":  "",
				"repository": "nginx",
				"tag":        "",
				"digest":     "",
			},
		},
		{
			input: "geoffh1977/chrony",
			expected: map[string]string{
				"registry":   "",
				"namespace":  "geoffh1977",
				"repository": "chrony",
				"tag":        "",
				"digest":     "",
			},
		},
		{
			input: "wanglet/kube-eagle:1.1.4",
			expected: map[string]string{
				"registry":   "",
				"namespace":  "wanglet",
				"repository": "kube-eagle",
				"tag":        "1.1.4",
				"digest":     "",
			},
		},
		{
			input: "172.18.60.199:5000/flannelcni/flannel-cni-plugin:v1.0.1-amd64",
			expected: map[string]string{
				"registry":   "172.18.60.199:5000",
				"namespace":  "flannelcni",
				"repository": "flannel-cni-plugin",
				"tag":        "v1.0.1-amd64",
				"digest":     "",
			},
		},
		{
			input: "busybox:1.28@sha256:7d602b12b1d9c1bdbf4c9255c0ba276ac0d9e0cd781a7c13461e4875cfcae509",
			expected: map[string]string{
				"registry":   "",
				"namespace":  "",
				"repository": "busybox",
				"tag":        "1.28",
				"digest":     "sha256:7d602b12b1d9c1bdbf4c9255c0ba276ac0d9e0cd781a7c13461e4875cfcae509",
			},
		},
	}

	for _, tc := range testCases {
		image := collector.Image{Source: tc.input}
		actualValue, err := image.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(actualValue, tc.expected) {
			t.Errorf("\nExpected: %s\n but got: %s", actualValue, tc.expected)
		}

	}
}
