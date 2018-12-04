package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// KeyBinding structs is helpful for not re-writinh the same function over and
// over again. it hold useful values to generate a controls view
type KeyBinding struct {
	View        string
	Handler     func(*gocui.Gui, *gocui.View) error
	Key         interface{}
	Modifier    gocui.Modifier
	Display     string
	Description string
	Vital       bool
}

// generate the gui's controls a.k.a. keybindings
func (gui *Gui) generateKeybindings() error {
	gui.KeyBindings = []*KeyBinding{
		{
			View:        "",
			Key:         gocui.KeyCtrlC,
			Modifier:    gocui.ModNone,
			Handler:     gui.quit,
			Display:     "ctrl + c",
			Description: "Force application to quit",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'q',
			Modifier:    gocui.ModNone,
			Handler:     gui.quit,
			Display:     "q",
			Description: "Quit",
			Vital:       true,
		}, {
			View:        mainViewFeature.Name,
			Key:         gocui.KeyTab,
			Modifier:    gocui.ModNone,
			Handler:     gui.switchMode,
			Display:     "tab",
			Description: "Switch mode",
			Vital:       true,
		}, {
			View:        mainViewFeature.Name,
			Key:         gocui.KeyArrowUp,
			Modifier:    gocui.ModNone,
			Handler:     gui.cursorUp,
			Display:     "↑",
			Description: "Up",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         gocui.KeyArrowDown,
			Modifier:    gocui.ModNone,
			Handler:     gui.cursorDown,
			Display:     "↓",
			Description: "Down",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'k',
			Modifier:    gocui.ModNone,
			Handler:     gui.cursorUp,
			Display:     "k",
			Description: "Up",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'j',
			Modifier:    gocui.ModNone,
			Handler:     gui.cursorDown,
			Display:     "j",
			Description: "Down",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'b',
			Modifier:    gocui.ModNone,
			Handler:     gui.nextBranch,
			Display:     "b",
			Description: "Iterate over branches",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'r',
			Modifier:    gocui.ModNone,
			Handler:     gui.nextRemote,
			Display:     "r",
			Description: "Iterate over remotes",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'e',
			Modifier:    gocui.ModNone,
			Handler:     gui.nextRemoteBranch,
			Display:     "e",
			Description: "Iterate over remote branches",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         's',
			Modifier:    gocui.ModNone,
			Handler:     gui.nextCommit,
			Display:     "s",
			Description: "Iterate over commits",
			Vital:       false,
		},{
			View:        mainViewFeature.Name,
			Key:         's',
			Modifier:    gocui.ModAlt,
			Handler:     gui.prevCommit,
			Display:     "alt + s",
			Description: "Iterate over commits",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'd',
			Modifier:    gocui.ModNone,
			Handler:     gui.openCommitDiffView,
			Display:     "d",
			Description: "Show commit diff",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'c',
			Modifier:    gocui.ModNone,
			Handler:     gui.openCheatSheetView,
			Display:     "c",
			Description: "Controls",
			Vital:       true,
		}, {
			View:        mainViewFeature.Name,
			Key:         gocui.KeyEnter,
			Modifier:    gocui.ModNone,
			Handler:     gui.startQueue,
			Display:     "enter",
			Description: "Start queue",
			Vital:       true,
		}, {
			View:        mainViewFeature.Name,
			Key:         gocui.KeySpace,
			Modifier:    gocui.ModNone,
			Handler:     gui.markRepository,
			Display:     "space",
			Description: "Add to queue",
			Vital:       true,
		}, {
			View:        mainViewFeature.Name,
			Key:         gocui.KeyCtrlSpace,
			Modifier:    gocui.ModNone,
			Handler:     gui.markAllRepositories,
			Display:     "ctrl + space",
			Description: "Add all to queue",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         gocui.KeyBackspace2,
			Modifier:    gocui.ModNone,
			Handler:     gui.unmarkAllRepositories,
			Display:     "backspace",
			Description: "Remove all from queue",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'n',
			Modifier:    gocui.ModNone,
			Handler:     gui.sortByName,
			Display:     "n",
			Description: "Sort repositories by Name",
			Vital:       false,
		}, {
			View:        mainViewFeature.Name,
			Key:         'm',
			Modifier:    gocui.ModNone,
			Handler:     gui.sortByMod,
			Display:     "m",
			Description: "Sort repositories by Modification date",
			Vital:       false,
		}, {
			View:        commitDiffViewFeature.Name,
			Key:         'c',
			Modifier:    gocui.ModNone,
			Handler:     gui.closeCommitDiffView,
			Display:     "c",
			Description: "close/cancel",
			Vital:       true,
		}, {
			View:        commitDiffViewFeature.Name,
			Key:         gocui.KeyArrowUp,
			Modifier:    gocui.ModNone,
			Handler:     gui.commitCursorUp,
			Display:     "↑",
			Description: "Page up",
			Vital:       true,
		}, {
			View:        commitDiffViewFeature.Name,
			Key:         gocui.KeyArrowDown,
			Modifier:    gocui.ModNone,
			Handler:     gui.commitCursorDown,
			Display:     "↓",
			Description: "Page down",
			Vital:       true,
		}, {
			View:        commitDiffViewFeature.Name,
			Key:         'k',
			Modifier:    gocui.ModNone,
			Handler:     gui.commitCursorUp,
			Display:     "k",
			Description: "Page up",
			Vital:       false,
		}, {
			View:        commitDiffViewFeature.Name,
			Key:         'j',
			Modifier:    gocui.ModNone,
			Handler:     gui.commitCursorDown,
			Display:     "j",
			Description: "Page down",
			Vital:       false,
		}, {
			View:        cheatSheetViewFeature.Name,
			Key:         'c',
			Modifier:    gocui.ModNone,
			Handler:     gui.closeCheatSheetView,
			Display:     "c",
			Description: "close/cancel",
			Vital:       true,
		}, {
			View:        errorViewFeature.Name,
			Key:         'c',
			Modifier:    gocui.ModNone,
			Handler:     gui.closeErrorView,
			Display:     "c",
			Description: "close/cancel",
			Vital:       true,
		},
	}
	return nil
}

// set the guis by iterating over a slice of the gui's keybindings struct
func (gui *Gui) keybindings(g *gocui.Gui) error {
	for _, k := range gui.KeyBindings {
		if err := g.SetKeybinding(k.View, k.Key, k.Modifier, k.Handler); err != nil {
			return err
		}
	}
	return nil
}

// the bottom line of the gui is mode indicator and keybindings view. Only the
// important controls (marked as vital) are shown
func (gui *Gui) updateKeyBindingsView(g *gocui.Gui, viewName string) error {
	v, err := g.View(keybindingsViewFeature.Name)
	if err != nil {
		return err
	}
	v.Clear()
	v.BgColor = gocui.ColorWhite
	v.FgColor = gocui.ColorBlack
	v.Frame = false
	fmt.Fprint(v, ws)
	modeLabel := ""
	switch mode := gui.State.Mode.ModeID; mode {
	case FetchMode:
		v.BgColor = gocui.ColorBlue
		v.FgColor = gocui.ColorWhite
		modeLabel = fetchSymbol + ws + bold.Sprint("FETCH")
	case PullMode:
		v.BgColor = gocui.ColorMagenta
		v.FgColor = gocui.ColorWhite
		modeLabel = pullSymbol + ws + bold.Sprint("PULL")
	case MergeMode:
		v.BgColor = gocui.ColorCyan
		v.FgColor = gocui.ColorBlack
		modeLabel = mergeSymbol + ws + black.Sprint(bold.Sprint("MERGE"))
	default:
		modeLabel = "No mode selected"
	}

	fmt.Fprint(v, ws+modeLabel+ws+modeSeperator)

	for _, k := range gui.KeyBindings {
		if k.View == viewName && k.Vital {
			binding := keyBindingSeperator + ws + k.Display + ":" + ws + k.Description + ws
			fmt.Fprint(v, binding)
		}
	}
	return nil
}
