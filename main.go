package main 

import (
  "github.com/gdamore/tcell/v2"
  "github.com/rivo/tview"
)

var app = tview.NewApplication()
var mainBox = tview.NewBox().
  SetBorder(true).
  SetTitle("Journalctl but TUI")

var searchBox = tview.NewBox().
  SetBorder(true).
  SetTitle("Search service")
  

func main(){
  mainBox.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    if event.Rune() == 113 {
      app.Stop()
    }
    return event
  })

  if err := app.SetRoot(mainBox, true).EnableMouse(true).Run(); err != nil {
    panic(err)
  }
}

