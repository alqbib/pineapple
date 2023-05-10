package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"opensearch-tool/tool/config"
	"opensearch-tool/tool/router"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "opensearch tool",
	Short: "opensearch tool",
	Run: func(cmd *cobra.Command, args []string) {
		if cfgFile == "" {
			log.Fatal("缺少配置文件!")
		}
		cfg, err := config.InitConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
		router.Start(cfg)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "", "config file")
}
