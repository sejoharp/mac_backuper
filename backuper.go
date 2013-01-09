package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

type Params struct {
	MacAddress, Domain, ClientBroadcast, User string
	TimeoutForWakeup                          int
}

func main() {
	params := parseParams()
	checkParams(params)
	fmt.Println(params)

}

func ping(domain string, port string) bool {
	conn, err := net.DialTimeout("tcp", domain+":"+port, 500*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func isAfpOnline(domain string) bool {
	return ping(domain, "548")
}

func isServerOnline(domain string) bool {
	return ping(domain, "22")
}

func executeCommand(command string) string {
	output, err := exec.Command(command).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}

func isBackupInProcess() bool {
	return executeCommand("ps -ax | grep /CoreServices/[b]ackupd 2>&1") != ""
}

func startTimeMachine() {
	executeCommand("/System/Library/CoreServices/backupd.bundle/Contents/Resources/backupd-helper")
}

func shutdownServer(user string, domain string) {
	executeCommand("ssh " + user + "@" + domain + " sudo shutdown -p now")
}

func wakeupBackupMachine(macAddress string) {
	executeCommand("/usr/local/Cellar/wol/HEAD/bin/wol " + macAddress)
}

func isClientAtHome(clientBroadcast string) bool {
	return executeCommand("ifconfig | grep "+clientBroadcast) != ""
}

func parseParams() *Params {
	params := new(Params)
	flag.StringVar(&params.MacAddress, "m", "", "macaddress to wakeup the backup server e.g. e4:ae:9f:4c:10:b3")
	flag.StringVar(&params.Domain, "s", "", "adress from the backup server e.g. 192.168.0.1")
	flag.StringVar(&params.ClientBroadcast, "b", "", "broadcastaddress of the network e.g. 192.168.0.255")
	flag.StringVar(&params.User, "u", "", "username to login to backup server")
	flag.IntVar(&params.TimeoutForWakeup, "t", 500, "timeout in milliseconds for wakeing up e.g. 500")
	flag.Usage = usage
	flag.Parse()
	return params
}

func checkParams(params *Params) {
	if params.ClientBroadcast == "" || params.Domain == "" || params.MacAddress == "" || params.TimeoutForWakeup == 0 || params.User == "" {
		usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: mac_backuper  -msbut \n")
	flag.PrintDefaults()
	os.Exit(2)
}
