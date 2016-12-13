package elkm1gold

import (
	"bufio"
	"log"
)

var Log *log.Logger

func (e *Driver) listen() {
	scanner := bufio.NewScanner(e.stream)
	for scanner.Scan() {
		s := scanner.Text()
		e.listenChan <- s
	}
	if err := scanner.Err(); err != nil {
		Log.Println("Scanner error", err)
	}
}

func (e *Driver)handle() {
	for {
		select {
		case cmd := <-e.requestChan:
			Log.Print("SEND: ", cmd.Command)
			e.stream.Write([]byte(cmd.Command))
			if cmd.Response != nil {
				r := <-e.listenChan
				Log.Print("RECV: ", cmd.Command)
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
			Log.Println("EVNT[", event[2:4], "]", event)
		}
	}
}