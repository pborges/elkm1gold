package elkm1gold

import (
	"bufio"
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
				r := <-e.listenChan
				log.Print("RECV: ", cmd.Command)
				cmd.Response <- r
			}
		case event := <-e.listenChan:
			switch event[2:4] {
			case "ZC":// ZONE CHANGE UPDATE
				if e.zoneChangeHandler != nil {
					e.zoneChangeHandler(decodeZoneChange(event))
				}
			case "AS":// ARMING STATUS REPORT UPDATE
				if e.armingStatusReportHandler != nil {
					e.armingStatusReportHandler(decodeArmingStatusReport(event))
				}
			}
			log.Println("EVNT[", event[2:4], "]", event)
		}
	}
}