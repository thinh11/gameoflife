package main 

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
)


var (
	Cell__Black *gui.QColor = gui.NewQColor3(0,0,0,255)
	Cell__White *gui.QColor = gui.NewQColor3(255,255,255,255)
	Cell__Gray *gui.QColor = gui.NewQColor3(128,128,128,255)
)

var (
	CellWidth int = 10
	CellHeight int = 10
)

type Cell struct {
	*widgets.QGraphicsItem
	Live **int
}

func NewCell(live **int) *Cell {
	cell := &Cell{widgets.NewQGraphicsItem(nil), live}
	cell.ConnectBoundingRect(CellBoundingRect(cell))
	cell.ConnectShape(CellShape(cell))
	cell.ConnectPaint(Paint(cell))
	cell.ConnectMousePressEvent(CellMousePressEvent(cell))
	cell.ConnectAdvance(CellAdvance(cell))
	
	return cell
}

func (cell *Cell) Change() {
	if **cell.Live != 0 {
		*cell.Live = &Dead
	} else {
		*cell.Live = &Live
	}
	cell.Update(core.NewQRectF())
}


func (cell *Cell) Color() *gui.QColor {
	if **cell.Live == 1{
		return Cell__White
	}
	return Cell__Black
}


func CellBoundingRect(cell *Cell) func() *core.QRectF {
	return func() *core.QRectF {
		w := float64(CellWidth)
		h := float64(CellHeight)
		return core.NewQRectF4(-w/2.0, -h/2.0, w, h)
	}
}

func CellShape(cell *Cell) func() *gui.QPainterPath {
	return func() *gui.QPainterPath{
		path := gui.NewQPainterPath()
		path.AddRect(cell.BoundingRect())
		return path
	}
}

func Paint(cell *Cell) func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget) {
	return func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget) {
		pen := gui.NewQPen3(Cell__Gray)
		painter.SetPen(pen)
		painter.SetBrush(gui.NewQBrush3(cell.Color(), core.Qt__SolidPattern))
		painter.DrawRect(cell.BoundingRect())
	}
}

func CellMousePressEvent(cell *Cell) func(*widgets.QGraphicsSceneMouseEvent) {
	return func (event *widgets.QGraphicsSceneMouseEvent) {
		cell.Change()
	}
}


func CellAdvance(cell *Cell) func(int) {
	return func(n int) { 
		if n > 0 {
			cell.Update(core.NewQRectF())
		}
	}
}
