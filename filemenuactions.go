package main 

import (
	"github.com/therecipe/qt/widgets"
	//"github.com/therecipe/qt/gui"
	//"github.com/therecipe/qt/core"
	"os"
	"io/ioutil"
	"bufio"
	"fmt"
	"errors"
	"strings"
)

func (window *MainWindow) OkToContinue() bool {
	if window.IsWindowModified() {
		b := widgets.QMessageBox_Warning(window, "Game of Life", "Save changes?",
			widgets.QMessageBox__Yes | widgets.QMessageBox__No | 
			widgets.QMessageBox__Cancel, widgets.QMessageBox__Yes)
		if b == widgets.QMessageBox__Yes {
			window.Save(true)
		} else if b == widgets.QMessageBox__Cancel {
			return false
		}
	}
	return true
}

func (window *MainWindow) SetCurrentFile (filename string) {
	window.SetWindowModified(false)
	window.CurrentFile = filename
	shownfilename := "Untitled"
	if (len(filename) > 0) {
		shownfilename = filename
	}
	window.SetWindowTitle(shownfilename + " - Game of Life")
}

func (window *MainWindow) WriteFile(filename string) error{
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, row := range window.GameState {
		for _, b := range row {
			if b {
				_, err := writer.WriteString("1")
				if err != nil {
					return err
				}
			} else {
				_, err := writer.WriteString("0")
				if err != nil {
					return err
				}
			}
		}
		writer.WriteString("\n")
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (window *MainWindow) LoadFile(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		window.StatusBar().ShowMessage(fmt.Sprintf("%v", err), 0)
		return err
	}
	rows := strings.Split(strings.Trim(string(bytes), "\n"), "\n")
	M := len(rows)
	if M < SceneMin || M > SceneMax {
		err := errors.New(
			fmt.Sprintf(
				"Number of rows %v not within range [%v, %v]",
				M, SceneMin, SceneMax))
		return err
	}
	gamestate := make([][]bool, M, M)
	n := len(rows[0])
	for i, row := range rows {
		N := len(row)
		if N < SceneMin || N > SceneMax {
			err := errors.New(
				fmt.Sprintf(
					"Number of rows %v not within range [%v, %v]",
					M, SceneMin, SceneMax))
			return err
		}
		if N != n {
			err := errors.New("Uneven rows")
			return err
		}	
		gamestate[i] = make([]bool, N, N)
		for j:=0; j < N; j++ {
			r := row[j]
			if r == '0' {
				gamestate[i][j] = false
			} else if r == '1' {
				gamestate[i][j] = true
			} else {
				err := errors.New("Invalid character")
				return err
			}
			
			
		}
	}
	window.GameState = gamestate
	window.SetCurrentFile(filename)

	return nil
}

func New(window *MainWindow) func(bool) {
	return func(bool) {
		if (window.OkToContinue()) {
			//window.Scene.KillCells()
			window.SetCurrentFile("")
		}
	}
}

func Open(window *MainWindow) func(bool) {
	return func(bool) {
		if window.OkToContinue() {
			filename := widgets.QFileDialog_GetOpenFileName(window,
				"Open Game of Life", "", "", "", 0)
				
			if len(filename) > 0 {
				err := window.LoadFile(filename)
				if err != nil {
					window.StatusBar().ShowMessage(fmt.Sprintf("%v", err), 0)
				} else {
					window.StatusBar().ShowMessage("Loaded "+filename, 0)
				}
			}
		}
	}
}


func Save(window *MainWindow) func(bool) {
	return func(checked bool) {
		if len(window.CurrentFile)==0 {
			window.SaveAs(checked)
		} else {
			err := window.WriteFile(window.CurrentFile)
			if err != nil {
				window.StatusBar().ShowMessage(fmt.Sprintf("%v", err), 0)
			}
		}
	}
}

func SaveAs(window *MainWindow) func(bool) {
	return func(bool) {
		filename := widgets.QFileDialog_GetSaveFileName(window,
			"Save Game of Life", "", "", "", 0)
		if len(filename)==0 {
			window.StatusBar().ShowMessage("Empty filename", 0)
		} else {
			err := window.WriteFile(filename)
			if err != nil {
				window.StatusBar().ShowMessage(fmt.Sprintf("%v", err), 0)
			}
		}
	}
}

func Exit(window *MainWindow) func(bool) {
	return func(bool) {
		window.Close()
	}
}