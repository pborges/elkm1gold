package elkm1gold

import (
	"bufio"
	"strconv"
	"strings"
	"log"
)

func (e *Driver) listen() {
	scanner := bufio.NewScanner(e.stream)
	for scanner.Scan() {
		s := scanner.Text()
		e.listenChan <- s
	}
	if err := scanner.Err(); err != nil {
		log.Println("Scanner error", err)
	}
}

func (e *Driver)handle() {
	for {
		select {
		case cmd := <-e.requestChan:
			log.Print("SEND: ", cmd.Command)
			e.stream.Write([]byte(cmd.Command))
			if cmd.Response != nil {
				cmd.Response <- <-e.listenChan
			}
		case event := <-e.listenChan:
			log.Println("EVNT[", event[2:4], "]", event)
			switch event[2:4] {
			case "ZC":// ZONE CHANGE UPDATE
				if e.zoneChangeHandler != nil {
					z := ZoneChangeUpdate{}
					z.Number, _ = strconv.Atoi(strings.TrimLeft(event[4:6], "0"))
					switch event[7]{
					case '0':
						z.Status = ZoneStatusNormal
						z.SubStatus = ZoneSubStatusUnconfigured
					case '1':
						z.Status = ZoneStatusNormal
						z.SubStatus = ZoneSubStatusOpen
					case '2':
						z.Status = ZoneStatusNormal
						z.SubStatus = ZoneSubStatusEOL
					case '3':
						z.Status = ZoneStatusNormal
						z.SubStatus = ZoneSubStatusShort
					case '5':
						z.Status = ZoneStatusTrouble
						z.SubStatus = ZoneSubStatusOpen
					case '6':
						z.Status = ZoneStatusTrouble
						z.SubStatus = ZoneSubStatusEOL
					case '7':
						z.Status = ZoneStatusTrouble
						z.SubStatus = ZoneSubStatusShort
					case '9':
						z.Status = ZoneStatusViolated
						z.SubStatus = ZoneSubStatusOpen
					case 'A':
						z.Status = ZoneStatusViolated
						z.SubStatus = ZoneSubStatusEOL
					case 'B':
						z.Status = ZoneStatusBypassed
						z.SubStatus = ZoneSubStatusShort
					case 'D':
						z.Status = ZoneStatusBypassed
						z.SubStatus = ZoneSubStatusOpen
					case 'E':
						z.Status = ZoneStatusBypassed
						z.SubStatus = ZoneSubStatusEOL
					case 'F':
						z.Status = ZoneStatusBypassed
						z.SubStatus = ZoneSubStatusShort
					}

					e.zoneChangeHandler(z)
				}
			}
		}
	}
}