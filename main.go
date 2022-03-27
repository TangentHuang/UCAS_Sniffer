package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	myApp := app.NewWithID("UCAS")
	myWin := myApp.NewWindow("UCAS")

	interfaceCard := widget.NewLabel(getINTERFACENAME())

	myWin.SetMainMenu(makeMenu(myApp, myWin))
	myWin.SetContent(interfaceCard)
	go func() {
		for range time.Tick(time.Second) {
			interfaceCard.SetText(getINTERFACENAME())
			fmt.Println(getINTERFACENAME())
		}
	}()
	myWin.Resize(fyne.NewSize(640*2, 640))

	myWin.ShowAndRun()
}

func makeInterfaceWin(a fyne.App, w fyne.Window) fyne.Window {
	//实现接口选择窗口
	var interfacename string
	label := widget.NewLabel("Select A Interface From The List")
	hbox := container.NewHBox(label)
	interfaceNameList := getInterfaceList()
	interfaceWindow := a.NewWindow("Interface Choice")
	//设置list
	interfaceList := widget.NewList(
		func() int {
			return len(interfaceNameList)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("interface "))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(interfaceNameList[id])
		})
	interfaceList.OnSelected = func(id widget.ListItemID) {
		interfacename = interfaceNameList[id]
		label.SetText(getInterfaceDetails(interfaceNameList[id]))
	}
	interfaceList.OnUnselected = func(id widget.ListItemID) {
		label.SetText("lect An Item From The List")
	}
	interfaceList.Select(125)

	choiceButton := widget.NewButton("confirm", func() {
		setINTERFACENAME(interfacename)
		interfaceWindow.Close()
	})
	interfaceWindow.SetContent(container.NewHSplit(interfaceList, container.NewVBox(hbox, choiceButton)))
	interfaceWindow.Resize(fyne.NewSize(640*2, 640))
	return interfaceWindow

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
	interfaceMenu := fyne.NewMenuItem("Interface", func() {
		interfaceWindow := makeInterfaceWin(a, w)
		interfaceWindow.Show()

	})

	// 定义help菜单
	helpDocItem := fyne.NewMenuItem("Doc", nil)
	helpAboutItem := fyne.NewMenuItem("About", nil)
	helpMenu := fyne.NewMenu("Help", helpDocItem, helpAboutItem)

	return fyne.NewMainMenu(
		fileMenu,
		editMenu,
		fyne.NewMenu("interfaceMenu", interfaceMenu),
		helpMenu,
	)
}
