package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var grid *fyne.Container
var myApp fyne.App
var myWindow fyne.Window

func makeNewUi(maze [][]string) {

	myApp = app.New()
	myWindow = myApp.NewWindow("Maze UI")

	grid = container.NewGridWithColumns(len(maze[0]))

	for _, row := range maze {
		for _, cell := range row {
			var rect *canvas.Rectangle
			if isNotWall(cell) {
				rect = canvas.NewRectangle(color.Black)
			} else {
				rect = canvas.NewRectangle(color.White)
			}
			rect.SetMinSize(fyne.NewSize(20, 20))
			grid.Add(rect)
		}
	}

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Maze"),
		grid,
	))

	myWindow.ShowAndRun()
}

func updateUI(maze [][]string) {
	grid.Objects = nil
	for _, row := range maze {
		for _, cell := range row {
			var rect *canvas.Rectangle
			if isNotWall(cell) {
				rect = canvas.NewRectangle(color.Black)
			} else {
				rect = canvas.NewRectangle(color.White)
			}
			rect.SetMinSize(fyne.NewSize(20, 20))
			grid.Add(rect)
		}
	}
	grid.Refresh()
}

func win_ui() {

}
