package main
	
import (
	"github.com/therecipe/qt/widgets"
)

type Scene struct {
	*widgets.QGraphicsScene
}

func NewScene() *Scene {
	return &Scene{widgets.NewQGraphicsScene(nil)}
}

