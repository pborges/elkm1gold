package elkm1gold

import (
	"fmt"
	"log"
)

type elkRequest struct {
	Command  string
	Response chan string
}

func (e *Driver)request(cmd string, response bool) (string) {
	r := elkRequest{}
	c := fmt.Sprintf("%02X%s", len(cmd) + 2, cmd)
	r.Command = fmt.Sprintf("%s%02X\r\n", c, crc([]byte(c)))
	if response {
		r.Response = make(chan string)
	}
	e.requestChan <- r
	if response {
		req := <-e.listenChan
		log.Println("RECV[", req[2:4], "]", req)
		return req
	}
	return ""
}
