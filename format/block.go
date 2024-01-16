package format

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

// ╭───────╮
// │ _msg  │
// ├───────┤
// │ str   │
// │ str   │
// │ str   │
// ╰───────╯
//

const (
	TopLeft   = "╭"
	TopMiddle = "┬"
	TopRight  = "╮"

	MiddleLeft   = "├"
	MiddleCenter = "┼"
	MiddleRight  = "┤"

	HLine = "─"
	VLine = "│"

	BottomLeft   = "╰"
	BottomMiddle = "┴"
	BottomRight  = "╯"
)

var (
	FormatStringBuilder = strings.Builder{}
)

func MaxLength(s []string) int {
	currentMax := 0
	for _, v := range s {
		if len(v) > currentMax {
			currentMax = len(v)
		}
	}
	return currentMax
}

func BlockMessage(title string, content []string) {
	//	1: find max length of content
	maxLength := MaxLength(content) + 2

	DrawBlockLine(maxLength, TopLeft, HLine, TopRight)
	//	绘制标题
	DrawBlockLine(1, VLine, TitleMiddleAlign(color.Note.Render(title), maxLength), VLine)
	//	绘制中间
	DrawBlockLine(maxLength, MiddleLeft, HLine, MiddleRight)
	//	绘制内容
	for _, v := range content {
		DrawBlockLine(1, VLine, ContentLeftAlign(v, maxLength), VLine)
	}
	//	绘制底部
	DrawBlockLine(maxLength, BottomLeft, HLine, BottomRight)
	println(FormatStringBuilder.String())
	FormatStringBuilder.Reset()
}

func DrawBlockLine(length int, s1 string, s2 string, s3 string) {
	FormatStringBuilder.WriteString(s1)
	FormatStringBuilder.WriteString(strings.Repeat(s2, length))
	FormatStringBuilder.WriteString(s3)
	FormatStringBuilder.WriteString("\n")
}

func TitleMiddleAlign(ctx string, length int) string {
	ctxLength := ColorTextLength(ctx)
	return fmt.Sprintf("%*s%s%*s", (length-ctxLength)/2, "", ctx, (length - (length-ctxLength)/2 - ctxLength), "")
}

func ContentLeftAlign(ctx string, length int) string {
	ctxLength := len(ctx)
	return fmt.Sprintf("%s%*s", ctx, length-ctxLength, "")
}

func ColorTextLength(ctx string) int {
	return len(ctx) - 11
}
