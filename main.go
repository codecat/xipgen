package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: xipgen <ip> [subdomain]")
		return
	}

	sip := os.Args[1]
	sub := ""
	if len(os.Args) >= 3 {
		sub = os.Args[2]
	}
	if sub != "" {
		sub += "."
	}

	bip := net.ParseIP(sip).To4()

	littleEndianIP := binary.LittleEndian.Uint32(bip)
	bigEndianIP := binary.BigEndian.Uint32(bip)

	base36IP := strconv.FormatInt(int64(littleEndianIP), 36)

	color.Red(":: xip.io")
	fmt.Printf("%s%s.xip.io\n", sub, base36IP)
	fmt.Printf("%s%d.%d.%d.%d.xip.io\n", sub, bip[0], bip[1], bip[2], bip[3])
	fmt.Println()

	color.Red(":: nip.io")
	fmt.Printf("%s%d.%d.%d.%d.nip.io\n", sub, bip[0], bip[1], bip[2], bip[3])
	fmt.Printf("%s%d-%d-%d-%d.nip.io\n", sub, bip[0], bip[1], bip[2], bip[3])
	fmt.Printf("%s%08x.nip.io\n", sub, bigEndianIP)
	fmt.Println()

	color.Red(":: sslip.io")
	fmt.Printf("%s%d.%d.%d.%d.sslip.io\n", sub, bip[0], bip[1], bip[2], bip[3])
	fmt.Printf("%s%d-%d-%d-%d.sslip.io\n", sub, bip[0], bip[1], bip[2], bip[3])
	fmt.Println()
}
