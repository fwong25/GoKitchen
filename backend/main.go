package main

import (
	"db_mgmt"
	"encoding/json"
	"fmt"
	interfaces "interfaces"
	"io"
	"net/http"
	"os"
	"strconv"
)

type ServerConfig struct {
	Name string `json:"Name"`
	Ip   string `json:"Ip"`
	Port int    `json:"Port"`
}

type ServerConfigAll struct {
	backend_server_idx  int
	frontend_server_idx int
	Configs             []ServerConfig `json:"Configs"`
}

func main() {
	db_mgmt.ConnectDB()

	host := readServerConfig()

	http.HandleFunc("/", interfaces.Handler)
	// http.ListenAndServe("192.168.18.27:8000", nil)
	http.ListenAndServe(host, nil)

	db_mgmt.CloseDB()
}

func readServerConfig() string {
	var serverConfigAll ServerConfigAll

	jsonFile, err := os.Open("../ipcfg.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully loaded ipcfg.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &serverConfigAll)

	// for i := 0; i < len(serverConfigAll.Configs); i++ {
	// 	fmt.Println(serverConfigAll.Configs[i])
	// }

	host := serverConfigAll.Configs[0].Ip + ":" + strconv.Itoa(serverConfigAll.Configs[serverConfigAll.backend_server_idx].Port)
	fmt.Println("Backend Server Host: ", host)
	return host
}
