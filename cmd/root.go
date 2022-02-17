package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/quantumsheep/doka/visitor"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Version string
)

var rootCmd = &cobra.Command{
	Use:     "doka",
	Short:   "Extended Dockerfile syntax",
	Version: Version,
	Run:     run,
}

func init() {
	flags := rootCmd.PersistentFlags()
	flags.StringArrayP("var", "v", []string{
		fmt.Sprintf("os=%s", runtime.GOOS),
		fmt.Sprintf("arch=%s", runtime.GOARCH),
	}, "Variables to be used in the Dokafile")
	flags.StringP("file", "f", "Dokafile", "Dokafile path")
	flags.BoolP("full-errors", "e", false, "Enable full error messages")

	viper.SetDefault("author", "quantumsheep <nathanael.dmc@outlook.fr>")
	viper.SetDefault("license", "MIT")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()

	variablesArguments, e := flags.GetStringArray("var")
	if e != nil {
		log.Fatal(e)
	}

	file, e := flags.GetString("file")
	if e != nil {
		log.Fatal(e)
	}

	fullErrors, e := flags.GetBool("full-errors")
	if e != nil {
		log.Fatal(e)
	}

	variables, e := variablesArgumentsToMap(variablesArguments)
	if e != nil {
		log.Fatal(e)
	}

	if _, ok := variables["os"]; !ok {
		variables["os"] = runtime.GOOS
	}

	if _, ok := variables["arch"]; !ok {
		variables["arch"] = runtime.GOARCH
	}

	output, errors := visitor.ParseFile(file, fullErrors, variables)
	if len(errors) > 0 {
		for _, e := range errors {
			fmt.Fprintln(os.Stderr, e)
		}

		os.Exit(1)
	}

	fmt.Println(output)
}

func variablesArgumentsToMap(variablesArguments []string) (map[string]string, error) {
	variables := make(map[string]string)
	rx := regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")

	for _, v := range variablesArguments {
		parts := strings.SplitN(v, "=", 2)

		if !rx.MatchString(parts[0]) {
			return nil, fmt.Errorf("invalid variable name: %s", parts[0])
		}

		if len(parts) == 1 {
			variables[parts[0]] = os.Getenv(parts[0])
		} else if len(parts) == 2 {
			variables[parts[0]] = parts[1]
		} else {
			return nil, fmt.Errorf("invalid variable: '%s', accepted format is 'NAME=VALUE'", v)
		}
	}

	return variables, nil
}
