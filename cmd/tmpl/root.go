package main

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

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
	rootCmd.Flags().StringVarP(&sourcesPath, "source", "s", "", "path to sources")
	rootCmd.Flags().StringVarP(&delimiter, "delimiter", "d", "\n", "template delimiter")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "tmpl.out", "output file path")
	rootCmd.Flags().BoolVarP(&test, "test", "t", false, "print result to stdout")
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

	results, err := tmpl.GenerateFromFiles(parametersPath, sourcesPath)
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

	for i, r := range results {
		_, err := w.Write(bytes.TrimSpace(r.Data))
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
