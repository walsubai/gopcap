package main

import (
	"fmt"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {

	var pcapPath string = "/Users/walsubai/Library/CloudStorage/OneDrive-Cisco/Desktop/Work/Log Redactor/PCAP/7junepcap.pcapng"
	pcapF, err := pcap.OpenOffline(pcapPath)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer pcapF.Close()

	var userInput string = "192.168.0.117"
	packets := gopacket.NewPacketSource(pcapF, pcapF.LinkType())

	for packet := range packets.Packets() {
		//fmt.Println(packet)
		var packetString string = packet.String()
		if strings.Contains(packetString, userInput) {
			fmt.Println("match")
		} else {
			fmt.Println("NO MATCH")
		}
		//fmt.Println(packetString)
		//fmt.Println(reflect.TypeOf(packetString).Kind())

	}

}

/*
	//pcapRead := bufio.NewScanner(pcapF)
	var userInput string = "192.168.0.117"
	x := convertStringByte(userInput)
	//fmt.Println(x)

	for pcapRead.Scan() {
		y := pcapRead.Bytes()
		pcapString := string(y[:])
		fmt.Println(pcapString)
		match := cmp.Equal(y, x)
		if match {
			fmt.Println("Aaaaaaaaa")
			fmt.Println("MATCH FOUND:", y, " ", x)

		}

	}

	//fmt.Println(x)
	//fmt.Println(reflect.ValueOf(x).Kind())

}



func convertStringByte(userString string) []byte {

	userByte := []byte(userString)
	return userByte

}

func redactPcap(encodedHex string, pcapFile os.File) {

}

*/
