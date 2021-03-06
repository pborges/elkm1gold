package elkm1gold

import (
	"io"
	"io/ioutil"
	"log"
)

func init() {
	Log = log.New(ioutil.Discard, "[elkm1gold] ", log.LstdFlags | log.Lshortfile)
}

func NewDriver(stream io.ReadWriter) (e *Driver) {
	e = new(Driver)
	e.stream = stream
	e.requestChan = make(chan elkRequest)
	e.listenChan = make(chan string)
	go e.listen()
	go e.handle()
	return
}

type Driver struct {
	stream                    io.ReadWriter
	requestChan               chan elkRequest
	listenChan                chan string
	zoneChangeHandler         func(ZoneChangeUpdate)
	armingStatusReportHandler func(ArmingStatusReport)
}