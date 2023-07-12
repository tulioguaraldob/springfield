package infologger

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	success   *color.Color
	successBg *color.Color
	errorLog  *color.Color
	errorBg   *color.Color
	infoLog   *color.Color
	infoBg    *color.Color
)

func CreateLog() {
	// Success
	success = color.New(color.FgHiGreen, color.Bold)
	successBg = color.New(color.FgWhite, color.BgGreen, color.Bold)

	// Error
	errorLog = color.New(color.FgHiRed, color.Bold)
	errorBg = color.New(color.FgWhite, color.BgRed, color.Bold)

	// Info
	infoLog = color.New(color.FgHiYellow, color.Bold)
	infoBg = color.New(color.FgBlack, color.BgYellow, color.Bold)
}

func Success(message string) {
	fmt.Printf("%s %s.\n", successBg.Sprintf("| SUCCESS |"), success.Sprintf(message))
}

func Error(message string) {
	fmt.Printf("%s %s.\n", errorBg.Sprintf("| ERROR |"), errorLog.Sprintf(message))
}

func Info(message string) {
	fmt.Printf("%s %s.\n", infoBg.Sprintf("| INFO |"), infoLog.Sprintf(message))
}
