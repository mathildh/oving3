//package udp
package main

import (
	. "fmt"
	. "net"
	"time"
)

//var chan_con chan *UDPConn
//var chan_send_addr chan *UDPAddr

func ListenToNetwork(chan_con chan *UDPConn, port string) {

	//Finding the broadcast address using port 
	udp_address, err := ResolveUDPAddr("udp4", port)
	check_if_error(err, "ERROR RESOLVING UDP ADDRESS ON PORT: " + port)
	
	//creates a connection that listens to the broadcast address we just found 
	conn, err := ListenUDP("udp", udp_address)
	check_if_error(err, "ERROR LISTENING TO UDP ON PORT: " + port)
	
	chan_con <- conn

	//defer conn.Close()
	//creating the msg to send 
	var msg []byte = make([]byte, 1500)

	for {

		time.Sleep(100*time.Millisecond)
		//trying to read from UDP channel 
		num_bytes, sender_address, err := conn.ReadFromUDP(msg)
		check_if_error(err, "ERROR READING DATA FROM CONNECTION")
		//Print msg
		if sender_address != nil {

			Println("got message from ", sender_address, " with n = ", num_bytes)

			if num_bytes > 0 {
				Println("from address", sender_address, "got message:", string(msg[0:num_bytes]), num_bytes)
			}
		}
	}

}

func SendToNetwork(chan_con chan *UDPConn,port string) {
	//send_address we found in broadcast addr in Listen 
	receiver_address, err := ResolveUDPAddr("udp4", "129.241.187.255" + port)
	check_if_error(err, "ERROR WHILE RESOLVING UDP ADDRESS")
	

	Println("Muggles are sending a greeting?")

	//Deler socket, collecting connection from channel 
	connection := <-chan_con
	greeting_msg := []byte("Liker du ost? Isåfall, hvilken type?")

	//send constantly 
	for {
		connection.WriteToUDP(greeting_msg, receiver_address)
		time.Sleep(200*time.Millisecond)
	}

}
func check_if_error(err error, error_msg string){
	if err != nil{
		Println("Error of type: " + error_msg)
	}
}

func main() {
	port := ":20002"
	chan_con := make(chan *UDPConn, 1)

	go ListenToNetwork(chan_con,port)
	go SendToNetwork(chan_con,port)

	//Alive msg 
	for {
		time.Sleep(5*time.Second)
		Println("I see you muggles!")
	}
	

}
