package hamcrest

import (
	"fmt"
	"strings"
)

type FormattingDescriptionWriter struct {
	builder     strings.Builder
	indentLevel int
}

func (me *FormattingDescriptionWriter) IncreaseIndent(level ...int) (resetLevel func()) {
	previous := me.indentLevel
	if len(level) > 0 {
		me.indentLevel += level[0]
	} else {
		me.indentLevel++
	}
	return func() {
		me.indentLevel = previous
	}
}
func (me *FormattingDescriptionWriter) DecreaseIndent(level ...int) {
	if len(level) > 0 {
		me.indentLevel -= level[0]
	} else {
		me.indentLevel--
	}
}
func (me *FormattingDescriptionWriter) NewLine() {
	me.builder.WriteRune('\n')
	for i := 0; i < me.indentLevel; i++ {
		me.builder.WriteString("  ")
	}
}

func (me *FormattingDescriptionWriter) WriteRune(r rune) {
	me.builder.WriteRune(r)
}
func (me *FormattingDescriptionWriter) WriteString(s string) {
	me.builder.WriteString(s)
}
func (me *FormattingDescriptionWriter) WriteStringf(s string, args ...interface{}) {
	me.builder.WriteString(fmt.Sprintf(s, args...))
}
func (me *FormattingDescriptionWriter) String() string {
	return me.builder.String()
}

// type NullDescriptionWriter struct{}

// func (me *NullDescriptionWriter) IncreaseIndent()                            {}
// func (me *NullDescriptionWriter) DecreaseIndent()                            {}
// func (me *NullDescriptionWriter) NewLine()                                   {}
// func (me *NullDescriptionWriter) WriteRune(r rune)                           {}
// func (me *NullDescriptionWriter) WriteString(s string)                       {}
// func (me *NullDescriptionWriter) WriteStringf(s string, args ...interface{}) {}
// func (me *NullDescriptionWriter) String() string                             { return "" }
