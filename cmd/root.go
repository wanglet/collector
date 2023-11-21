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
	TYPE     = []string{"image", "binary", "chart", "repo", "package", "pypi"}
	ARCH     = []string{"amd64", "arm64"}
	OS       = []string{"ubuntu:20.04", "ubuntu:22.04"}
	flagType = []string{}
	flagArch = []string{}
	flagOS   = []string{}
	cfgFile  string
	output   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "collector",
	Short: "A tool to collect offline packages.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Run args:", args)
		ctx := cmd.Context()

		apps, err := application.NewApplications(ctx, cfgFile)
		if err != nil {
			return err
		}

		for _, app := range apps {
			fmt.Println(app)
		}

		return nil
	},

	PreRunE: func(cmd *cobra.Command, args []string) error {
		for _, ft := range flagType {
			if !slices.Contains(TYPE, ft) {
				return fmt.Errorf("invalid type: '%s' Valid options are %v", ft, TYPE)
			}
		}
		for _, fa := range flagArch {
			if !slices.Contains(ARCH, fa) {
				return fmt.Errorf("invalid arch: '%s' Valid options are %v", fa, ARCH)
			}
		}
		for _, fo := range flagOS {
			if !slices.Contains(OS, fo) {
				return fmt.Errorf("invalid arch: '%s' Valid options are %v", fo, OS)
			}
		}
		return nil
	},
}

func ExecuteContext(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags()
	rootCmd.Flags().StringSliceVar(&flagArch, "arch", ARCH, fmt.Sprintf("Select architectures %v", ARCH))
	rootCmd.Flags().StringSliceVar(&flagType, "type", TYPE, fmt.Sprintf("Select types %v", TYPE))
	rootCmd.Flags().StringSliceVar(&flagOS, "os", OS, fmt.Sprintf("Select os %v", OS))
	rootCmd.Flags().StringVar(&output, "output-path", "", "output path")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.collector.yaml)")
}
