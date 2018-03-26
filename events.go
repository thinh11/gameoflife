package main

import (
	//"github.com/therecipe/qt/widgets"
	//"github.com/therecipe/qt/gui"
	//"github.com/therecipe/qt/core"
	"math/rand"
)



/*
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
*/

func WindowAdvance(window *MainWindow) func() {
	return func() {
		window.AdvanceGameState()
		window.Scene.Advance()
		window.SetWindowModified(true)
	}
}

func TimerToggle(window *MainWindow) func(bool) {
	return func(checked bool) {
		if checked {
			window.Timer.Start2()
			window.PauseButton.SetText("Pause")
		} else {
			window.Timer.Stop()
			window.PauseButton.SetText("Start")
		}
	}
}

func Transform(window *MainWindow) func(int) {
	return func(int) {
		s := float64(window.ZoomSlider.Value())/100.0
		r := float64(window.RotateSlider.Value())
		window.View.ResetMatrix()
		window.View.Scale(s, s)
		window.View.Rotate(r)
	}
}

func Randomize(window *MainWindow) func(bool) {
	return func(bool) {
		window.Pause()
		N := len(window.GameState)-2
		M := len(window.GameState[0])-2
		if window.OkToContinue() {
			for i := 1; i <= N; i++  {
				for j := 1; j <= M; j++ {
					if rand.Intn(2)==0 {
						window.GameState[i][j] = &Dead
					} else {
						window.GameState[i][j] = &Live
					}
				}
			}
		}
		window.SetWindowModified(true)
	}
}

func SetTimeInterval(window *MainWindow) func(int) {
	return func(n int) {
		window.Timer.SetInterval(n)
	}
}

func (window *MainWindow) Pause() {
	window.Timer.Stop()
	window.PauseButton.SetChecked(false)
	window.PauseButton.SetText("Start")
}

func (window *MainWindow) Start() {
	window.Timer.Start2()
	window.PauseButton.SetChecked(true)
	window.PauseButton.SetText("Start")
}

func (window *MainWindow) AdvanceGameState() {
	M := len(window.GameState)-2
	N := len(window.GameState[0])-2
	b := window.GameState
	for i:=1; i <= M; i++ {
		for j:=1; j <= N; j++ {
			//window.GameState[i][j] = !window.GameState[i][j]
			live := *b[i-1][j-1] + *b[i-1][j] + *b[i-1][j+1] +
				    *b[i][j-1]   + *b[i][j+1] +
				    *b[i+1][j-1] + *b[i+1][j] + *b[i+1][j+1]
			if live==3 || (live==2 && *b[i][j]==1) {
				ReservedState[i][j] = &Live
			} else {
				ReservedState[i][j] = &Dead
			}
		}
	}
	for i:=1; i <= M; i++ {
		copy(b[i], ReservedState[i])
	}
}