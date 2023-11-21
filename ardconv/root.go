package main

import (
	contextpkg "context"

	"github.com/spf13/cobra"
	"github.com/tliron/commonlog"
	"github.com/tliron/exturl"
	"github.com/tliron/go-ard"
	"github.com/tliron/kutil/terminal"
	"github.com/tliron/kutil/util"
)

var logTo string
var verbose int
var inputUrl string
var outputPath string
var inputFormat string
var outputFormat string
var colorize string
var strict bool
var pretty bool
var base64 bool

func init() {
	rootCommand.PersistentFlags().BoolVarP(&terminal.Quiet, "quiet", "q", false, "suppress output")
	rootCommand.PersistentFlags().StringVarP(&logTo, "log", "l", "", "log to file (defaults to stderr)")
	rootCommand.PersistentFlags().CountVarP(&verbose, "verbose", "v", "add a log verbosity level (can be used twice)")
	rootCommand.PersistentFlags().BoolVarP(&commonlog.Trace, "trace", "", false, "add stack trace to log messages")

	rootCommand.Flags().StringVarP(&inputUrl, "input-url", "i", "", "input URL (when empty will read from stdin)")
	rootCommand.Flags().StringVarP(&outputPath, "output-url", "o", "", "output path (when empty will write to stdout)")
	rootCommand.Flags().StringVarP(&inputFormat, "input-format", "n", "yaml", "force input format (\"yaml\", \"json\", \"xjson\", \"xml\", \"cbor\", or \"messagepack\")")
	rootCommand.Flags().StringVarP(&outputFormat, "output-format", "f", "", "force output format (\"yaml\", \"json\", \"xjson\", \"xml\", \"cbor\", \"messagepack\", or \"go\")")
	rootCommand.Flags().StringVarP(&colorize, "colorize", "z", "true", "colorize output (boolean or \"force\")")
	rootCommand.Flags().BoolVarP(&strict, "strict", "y", false, "strict output (for \"yaml\" format only)")
	rootCommand.Flags().BoolVarP(&pretty, "pretty", "p", true, "prettify output")
	rootCommand.PersistentFlags().BoolVarP(&base64, "base64", "", false, "output base64 (for \"cbor\", \"messagepack\" formats)")
}

var rootCommand = &cobra.Command{
	Use:   toolName,
	Short: "Convert between ARD formats",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		util.InitializeColorization(colorize)
		commonlog.Initialize(verbose, logTo)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if outputFormat == "" {
			outputFormat = inputFormat
		}

		Convert(contextpkg.TODO())
	},
}

func Execute() {
	err := rootCommand.Execute()
	util.FailOnError(err)
}

func Convert(context contextpkg.Context) {
	urlContext := exturl.NewContext()
	util.OnExitError(urlContext.Release)

	var url exturl.URL
	var err error
	if inputUrl == "" {
		log.Info("parsing stdin")
		url, err = urlContext.ReadToInternalURLFromStdin(context, inputFormat)
	} else {
		log.Infof("parsing %q", inputUrl)
		url, err = urlContext.NewValidURL(context, inputUrl, nil)
	}
	util.FailOnError(err)

	value, _, err := ard.ReadURL(context, url, inputFormat, true, false)
	util.FailOnError(err)

	/*switch inputFormat {
	case "":
		value = ard.CopyMapsToStringMaps(value)
	}*/

	util.FailOnError(Transcriber().Write(value))
}
