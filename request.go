package elkm1gold

import (
	"fmt"
	"strings"
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
		Log.Println("RECV[", req[2:4], "]", req)
		return req
	}
	return ""
}

func (e *Driver)LcdDisplayTextWithAck(l1 string, l2 string, keypad int) {
	l1 = strings.Replace(l1, "\n", "^", -1)
	l2 = strings.Replace(l2, "\n", "^", -1)
	var t1, t2 string
	t1 = l1[: min(len(l1), 16)]
	t2 = l2[: min(len(l2), 16)]
	e.request(fmt.Sprintf("dm%d%d%d%05d%-16s%-16s00", keypad, 1, 1, 0, t1, t2), false)
}

func (e *Driver)LcdDisplayTextTimeout(l1 string, l2 string, keypad int, timeout int) {
	l1 = strings.Replace(l1, "\n", "^", -1)
	l2 = strings.Replace(l2, "\n", "^", -1)
	var t1, t2 string
	t1 = l1[: min(len(l1), 16)]
	t2 = l2[: min(len(l2), 16)]
	e.request(fmt.Sprintf("dm%d%d%d%05d%-16s%-16s00", keypad, 2, 1, timeout, t1, t2), false)
}

func (e *Driver)LcdClearText(keypad int) {
	e.request(fmt.Sprintf("dm%d%d%d%05d%-16s%-16s00", keypad, 0, 1, 0, "", ""), false)
}
