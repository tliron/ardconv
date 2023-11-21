package main

import (
	"github.com/tliron/commonlog"
	"github.com/tliron/go-transcribe"
)

const toolName = "ardconv"

var log = commonlog.GetLogger(toolName)

func Transcriber() *transcribe.Transcriber {
	return &transcribe.Transcriber{
		File:        outputPath,
		Format:      outputFormat,
		ForTerminal: pretty,
		Strict:      strict,
		Base64:      base64,
	}
}
