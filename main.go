package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Init() {
	INTERFACENAME.Set("Please Choice  an Interface!")
}

func main() {
	Init()
	myApp := app.NewWithID("UCAS")
	myWin := myApp.NewWindow("UCAS")

	interfaceCard := widget.NewLabelWithData(INTERFACENAME)
	toolbar := makeToolBar(myApp, myWin)
	BPFentry := makeBPFEntry(myApp, myWin)
	BPFbotton := widget.NewButton("Submit", func() {
		setBPFString(BPFentry.Text)
	})
	myWin.SetMainMenu(makeMenu(myApp, myWin))

	toolsContent := container.New(layout.NewGridLayout(2), container.NewHBox(toolbar, widget.NewLabel("Interface:"), interfaceCard), container.New(layout.NewFormLayout(), widget.NewSeparator(), container.NewBorder(nil, nil, nil, BPFbotton, BPFentry)))
	mainContent := makeMainContent()
	myWin.SetContent(container.NewBorder(container.NewVBox(toolsContent, widget.NewSeparator()), nil, nil, nil, mainContent))

	myWin.Resize(fyne.NewSize(640*2, 640))
	myWin.SetMaster()
	myWin.ShowAndRun()
}

func makeBPFEntry(a fyne.App, w fyne.Window) *widget.Entry {
	BPFEntry := widget.NewEntry()
	BPFEntry.SetPlaceHolder("Please input BPF string")

	return BPFEntry

}

func makeToolBar(a fyne.App, w fyne.Window) *widget.Toolbar {
	startCapture := widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
		//!todo 添加开始抓包的逻辑
		go startCapture(a,w)
		fmt.Println("start capture")
	})
	stopCapture := widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
		//!todo 添加停止抓包的逻辑
		stopCapture()
		fmt.Println("stop capture")
	})
	toolbar := widget.NewToolbar(startCapture, stopCapture, widget.NewToolbarSeparator())
	return toolbar
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
