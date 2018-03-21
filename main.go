package main

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	
	cell1 := NewCell()
	cell1.SetPos2(5.0, 5.0)
	cell2 := NewCell()
	cell2.SetPos2(15.0, 50.0)
	cell3 := NewCell()
	cell3.SetPos2(-50.0, -50.0)
	
	
	
	scene := widgets.NewQGraphicsScene(nil)
	//scene.SetSceneRect2(0.0, 0.0, 100.0, 100.0)
	scene.AddItem(cell1)
	scene.AddItem(cell2)
	scene.AddItem(cell3)
	scene.SetBackgroundBrush(gui.NewQBrush3(gui.NewQColor3(100, 200, 200, 255), core.Qt__SolidPattern))
	scene.SetItemIndexMethod(widgets.QGraphicsScene__BspTreeIndex)
	
	
	view := widgets.NewQGraphicsView(nil)
	view.SetScene(scene)
	view.SetRenderHints(gui.QPainter__Antialiasing)
	
	
	
	view.Show()
	
	widgets.QApplication_Exec()
}