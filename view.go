package main

import (
	"github.com/therecipe/qt/widgets"
)

type View struct {
	*widgets.QGraphicsView
}

func NewView() *View {
	return  &View{widgets.NewQGraphicsView(nil)}
}

