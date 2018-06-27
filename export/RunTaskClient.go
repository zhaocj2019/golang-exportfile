package export

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

//RunTaskClient
type RunTaskClient struct {
	StartParams *ModelExport
	Domain      string
	Port        int
}

//Run
func (rtc *RunTaskClient) Run() (err error, reseve string) {
	conn, err := net.Dial("tcp", rtc.Domain+":"+strconv.Itoa(rtc.Port))
	checkError(err)
	defer conn.Close()
	var sendContent []byte
	sendContent, err = json.Marshal(*rtc.StartParams)
	if nil == err {
		conn.Write(sendContent)
	}
	return err, ""
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}
