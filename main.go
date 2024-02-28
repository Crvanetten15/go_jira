package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

var currentViewIndex int = -1

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()
    // Define titles for each panel
    titles := []string{"Issue Type", "Title", "Details", "Status", "Comments"}

    // Create panel1 with a title
    if v, err := g.SetView("panel1", 0, 0, maxX/2-1, maxY/5-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Title = titles[0] // Set the title for panel1
    }

    // Create panel2 with a title
    if v, err := g.SetView("panel2", maxX/2, 0, maxX-1, maxY/5-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Title = titles[1] // Set the title for panel2
    }

    // Create panel3 with a title
    if v, err := g.SetView("panel3", 0, maxY/5, maxX/2-1, 2*maxY/5-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Title = titles[2] // Set the title for panel3
    }

    // Create panel4 with a title
    if v, err := g.SetView("panel4", maxX/2, maxY/5, maxX-1, 2*maxY/5-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Title = titles[3] // Set the title for panel4
    }

    // Create panel5 with a title
    if v, err := g.SetView("panel5", 0, 2*maxY/5, maxX-1, 3*maxY/5-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Title = titles[4] // Set the title for panel5
    }

    // Automatically focus on the first view on startup
    if currentViewIndex == -1 {
        nextView(g, nil)
    }

    return nil
}


func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}
	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
    views := []string{"panel1", "panel2", "panel3", "panel4", "panel5"}
    currentViewIndex = (currentViewIndex + 1) % len(views)
    nextView, err := g.SetCurrentView(views[currentViewIndex])
    if err != nil {
        return err
    }
    
    // Enable edit mode for the current view
    nextView.Editable = true
    // Optionally, set the view to capture the cursor
    nextView.SetCursor(0, 0)
    
    return nil
}


func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
