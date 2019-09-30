package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	socket net.Conn
	data   chan []byte
}

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)

		if err != nil {
			client.socket.Close()
			fmt.Println("Disconect from server")
			startClientMode()
			break
		}

		if length > 0 {
			fmt.Println("RECEIVED: " + string(message) + " " + time.Now().Format("15:04:05"))
		}
	}
}

func startClientMode() {
	fmt.Println("Starting client...")

	// config.json has the json content
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	bb, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	doc := make(map[string]interface{})
	if err := json.Unmarshal(bb, &doc); err != nil {
		log.Fatal(err)
	}

	if server, contains := doc["IpAddress"]; contains {
		log.Printf("we have a values to connect: %q\n", server)

	} else {
		log.Println("Json document doesn't have a data for any connect")
	}

	serverAddress := fmt.Sprintf("%v:%v", doc["IpAddress"], doc["Port"])

	fmt.Println("Trying to connect to..." + serverAddress)

	for {

		connection, error := net.Dial("tcp", serverAddress)
		if error != nil {
			fmt.Println("no server connection")
			time.Sleep(time.Second)
			continue
		}

		client := &Client{socket: connection}
		go client.receive()
		for {
			time.Sleep(time.Second * 5)
			message := GetLocalIP()
			connection.Write([]byte(strings.TrimRight("IP "+message, "\n")))
		}
	}

}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	startClientMode()
}
