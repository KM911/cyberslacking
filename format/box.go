package format

import (
	"strings"

	"github.com/gookit/color"
)

// ╭─────────┬──────────╮
// │ _title  │ _content │
// ╰─────────┴──────────╯

func BoxMessage(_title string, _content string, l1 int, l2 int) {
	FormatStringBuilder.WriteString("╭")
	FormatStringBuilder.WriteString(strings.Repeat("─", l1+2))
	FormatStringBuilder.WriteString("┬")
	FormatStringBuilder.WriteString(strings.Repeat("─", l2+2))
	FormatStringBuilder.WriteString("╮\n│ ")
	FormatStringBuilder.WriteString(_title)
	FormatStringBuilder.WriteString(" │ ")
	FormatStringBuilder.WriteString(_content)
	FormatStringBuilder.WriteString(" │\n╰")
	FormatStringBuilder.WriteString(strings.Repeat("─", l1+2))
	FormatStringBuilder.WriteString("┴")
	FormatStringBuilder.WriteString(strings.Repeat("─", l2+2))
	FormatStringBuilder.WriteString("╯\n")
	println(FormatStringBuilder.String())
	FormatStringBuilder.Reset()
}

func BoxError(_title string, _content string) {

	BoxMessage(color.Error.Render(_title), _content, len(_title), len(_content))
}

func BoxWarning(_title string, _content string) {
	BoxMessage(color.Warn.Render(_title), _content, len(_title), len(_content))
}

func BoxInfo(_title string, _content string) {
	BoxMessage(color.Info.Render(_title), _content, len(_title), len(_content))
}

func BoxSuccess(_title string, _content string) {
	BoxMessage(color.Success.Render(_title), _content, len(_title), len(_content))
}

func BoxNote(_title string, _content string) {
	BoxMessage(color.Note.Render(_title), _content, len(_title), len(_content))
}

func BoxExample() {
	BoxError("error", "error")
	BoxWarning("warning", "warning")
	BoxInfo("info", "info")
	BoxSuccess("success", "success")
	BoxNote("note", "note")
}
