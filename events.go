package main

import (
	//"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
)

func WheelEvent(window *MainWindow) func(*gui.QWheelEvent) {
	return func(event *gui.QWheelEvent) {
		if event.Modifiers() & core.Qt__ControlModifier != 0 {
			if event.AngleDelta().Y() > 0 {
				window.View.Scale(0.01, 0.01)
			} else {
				window.View.Scale(-0.01, -0.01)
			}
		}
	}
}

func Zoom(window *MainWindow) func(int) {
	return func(n int) {
		s := float64(n)/100.0
		window.View.ResetMatrix()
		window.View.Scale(s, s)
	}
}

func Rotate(window *MainWindow) func(int) {
	return func(n int) {
		window.View.ResetMatrix()
		window.View.Rotate(float64(n))
	}
}

func WindowAdvance(window *MainWindow) func() {
	return func() {
		if window.PauseButton.IsChecked(){
			window.AdvanceGameState()
			window.Scene.Advance()
			window.SetWindowModified(true)
		}
	}
}

func PauseToggle(window *MainWindow) func(bool) {
	return func(checked bool) {
		if checked {
			window.PauseButton.SetText("Pause")
		} else {
			window.PauseButton.SetText("Start")
		}
	}
}

func (window *MainWindow) AdvanceGameState() {
	M := len(window.GameState)-2
	N := len(window.GameState[0])-2
	b := window.GameState
	for i:=1; i <= N; i++ {
		for j:=1; j <= M; j++ {
			//window.GameState[i][j] = !window.GameState[i][j]
			ns := []bool{b[i-1][j-1], b[i-1][j], b[i-1][j+1], 
				b[i][j-1], b[i][j+1], 
				b[i+1][j-1], b[i+1][j], b[i+1][j+1]}	
			live := 0
			for _, n := range ns {
				if n {
					live++
				}
			}
			if b[i][j] {
				if live < 2 || live > 3 {
					ReservedState[i][j] = false
				} else {
					ReservedState[i][j] = true
				}
			} else {
				if live==3 {
					ReservedState[i][j] = true
				}
			}
		}
	}
	for i:=1; i <= N; i++ {
		copy(b[i], ReservedState[i])
	}
}