package logger

import "time"

type Options struct {
	FileLocation       string        `json:"fileLocation"`
	FileName           string        `json:"fileName"`
	FileTdrLocation    string        `json:"fileTdrLocation"`
	FileMaxAge         time.Duration `json:"fileMaxAge"`
	Stdout             bool          `json:"stdout"`
	MaskingLogJsonPath string        `json:"maskingLogJsonPath"`
	PublishLog         bool          `json:"publishLog"`
}

type PublishOption struct {
	InstId       string        `json:"instId"`
	PublishLogTo string        `json:"publishLogTo"`
	Timeout      time.Duration `json:"timeout"`
	DebugMode    bool          `json:"debugMode"`
	SkipTLS      bool          `json:"skipTLS"`
}
