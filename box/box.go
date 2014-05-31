package box

import "bytes"
import "io"
import "regexp"
import "strings"

import "github.com/mattn/go-runewidth"

var ansiCutter = regexp.MustCompile("\x1B[^a-zA-Z]*[A-Za-z]")

func Print(nodes []string, width int, out io.Writer) {
	maxLen := 1
	for _, finfo := range nodes {
		length := runewidth.StringWidth(ansiCutter.ReplaceAllString(finfo, ""))
		if length > maxLen {
			maxLen = length
		}
	}
	nodePerLine := (width - 1) / (maxLen + 1)
	if nodePerLine <= 0 {
		nodePerLine = 1
	}
	nlines := (len(nodes) + nodePerLine - 1) / nodePerLine

	lines := make([]bytes.Buffer, nlines)
	for i, finfo := range nodes {
		lines[i%nlines].WriteString(finfo)
		lines[i%nlines].WriteString(
			strings.Repeat(" ", maxLen+1-
				runewidth.StringWidth(ansiCutter.ReplaceAllString(finfo, ""))))
	}
	for _, line := range lines {
		io.WriteString(out, strings.TrimSpace(line.String()))
		io.WriteString(out, "\n")
	}
}
