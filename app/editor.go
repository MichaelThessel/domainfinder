package app

import (
	"github.com/jroimartin/gocui"
)

type lineEditor struct {
	gocuiEditor gocui.Editor
}

var le lineEditor

// Edit sets up input handling
func (e *lineEditor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch key {
	// Disable line wrapping
	case gocui.KeyEnter:
		return

	// Disable line wrapping (right arrow key at line end wraps too)
	case gocui.KeyArrowRight:
		x, _ := v.Cursor()
		if x >= len(v.ViewBuffer())-1 {
			return
		}

	case gocui.KeyHome:
		v.SetCursor(0, 0)
		return

	case gocui.KeyEnd:
		v.SetCursor(len(v.ViewBuffer())-2, 0)
		return
	}

	e.gocuiEditor.Edit(v, key, ch, mod)
}
