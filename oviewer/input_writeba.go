package oviewer

import "github.com/gdamore/tcell/v2"

// setWriteBAMode sets the inputMode to WriteBA.
func (root *Root) setWriteBAMode() {
	input := root.input
	input.reset()
	input.Event = newWriteBAEvent(input.WriteBACandidate)
}

// eventWriteBA represents the writeBA input mode.
type eventWriteBA struct {
	tcell.EventTime
	clist *candidate
	value string
}

// newWriteBAEvent returns writeBAEvent.
func newWriteBAEvent(clist *candidate) *eventWriteBA {
	return &eventWriteBA{clist: clist}
}

// Mode returns InputMode.
func (*eventWriteBA) Mode() InputMode {
	return WriteBA
}

// Prompt returns the prompt string in the input field.
func (*eventWriteBA) Prompt() string {
	return "WriteAndQuit Before:After:"
}

// Confirm returns the event when the input is confirmed.
func (e *eventWriteBA) Confirm(str string) tcell.Event {
	e.value = str
	e.clist.toLast(str)
	e.SetEventNow()
	return e
}

// Up returns strings when the up key is pressed during input.
func (e *eventWriteBA) Up(_ string) string {
	return e.clist.up()
}

// Down returns strings when the down key is pressed during input.
func (e *eventWriteBA) Down(_ string) string {
	return e.clist.down()
}
