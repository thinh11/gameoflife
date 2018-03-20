package main 

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
)

var (
	Cell__Black CellColor = gui.NewQColor3(0,0,0,255)
	Cell__White CellColor = gui.NewQColor3(255,255,255,255)
	Cell__Gray CellColor = gui.NewQColor3(128,128,128,255)
)


type Cell struct {
	*widgets.QGraphicsItem
	Live bool
}

func NewCell() *Cell {
	cell := Cell{widgets.NewQGraphicsItem(nil), false}
	cell.ConnectBoundingRect(BoundingRect(&cell))
	cell.ConnectShape(Shape(&cell))
	cell.ConnectPaint(Paint(&cell))
	cell.ConnectMousePressEvent(MousePressEvent(&cell))
	return &cell
}
func (cell *Cell) Change() {
	cell.Live = !cell.Live
	cell.Update(core.NewQRectF())
}

func (cell *Cell) Color() *gui.QColor {
	if cell.Live {
		return Cell__White
	}
	return Cell__Black
}


func BoundingRect(cell *Cell) func() *core.QRectF {
	return func() *core.QRectF { return core.NewQRectF4(-5.0, -5.0, 10.0, 10.0) }
}

func Shape(cell *Cell) func() *gui.QPainterPath {
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

func MousePressEvent(cell *Cell) func(*widgets.QGraphicsSceneMouseEvent) {
	return func (event *widgets.QGraphicsSceneMouseEvent) { cell.Change() }
}
