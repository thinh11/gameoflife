package main

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"os"
)

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)
	
	window, _ := NewMainWindow(50, 50, 10, 10)
	
	window.Show()
	
	timer := core.NewQTimer(nil)
	timer.ConnectTimeout(WindowAdvance(window))
	timer.SetInterval(500)
	timer.Start2()
	
	
	
	app.Exec()
}
