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
	SceneMax int = 502
)

var (
	Live int = 1
	Dead int = 0
)

var ReservedState [][]*int

type MainWindow struct {
	*widgets.QMainWindow
	View *View
	Scene *Scene
	PauseButton *widgets.QToolButton
	ZoomSlider *widgets.QSlider
	RotateSlider *widgets.QSlider
	TimeSpinBox *widgets.QSpinBox
	Timer *core.QTimer
	GameState [][]*int
	CurrentFile string
	New func(bool)
	Open func(bool)
	Save func(bool)
	SaveAs func(bool)
	Exit func(bool)
}

func NewMainWindow(width int, height int, cellwidth int, cellheight int) (*MainWindow, error) {
	if height+2 < SceneMin || height+2 > SceneMax || width+2 < SceneMin || width+2 > SceneMax {
		return nil, errors.New(fmt.Sprintf(
			"Dimension %vx%v out of range [%v, %v]", height, width, SceneMin, SceneMax))
	}
	
	CellWidth = cellwidth
	CellHeight = cellheight
	
	ReservedState = make([][]*int, SceneMax, SceneMax)
	for i, _ := range ReservedState {
		ReservedState[i] = make([]*int, SceneMax, SceneMax)
		for j, _ := range ReservedState[i] {
			ReservedState[i][j] = &Dead
		}
	}
	
	view := NewView()
	scene := NewScene()
	gamestate := make([][]*int, height+2, height+2)
	for i, _ := range gamestate {
		gamestate[i] = make([]*int, width+2, width+2)
		for j, _ := range gamestate[i] {
			gamestate[i][j] = &Dead
		}
	}
	window := &MainWindow{widgets.NewQMainWindow(nil, core.Qt__Widget), 
		view, scene, nil, nil, nil, nil, nil, gamestate, "", nil, nil, nil, nil, nil}
	
	view.SetScene(scene)
	view.SetRenderHints(gui.QPainter__Antialiasing)
	
	
	
	window.New = New(window)
	window.Open = Open(window)
	window.Save = Save(window)
	window.SaveAs = SaveAs(window)
	window.Exit = Exit(window)
	
	window.SetupScene()
	window.CreateFileMenu()
	window.CreateStatusBar()
	window.CreateWidgets() //also set layout
	
	window.Timer = core.NewQTimer(nil)
	window.Timer.ConnectTimeout(WindowAdvance(window))
	window.Timer.SetInterval(300)
	
	
	window.SetWindowTitle("[*]Untitled - Game of Life")
	window.SetWindowModified(true)
	
	return window, nil
}

func (window *MainWindow) CreateWidgets() {
	widget := widgets.NewQWidget(nil, core.Qt__Widget)
	pausebutton := widgets.NewQToolButton(nil)
	pausebutton.SetText("Start")
	pausebutton.SetCheckable(true)
	pausebutton.ConnectClicked(TimerToggle(window))
	zoomslider := widgets.NewQSlider2(core.Qt__Vertical, nil)
	zoomslider.SetRange(50, 200)
	zoomslider.SetValue(100)
	zoomslider.ConnectValueChanged(Transform(window))
	rotateslider := widgets.NewQSlider2(core.Qt__Horizontal, nil)
	rotateslider.SetRange(-180, 180)
	rotateslider.SetValue(0)
	rotateslider.ConnectValueChanged(Transform(window))
	timespinbox := widgets.NewQSpinBox(nil)
	timespinbox.SetRange(100, 1000)
	timespinbox.SetValue(300)
	timespinbox.ConnectValueChanged(SetTimeInterval(window))
	window.PauseButton = pausebutton
	window.ZoomSlider = zoomslider
	window.RotateSlider = rotateslider
	window.TimeSpinBox = timespinbox
	
	randombutton := widgets.NewQToolButton(nil)
	randombutton.SetText("Randomize")
	randombutton.ConnectClicked(Randomize(window))
	
	toplayout := widgets.NewQHBoxLayout()
	toplayout.AddWidget(pausebutton, 0, 0)
	toplayout.AddWidget(randombutton, 0, 0)
	toplayout.AddStretch(0)
	toplayout.AddWidget(timespinbox, 0, 0)
	gridlayout := widgets.NewQGridLayout2()
	gridlayout.AddLayout(toplayout, 0, 0, 0)
	gridlayout.AddWidget(window.View, 1, 0, 0)
	gridlayout.AddWidget(zoomslider, 1, 1, 0)
	gridlayout.AddWidget(rotateslider, 3, 0, 0)
	widget.SetLayout(gridlayout)
	window.SetCentralWidget(widget)
}

func (window *MainWindow) CreateFileMenu() {
	NewAction := widgets.NewQAction2("New", nil)
	OpenAction := widgets.NewQAction2("Open", nil)
	SaveAction := widgets.NewQAction2("Save", nil)
	SaveAsAction := widgets.NewQAction2("Save as", nil)
	ExitAction := widgets.NewQAction2("Exit", nil)

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



