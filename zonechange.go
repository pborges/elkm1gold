package elkm1gold

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
