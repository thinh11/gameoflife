package main

import (
	"github.com/therecipe/qt/widgets"
	//"github.com/therecipe/qt/core"
	"os"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	app := widgets.NewQApplication(len(os.Args), os.Args)
	
	window, _ := NewMainWindow(100, 100, 10, 10)
	
	window.Show()
	
	app.Exec()
}
