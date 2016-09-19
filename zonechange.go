package elkm1gold

import (
	"strconv"
	"strings"
)

type ZoneStatus string

const (
	ZoneStatusNormal = "Normal"
	ZoneStatusTrouble = "Trouble"
	ZoneStatusViolated = "Violated"
	ZoneStatusBypassed = "Bypassed"
)

type ZoneSubStatus string

const (
	ZoneSubStatusUnconfigured = "Unconfigured"
	ZoneSubStatusOpen = "Open"
	ZoneSubStatusEOL = "EOL"
	ZoneSubStatusShort = "Short"
)

type ZoneChangeUpdate struct {
	Number    int
	Status    ZoneStatus
	SubStatus ZoneSubStatus
}

func decodeZoneChange(data string) ZoneChangeUpdate {
	z := ZoneChangeUpdate{}
	z.Number, _ = strconv.Atoi(strings.TrimLeft(data[4:6], "0"))
	switch data[7]{
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
	return z
}