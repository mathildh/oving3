//package udp
package main

import (
	"fmt"
	. "net"
	"time"
)

//var chan_con chan *UDPConn
//var chan_send_addr chan *UDPAddr

func ListenToNetwork(chan_con chan *UDPConn) {

	port := ":30000"

	my_baddr_udp, err := ResolveUDPAddr("udp4", port)

	if err != nil {
		fmt.Println("error resolving UDP address on ", port)
		fmt.Println(err)
		return
	}

	conn, err := ListenUDP("udp", my_baddr_udp)

	if err != nil {
		fmt.Println("error listening on UDP port ", port)
		fmt.Println(err)
		return
	}
	chan_con <- conn

	//defer conn.Close()

	var msg []byte = make([]byte, 1500)

	for {

		time.Sleep(100 * time.Millisecond)

		no_bytes, server_address, err := conn.ReadFromUDP(msg)

		if err != nil {
			fmt.Println("error reading data from connection")
			fmt.Println(err)
			return
		}

		if server_address != nil {

			fmt.Println("got message from ", server_address, " with n = ", no_bytes)

			if no_bytes > 0 {
				fmt.Println("from address", server_address, "got message:", string(msg[0:no_bytes]), no_bytes)
			}
		}
	}

}

func SendToNetwork(chan_con chan *UDPConn) {
	send_address, err := ResolveUDPAddr("udp4", "129.241.187.255:30000")

	if err != nil {
		println("ERROR while resolving UDP addr")
	}

	fmt.Println("Muggles are sending a greeting?")

	//Deler socket
	connection := <-chan_con
	greeting_msg := []byte("Liker du ost? Isåfall, hvilken type?")

	if connection == nil {
		fmt.Println("Error trying to write to server")
	}
	for {
		connection.WriteToUDP(greeting_msg, send_address)
		time.Sleep(200 * time.Millisecond)
	}

}

func main() {
	chan_con := make(chan *UDPConn, 1)

	go ListenToNetwork(chan_con)
	go SendToNetwork(chan_con)

	for {
		time.Sleep(5 * time.Second)
		fmt.Println("I see you muggles!")
	}

}
