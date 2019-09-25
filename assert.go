package hamcrest

type Asserter interface {
	That(interface{}, Matcher, ...interface{})
}
type DescriptionWriter interface {
	WriteRune(r rune)
	WriteString(s string)
	WriteStringf(string, ...interface{})
	IncreaseIndent(...int) func()
	DecreaseIndent(...int)
	NewLine()
	String() string
}

type Matcher interface {
	Match(interface{}) MatchResult
	WriteDescription(DescriptionWriter)
}

type MatchResult interface {
	Matched() bool
	WriteFailureReason(DescriptionWriter)
}

func DefaultDescriptionWriter() DescriptionWriter {
	return &FormattingDescriptionWriter{}
}

var makeDescriptionWriter = DefaultDescriptionWriter

// SetDescriptionWriter set the DescriptionWriter for your tests.
// This function allows you to create your own formatting rules for how descriptions are crafted
func SetDescriptionWriter(output func() DescriptionWriter) {
	makeDescriptionWriter = output
}

type TestT interface {
	Fatalf(string, ...interface{})
	Fatal(...interface{})
}

func Assert(t TestT) Asserter {
	return &FatalAsserter{t, makeDescriptionWriter}
}

type FatalAsserter struct {
	TestT
	NewWriter func() DescriptionWriter
}

func (me *FatalAsserter) That(to_match interface{}, m Matcher, args ...interface{}) {
	result := m.Match(to_match)
	if !result.Matched() {
		var description = me.NewWriter()
		description.WriteString("failed asserting that ")
		description.WriteStringf("%#v ", to_match)
		description.IncreaseIndent()
		description.NewLine()
		m.WriteDescription(description)
		reset := description.IncreaseIndent(2)
		description.NewLine()
		description.WriteString("because")
		reset()
		description.NewLine()
		result.WriteFailureReason(description)

		me.TestT.Fatal(description.String())
	}
}
