package main 

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
	"errors"
	"fmt"
)

const (
	SceneMin int = 10
	SceneMax int = 500
)

var ReservedState [][]bool

type MainWindow struct {
	*widgets.QMainWindow
	View *widgets.QGraphicsView
	Scene *Scene
	PauseButton *widgets.QPushButton
	GameState [][]bool
	CurrentFile string
	New func(bool)
	Open func(bool)
	Save func(bool)
	SaveAs func(bool)
	Exit func(bool)
	/*
	Copy func(bool)
	Paste func(bool)
	Cut func(bool)
	Delete func(bool)
	NewAction *widgets.QAction
	OpenAction *widgets.QAction
	SaveAction *widgets.QAction
	SaveAsAction *widgets.QAction
	ExitAction *widgets.QAction
	CopyAction *widgets.QAction
	PasteAction *widgets.QAction
	CutAction *widgets.QAction
	DeleteAction *widgets.QAction
	*/
}

func NewMainWindow(width int, height int, cellwidth int, cellheight int) (*MainWindow, error) {
	if height < SceneMin || height > SceneMax || width < SceneMin || width > SceneMax {
		return nil, errors.New(fmt.Sprintf(
			"Dimension %vx%v out of range [%v, %v]", height, width, SceneMin, SceneMax))
	}
	
	CellWidth = cellwidth
	CellHeight = cellheight
	
	
	view := widgets.NewQGraphicsView(nil)
	scene := NewScene()
	pausebutton := widgets.NewQPushButton2("Start", nil)
	
	gamestate := make([][]bool, height+2, height+2)
	ReservedState = make([][]bool, height+2, height+2)
	
	for i:=0; i < height+2; i++ {
		gamestate[i] = make([]bool, width+2, width+2)
		ReservedState[i] = make([]bool, width+2, width+2)
	}
	
	
	window := &MainWindow{widgets.NewQMainWindow(nil, core.Qt__Widget), 
		view, scene, pausebutton, gamestate, "", nil, nil, nil, nil, nil}
	
	view.SetScene(scene)
	view.SetRenderHints(gui.QPainter__Antialiasing)
	window.SetCentralWidget(view)
	
	pausebutton.SetCheckable(true)
	pausebutton.ConnectClicked(PauseToggle(window))
	
	window.New = New(window)
	window.Open = Open(window)
	window.Save = Save(window)
	window.SaveAs = SaveAs(window)
	window.Exit = Exit(window)
	
	window.SetupScene()
	window.CreateFileMenu()
	window.CreateToolBar()
	window.CreateStatusBar()
	
	
	
	window.SetWindowTitle("Game of Life")
	//window.SetUnifiedTitleAndToolBarOnMac(true)
	
	return window, nil
}

func (window *MainWindow) CreateFileMenu() {
	NewAction := widgets.NewQAction2("New", nil)
	OpenAction := widgets.NewQAction2("Open", nil)
	SaveAction := widgets.NewQAction2("Save", nil)
	SaveAsAction := widgets.NewQAction2("Save as", nil)
	ExitAction := widgets.NewQAction2("Exit", nil)
	/*
	CopyAction := widgets.NewQAction2("Copy", window)
	PasteAction := widgets.NewQAction2("Paste", window)
	CutAction := widgets.NewQAction2("Cut", window)
	DeleteAction := widgets.NewQAction2("Delete", window)
	*/
	NewAction.ConnectTriggered(window.New)
	OpenAction.ConnectTriggered(window.Open)
	SaveAction.ConnectTriggered(window.Save)
	SaveAsAction.ConnectTriggered(window.SaveAs)
	ExitAction.ConnectTriggered(window.Exit)
	
	filemenu := window.MenuBar().AddMenu2("File")
	filemenu.QWidget.AddAction(NewAction)
	filemenu.QWidget.AddAction(OpenAction)
	filemenu.QWidget.AddAction(SaveAction)
	filemenu.QWidget.AddAction(SaveAsAction)
	filemenu.QWidget.AddAction(ExitAction)
}

func (window *MainWindow) CreateToolBar() {
	toolbar := window.AddToolBar3("Toggling")
	toolbar.AddWidget(window.PauseButton)
	toolbar.AddSeparator()
	toolbar.QWidget.AddAction(widgets.NewQAction2("Random", nil))
}

func (window *MainWindow) CreateStatusBar() {
	statusbar := window.StatusBar()
	statusbar.ShowMessage("Normal", 0)
	
}
func (window *MainWindow) SetupScene() {
	M := len(window.GameState)-2
	N := len(window.GameState[0])-2
	window.Scene.SetSceneRect2(0.0, 0.0, float64(CellWidth*N), float64(CellHeight*M))
	//window.Scene.SetBackgroundBrush(gui.NewQBrush3(gui.NewQColor3(100, 200, 200, 255), core.Qt__SolidPattern))
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			cell := NewCell(&window.GameState[i][j])
			window.Scene.AddItem(cell)
			x := float64(CellWidth)*(float64(j)-0.5)
			y := float64(CellHeight)*(float64(i)-0.5)
			cell.SetPos2(x, y)
		}
	}
}


func WindowAdvance(window *MainWindow) func() {
	return func() {
		if window.PauseButton.IsChecked(){
			window.AdvanceGameState()
			window.Scene.Advance()
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

/*
func GetNeighbors(b [][]bool, M int, N int, i int, j int) []bool {
	if i==0 {
		if j==0 {
			return []bool{b[1][0], b[1][1], b[0][1]}
		}
		if j==N-1 {
			return []bool{b[0][N-2], b[1][N-2], b[1][N-1]}
		}
		return []bool{b[0][j-1], b[1][j-1], b[1][j], b[1][j+1], b[0][j+1]}
	}
	if i==M-1 {
		if j==0 {
			return []bool{b[M-1][1], b[M-2][1], b[M-2][0]}
		}
		if j==N-1 {
			return []bool{b[M-2][N-1], b[M-2][N-2], b[M-1][N-2]}
		}
		return []bool{b[M-1][j-1], b[M-1][j+1], b[M-2][j+1], b[M-2][j], b[M-1][j-1]}
	}
	if j==0 {
		return []bool{b[i+1][0], b[i+1][1], b[i][1], b[i-1][1], b[i-1][0]}
	}
	if j==N-1 {
		return []bool{b[i-1][N-1], b[i-1][N-2], b[i][N-2], b[i+1][N-2], b[i+1][N-1]}
	}
	return []bool{b[i-1][j-1], b[i-1][j], b[i-1][j+1], b[i][j-1], b[i][j+1], b[i+1][j-1], b[i+1][j], b[i+1][j+1]}	
}
func Count(b []bool) int {
	live := 0
	for _, c := range b {
		if c {
			live++
		}
	}
	return live
}
*/






