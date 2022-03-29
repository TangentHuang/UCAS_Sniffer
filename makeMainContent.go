package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/gopacket"
	"time"
)

type netPacket struct {
	SrcIP    string
	DstIP   string
	Protocol string
	SrcPort  string
	DstPort  string
	Length int
	packet gopacket.Packet
}
//packet的长度
var netPacketLen int



func makeMainContent() fyne.CanvasObject {
	head := container.NewGridWithColumns(4, widget.NewLabel("Source"),
		widget.NewLabel("Destination"),
		widget.NewLabel("Protocol"),
		widget.NewLabel("Length"))
	list := widget.NewList(func() int {
		return netPacketLen
	} ,func() fyne.CanvasObject {
		// !todo 获取信息
		hb := container.NewGridWithColumns(4)
		for i := 0; i < 4; i++ {
			hb.Add(widget.NewLabel(fmt.Sprintf("Field %d", i)))
		}
		return hb
	}, func(id widget.ListItemID, object fyne.CanvasObject) {

	})
	//定时刷新list
	go func() {
		for range time.Tick(2*time.Second){
			list.Refresh()
		}
	}()

	listContent := container.NewBorder(head, nil, nil, nil, list)
	return container.NewVSplit(listContent, container.NewHSplit(widget.NewLabel("11"), widget.NewLabel("22")))
}
