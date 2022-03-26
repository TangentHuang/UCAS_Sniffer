package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.NewWithID("UCAS")
	myWin := myApp.NewWindow("UCAS")
	myWin.SetMainMenu(makeMenu(myApp, myWin))
	myWin.Resize(fyne.NewSize(640*2, 640))
	myWin.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	//主界面头部菜单
	// 定义file菜单，包含了open，save，save as，Quit
	fileOpenItem := fyne.NewMenuItem("Open", nil)
	fileSaveItem := fyne.NewMenuItem("Save", nil)
	fileSaveAsItem := fyne.NewMenuItem("Save AS", nil)
	fileMenu := fyne.NewMenu("File", fileOpenItem, fileSaveItem, fileSaveAsItem)

	// 定义edit菜单，包含Copy，Paste，Cut
	editCopyItem := fyne.NewMenuItem("Copy", nil)
	editPasteItem := fyne.NewMenuItem("Paste", nil)
	editCutItem := fyne.NewMenuItem("Cut", nil)
	editMenu := fyne.NewMenu("Edit", editCopyItem, editPasteItem, editCutItem)

	//定义interface菜单
	interfaceMenu := fyne.NewMenu("Interface")

	// 定义help菜单
	helpDocItem := fyne.NewMenuItem("Doc", nil)
	helpAboutItem := fyne.NewMenuItem("About", nil)
	helpMenu := fyne.NewMenu("Help", helpDocItem, helpAboutItem)

	return fyne.NewMainMenu(
		fileMenu,
		editMenu,
		interfaceMenu,
		helpMenu,
	)
}
