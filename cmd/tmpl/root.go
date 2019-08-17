package main

import (
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tinrab/tmpl"
)

var (
	parametersPath string
	sourcesPath    string
	delimiter      string
	outputPath     string
	test           bool
)

var rootCmd = &cobra.Command{
	Use:   "tmpl",
	Short: "A tool for generating parameterized files from templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().StringVarP(&parametersPath, "parameters", "p", "", "path to parameters")
	_ = rootCmd.MarkFlagRequired("parameters")
	rootCmd.Flags().StringVarP(&sourcesPath, "source", "s", "", "path to sources")
	_ = rootCmd.MarkFlagRequired("source")
	rootCmd.Flags().StringVarP(&delimiter, "delimiter", "d", "\n", "template delimiter")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "tmpl.out", "output file path")
	rootCmd.Flags().BoolVarP(&test, "test", "t", false, "print result to stdout")

	rootCmd.Flags().String("consul-address", "", "consul address")
	_ = viper.BindPFlag("consul-address", rootCmd.Flags().Lookup("consul-address"))
	rootCmd.Flags().String("consul-token", "", "consul token")
	_ = viper.BindPFlag("consul-token", rootCmd.Flags().Lookup("consul-token"))

	rootCmd.Flags().String("vault-address", "", "vault address")
	_ = viper.BindPFlag("vault-address", rootCmd.Flags().Lookup("vault-address"))
	rootCmd.Flags().String("vault-token", "", "vault token")
	_ = viper.BindPFlag("vault-token", rootCmd.Flags().Lookup("vault-token"))

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("tmpl")
	viper.AutomaticEnv()
}

func run() error {
	special := map[string]string{
		"\\a": "\a",
		"\\b": "\b",
		"\\f": "\f",
		"\\n": "\n",
		"\\t": "\t",
		"\\v": "\v",
	}
	for original, escaped := range special {
		delimiter = strings.ReplaceAll(delimiter, original, escaped)
		delimiter = strings.ReplaceAll(delimiter, "\\"+escaped, original)
	}

	parametersMap, err := tmpl.LoadParameters(parametersPath)
	if err != nil {
		return err
	}
	sources, err := tmpl.LoadSources(sourcesPath)
	if err != nil {
		return err
	}
	options := tmpl.Options{
		Parameters:   parametersMap,
		Sources:      sources,
		ConsulConfig: getConsulConfig(),
		VaultConfig:  getVaultConfig(),
	}

	results, err := tmpl.Generate(options)
	if err != nil {
		return err
	}
	var w io.Writer
	if test {
		w = os.Stdout
	} else {
		f, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	}
	return writeResults(results, w)
}

func getConsulConfig() *tmpl.ConsulConfig {
	var consulConfig *tmpl.ConsulConfig
	if viper.GetString("consul-address") != "" {
		consulConfig = &tmpl.ConsulConfig{
			Address: viper.GetString("consul-address"),
			Token:   viper.GetString("consul-token"),
		}
	}
	return consulConfig
}

func getVaultConfig() *tmpl.VaultConfig {
	var vaultConfig *tmpl.VaultConfig
	if viper.GetString("vault-address") != "" {
		vaultConfig = &tmpl.VaultConfig{
			Address: viper.GetString("vault-address"),
			Token:   viper.GetString("vault-token"),
		}
	}
	return vaultConfig
}

func writeResults(results []tmpl.Result, w io.Writer) error {
	for i, r := range results {
		_, err := w.Write(r.Data)
		if err != nil {
			return err
		}
		if i < len(results)-1 || delimiter == "" {
			_, err = w.Write([]byte(delimiter))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
