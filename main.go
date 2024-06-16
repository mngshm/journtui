package main 

import (
  "github.com/gdamore/tcell/v2"
  "github.com/rivo/tview"
)

var app = tview.NewApplication()
var mainBox = tview.NewBox().
  SetBorder(true).
  SetTitle("service logs")

var searchBox = tview.NewInputField() 

func main(){
  flex := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(searchBox, 2, 1, false).AddItem(mainBox, 0, 2, false)

  searchBox.SetLabel("Search Service: ")
  searchBox.SetDoneFunc(func(key tcell.Key){
    app.Stop()
  })

  flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    if event.Rune() == 113 {
      app.Stop()
    }
    return event
  })

  if err := app.SetRoot(flex, true).SetFocus(searchBox).EnableMouse(true).Run(); err != nil {
    panic(err)
  }
}

