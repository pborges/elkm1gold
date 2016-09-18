package elkm1gold

import (
	"fmt"
	"errors"
)

func (e *Driver) ArmDisarm(level ArmedStatus, area int, pin int) (error) {
	if level < 0 || level >= 9 {
		return errors.New("level must be between 1-8")
	}
	if area < 1 || area >= 8 {
		return errors.New("zone must be between 1-8")
	}
	e.request(fmt.Sprintf("a%d%d%06d00", level, area, pin), false)
	return nil
}

func decodeArmingStatusReport(text string) (a ArmingStatusReport) {
	for i := 0; i < len(a); i++ {
		a[i].ArmedStatus = ArmedStatus(text[4 + i])
		a[i].ArmUpState = ArmUpState(text[12 + i])
		a[i].AlarmState = AlarmState(text[20 + i])
	}
	return
}

func (e *Driver) Disarm(zone int, pin int) (error) {
	return e.ArmDisarm(ARMING_STATUS_DISARMED, zone, pin)
}

func (e *Driver) ArmAway(zone int, pin int) (error) {
	return e.ArmDisarm(ARMING_STATUS_ARM_STEP_TO_NEXT_AWAY_MODE, zone, pin)
}

func (e *Driver) ArmStay(zone int, pin int) (error) {
	return e.ArmDisarm(ARMING_STATUS_ARM_STEP_TO_NEXT_STAY_MODE, zone, pin)
}

type ArmingStatusReport [8]ArmingStatusZone
type ArmingStatusZone struct {
	ArmedStatus ArmedStatus
	ArmUpState  ArmUpState
	AlarmState  AlarmState
}

type ArmedStatus byte

const (
	ARMING_STATUS_DISARMED ArmedStatus = 0
	ARMING_STATUS_ARMED_AWAY ArmedStatus = 1
	ARMING_STATUS_ARMED_STAY ArmedStatus = 2
	ARMING_STATUS_ARMED_STAY_INSTANT ArmedStatus = 3
	ARMING_STATUS_ARMED_TO_NIGHT ArmedStatus = 4
	ARMING_STATUS_ARMED_TO_NIGHT_INSTANT ArmedStatus = 5
	ARMING_STATUS_ARMED_TO_VACATION ArmedStatus = 6
	ARMING_STATUS_ARM_STEP_TO_NEXT_AWAY_MODE ArmedStatus = 7
	ARMING_STATUS_ARM_STEP_TO_NEXT_STAY_MODE ArmedStatus = 8
)

func (a ArmedStatus) String() string {
	switch a {
	case 0:
		return "Disarmed"
	case 1:
		return "Armed Away"
	case 2:
		return "Armed Stay"
	case 3:
		return "Armed Stay Instant"
	case 4:
		return "Armmed to Night"
	case 5:
		return "Armed to Night Instant"
	case 6:
		return "Armed to Vacation"
	}
	return "UNKNOWN"
}

type ArmUpState byte
type AlarmState byte

