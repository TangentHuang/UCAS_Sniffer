package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

var netPacketList []netPacket

var handle *pcap.Handle

func captureWithBPF(a fyne.App, w fyne.Window) {
	filename, _ := PcapFilePath.Get()
	BPFstr, _ := BPFString.Get()
	if filename == "" {
		if handle != nil {
			handle.Close()
		}
		netPacketList = netPacketList[0:0]
		netPacketLen = 0
		name, _ := INTERFACENAME.Get()
		var err error
		handle, err = pcap.OpenLive(name, 1600, true, time.Second)
		handle.SetBPFFilter(BPFstr)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			//fmt.Println(packet.Dump())
			analysisPacket(packet)
		}
	} else {
		if handle != nil {
			handle.Close()
		}
		netPacketList = netPacketList[0:0]
		netPacketLen = 0
		filePath, _ := PcapFilePath.Get()
		log.Println(filePath)
		var err error
		handle, err = pcap.OpenOffline(filePath)
		handle.SetBPFFilter(BPFstr)
		if err != nil {
			//dialog.ShowError(err, w)
			return
		}
		netPacketList = netPacketList[0:0]
		netPacketLen = 0
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			//fmt.Println(packet.Dump())
			analysisPacket(packet)
		}
	}
}

func parsePcapFile(a fyne.App, w fyne.Window) {
	filePath, _ := PcapFilePath.Get()
	log.Println(filePath)
	var err error
	handle, err = pcap.OpenOffline(filePath)
	if err != nil {
		//dialog.ShowError(err, w)
		return
	}
	netPacketList = netPacketList[0:0]
	netPacketLen = 0
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//fmt.Println(packet.Dump())
		analysisPacket(packet)
	}
}

func startCapture(a fyne.App, w fyne.Window) {
	name, _ := INTERFACENAME.Get()
	var err error
	handle, err = pcap.OpenLive(name, 1600, true, time.Second)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//fmt.Println(packet.Dump())
		analysisPacket(packet)
	}
}

func stopCapture() {
	handle.Close()
	log.Println("stop Capture")
}

func analysisPacket(packet gopacket.Packet) {
	var onePacket netPacket
	//解析IP层
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		onePacket.SrcIP = ip.SrcIP.String()
		onePacket.DstIP = ip.DstIP.String()
		onePacket.Protocol = ip.Protocol.String()
		onePacket.Length = len(packet.Data())
		onePacket.packet = packet
		defer func() {
			netPacketLen += 1
			netPacketList = append(netPacketList, onePacket)
		}()
	}
	//解析tcp
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)
		onePacket.SrcPort = tcp.SrcPort.String()
		onePacket.DstPort = tcp.DstPort.String()
	}
	//解析upd
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		fmt.Println("TCP layer detected.")
		udp, _ := udpLayer.(*layers.UDP)
		onePacket.SrcPort = udp.SrcPort.String()
		onePacket.DstPort = udp.DstPort.String()
	}
}
