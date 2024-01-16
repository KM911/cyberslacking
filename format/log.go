package format

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
)

// Current Log Level
// Error Warning Info Note Success

var ()

func FileLogger(_src string) {
	logFile, err := os.OpenFile(_src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}
func Error(_msg string) {
	ErrorMessage("Error", _msg)
}
func LogErorr(_msg string) {
	log.Println(color.BgRed.Render("[Error]"), ": ", color.Error.Render(_msg))
}

func ErrorMessage(error string, _msg string) {
	fmt.Println(color.BgRed.Render(error), ": ", color.Error.Render(_msg))
}

func Warning(_msg string) {
	WarningMessage("Warning", _msg)
}
func LogWarning(_msg string) {
	log.Println(color.BgYellow.Render("[Warning]"), ": ", color.Warn.Render(_msg))
}

func WarningMessage(warning string, _msg string) {
	fmt.Println(color.BgYellow.Render(warning), ": ", color.Warn.Render(_msg))
}

func Info(_msg string) {
	InfoMessage("Info", _msg)
}

func LogInfo(_msg string) {
	log.Println(color.BgBlue.Render("[Info]"), ": ", color.Info.Render(_msg))
}

func InfoMessage(info string, _msg string) {
	fmt.Println(color.BgBlue.Render(info), ": ", color.Info.Render(_msg))
}

func Note(_msg string) {
	NoteMessage("Note", _msg)
}
func LogNote(_msg string) {
	log.Println(color.BgHiBlue.Render("[Note]"), ": ", color.Note.Render(_msg))
}

func NoteMessage(_note string, _msg string) {
	fmt.Println(color.BgHiBlue.Render(_note), ": ", color.Note.Render(_msg))
}

func Success(_msg string) {
	SuccessMessage("Success", _msg)
}

func LogSuccess(_msg string) {
	log.Println(color.BgGreen.Render("[Success]"), ": ", color.Success.Render(_msg))
}

func SuccessMessage(success string, _msg string) {
	fmt.Println(color.BgGreen.Render(success), ": ", color.Success.Render(_msg))
}

func LogExample() {
	LogErorr("error")
	LogWarning("warning")
	LogInfo("info")
	LogNote("_note")
	LogSuccess("success")
}

func Example() {
	Error("error")
	Warning("warning")
	Info("info")
	Note("_note")
	Success("success")
}
