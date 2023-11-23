package cmd

import "slices"

type Options []string

func (o *Options) has(val string) bool {
	return slices.Contains(*o, val)
}
