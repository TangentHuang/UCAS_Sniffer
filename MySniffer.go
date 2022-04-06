package main

import (
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
)

var INTERFACENAME = binding.NewString()
var PcapFilePath = binding.NewString()
var BPFString = binding.NewString()

type Interface struct {
	Name        string //设备名称
	Description string //设备描述信息
	Flags       uint32
	Addresses   []InterfaceAddress //网口的地址信息列表
}

// InterfaceAddress describes an address associated with an Interface.
// Currently, it's IPv4/6 specific.
type InterfaceAddress struct {
	IP        net.IP
	Netmask   net.IPMask // Netmask may be nil if we were unable to retrieve it.
	Broadaddr net.IP     // Broadcast address for this IP may be nil
	P2P       net.IP     // P2P destination address for this IP may be nil
}

var allInterfaceInfo = make(map[string]Interface)

func getInterfaceList() []string {
	var interfaceList []string
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range devices {
		interfaceList = append(interfaceList, d.Name)
		var addressList []InterfaceAddress
		for _, address := range d.Addresses {
			addressList = append(addressList, InterfaceAddress{address.IP, address.Netmask, address.Broadaddr, address.P2P})
		}
		allInterfaceInfo[d.Name] = Interface{d.Name, d.Description, d.Flags, addressList}
	}
	return interfaceList
}

func getInterfaceDetails(interfaceName string) string {
	if interfaceInfo, ok := allInterfaceInfo[interfaceName]; ok {
		var interfaceInfoString string
		interfaceInfoString += "Description: "
		interfaceInfoString += interfaceInfo.Description + "\n"
		for _, address := range interfaceInfo.Addresses {
			interfaceInfoString += "- IP address: "
			interfaceInfoString += address.IP.String() + "\n"
			interfaceInfoString += "- Subnet mask: " + address.Netmask.String() + "\n"
		}
		fmt.Println(interfaceInfoString)
		return interfaceInfoString
	}
	return "??"
}

func setINTERFACENAME(name string) {
	INTERFACENAME.Set(name)
	log.Println("set interface is " + name)
}

func setBPFString(str string) {
	BPFString.Set(str)
	log.Println("set BPFString is " + str)
}
