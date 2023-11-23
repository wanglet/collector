/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"
	application "github.com/wanglet/collector/pkg"
)

var (
	Types     = Options{"image", "binary", "chart", "repo", "package", "pypi"}
	Archs     = Options{"amd64", "arm64"}
	OSs       = Options{"ubuntu:20.04", "ubuntu:22.04"}
	flagTypes = []string{}
	flagArchs = []string{}
	flagOSs   = []string{}
	cfgFile   string
	output    string
)

var rootCmd = &cobra.Command{
	Use:   "collector",
	Short: "A tool to collect offline packages.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		apps, err := application.NewApplications(cfgFile)
		if err != nil {
			return err
		}

		// 没有传递参数时，收集所有应用
		if len(args) == 0 {
			for _, app := range apps {
				args = append(args, app.Name)
			}
		}

		for _, app := range apps {
			if !slices.Contains(args, app.Name) {
				continue
			}

			if err = app.Validate(); err != nil {
				return err
			}

			err = app.Collect(ctx, flagArchs, flagOSs, flagTypes, output)
			if err != nil {
				return err
			}
		}

		return nil
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {
		for _, ft := range flagTypes {
			if !Types.has(ft) {
				return fmt.Errorf("invalid type: '%s' Valid options are %v", ft, Types)
			}
		}
		for _, fa := range flagArchs {
			if !Archs.has(fa) {
				return fmt.Errorf("invalid arch: '%s' Valid options are %v", fa, Archs)
			}
		}
		for _, fo := range flagOSs {
			if !OSs.has(fo) {
				return fmt.Errorf("invalid os: '%s' Valid options are %v", fo, OSs)
			}
		}
		return nil
	},
}

func ExecuteContext(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags()
	rootCmd.Flags().StringSliceVar(&flagArchs, "arch", Archs, fmt.Sprintf("Select architectures %v", Archs))
	rootCmd.Flags().StringSliceVar(&flagTypes, "type", Types, fmt.Sprintf("Select types %v", Types))
	rootCmd.Flags().StringSliceVar(&flagOSs, "os", OSs, fmt.Sprintf("Select os %v", OSs))

	rootCmd.Flags().StringVar(&output, "output-path", "", "output path")
	rootCmd.MarkFlagRequired("output-path")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.collector.yaml)")
	rootCmd.MarkFlagRequired("config")
}
