package main
	
import (
	"github.com/therecipe/qt/widgets"
	//"github.com/therecipe/qt/core"
	//"github.com/therecipe/qt/gui"
)

type Scene struct {
	*widgets.QGraphicsScene
	Cut func(bool)
	Copy func(bool)
	Paste func(bool)
	Delete func(bool)
}

func NewScene() *Scene {
	scene := &Scene{widgets.NewQGraphicsScene(nil), nil, nil, nil, nil}
	return scene
}


func Cut (scene *Scene) func(bool) {
	return func(checked bool) {
		scene.Copy(checked)
		scene.Delete(checked)
	}
}
func Copy (scene *Scene) func(bool) {
	return func(bool) {
	
	}
}
func Paste (scene *Scene) func(bool) {
	return func(bool) {
	
	}
}
func Delete (scene *Scene) func(bool) {
	return func(bool) {
	
	}
}