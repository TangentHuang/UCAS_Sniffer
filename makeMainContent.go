package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/gopacket"
	"strconv"
	"time"
)

type netPacket struct {
	SrcIP    string
	DstIP    string
	Protocol string
	SrcPort  string
	DstPort  string
	Length   int
	packet   gopacket.Packet
}

//packet的长度
var netPacketLen int

func makeMainContent() fyne.CanvasObject {
	packetDetails := widget.NewTextGrid()
	packetDump := widget.NewTextGrid()
	head := container.NewGridWithColumns(4,
		widget.NewLabel("Source"),
		widget.NewLabel("Destination"),
		widget.NewLabel("Protocol"),
		widget.NewLabel("Length"))
	list := widget.NewList(func() int {
		return netPacketLen
	}, func() fyne.CanvasObject {
		hb := container.NewGridWithColumns(4)
		for i := 0; i < 4; i++ {
			hb.Add(widget.NewLabel(fmt.Sprintf("Field %d", i)))
		}
		return hb
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		hb := object.(*fyne.Container)
		lbl0 := hb.Objects[0].(*widget.Label)
		lbl0.SetText(netPacketList[id].SrcIP + ": " + netPacketList[id].SrcPort)
		lbl1 := hb.Objects[1].(*widget.Label)
		lbl1.SetText(netPacketList[id].DstIP + ": " + netPacketList[id].DstPort)
		lbl2 := hb.Objects[2].(*widget.Label)
		lbl2.SetText(netPacketList[id].Protocol)
		lbl3 := hb.Objects[3].(*widget.Label)
		lbl3.SetText(strconv.Itoa(netPacketList[id].Length))
	})
	list.OnSelected = func(id widget.ListItemID) {
		packetDetails.SetText(netPacketList[id].packet.String())
		packetDump.SetText(netPacketList[id].packet.Dump())
	}

	//定时刷新list
	go func() {
		for range time.Tick(2 * time.Second) {
			list.Refresh()
		}
	}()

	listContent := container.NewBorder(head, nil, nil, nil, list)
	return container.NewVSplit(listContent, container.NewVSplit(container.NewScroll(packetDetails), container.NewScroll(packetDump)))
}
