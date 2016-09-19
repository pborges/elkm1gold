package elkm1gold

import (
	"fmt"
)

func (e *Driver) ArmDisarm(level ArmedStatus, area int, pin int) {
	e.request(fmt.Sprintf("a%s%d%06d00", string(level), area, pin), false)
	return
}

func decodeArmingStatusReport(text string) (a ArmingStatusReport) {
	for i := 0; i < len(a); i++ {
		a[i].Number = i + 1
		a[i].ArmedStatus = ArmedStatus(text[4 + i])
		a[i].ArmUpState = ArmUpState(text[12 + i])
		a[i].AlarmState = AlarmState(text[20 + i])
	}
	return
}

func (e *Driver) Disarm(zone int, pin int) {
	e.ArmDisarm(ARMING_STATUS_DISARMED, zone, pin)
	return
}

func (e *Driver) ArmAway(zone int, pin int) {
	// Im not sure why ARMING_STATUS_ARMED_AWAY does not work
	e.ArmDisarm(ARMING_STATUS_ARM_STEP_TO_NEXT_AWAY_MODE, zone, pin)
	return
}

func (e *Driver) ArmStay(zone int, pin int) {
	// Im not sure why ARMING_STATUS_ARMED_STAY does not work
	e.ArmDisarm(ARMING_STATUS_ARM_STEP_TO_NEXT_STAY_MODE, zone, pin)
	return
}

type ArmingStatusReport [8]ArmingStatusZone
type ArmingStatusZone struct {
	Number      int
	ArmedStatus ArmedStatus
	ArmUpState  ArmUpState
	AlarmState  AlarmState
}

type ArmedStatus byte

const (
	ARMING_STATUS_DISARMED ArmedStatus = '0'
	ARMING_STATUS_ARMED_AWAY ArmedStatus = '1'
	ARMING_STATUS_ARMED_STAY ArmedStatus = '2'
	ARMING_STATUS_ARMED_STAY_INSTANT ArmedStatus = '3'
	ARMING_STATUS_ARMED_TO_NIGHT ArmedStatus = '4'
	ARMING_STATUS_ARMED_TO_NIGHT_INSTANT ArmedStatus = '5'
	ARMING_STATUS_ARMED_TO_VACATION ArmedStatus = '6'
	ARMING_STATUS_ARM_STEP_TO_NEXT_AWAY_MODE ArmedStatus = '7'
	ARMING_STATUS_ARM_STEP_TO_NEXT_STAY_MODE ArmedStatus = '8'
	ARMING_STATUS_FORCE_AWAY ArmedStatus = '9'
	ARMING_STATUS_FORCE_STAY ArmedStatus = ':'
)

func (a ArmedStatus) String() string {
	switch a {
	case ARMING_STATUS_DISARMED:
		return "Disarmed"
	case ARMING_STATUS_ARMED_AWAY:
		return "Armed Away"
	case ARMING_STATUS_ARMED_STAY:
		return "Armed Stay"
	case ARMING_STATUS_ARMED_STAY_INSTANT:
		return "Armed Stay Instant"
	case ARMING_STATUS_ARMED_TO_NIGHT:
		return "Armed to Night"
	case ARMING_STATUS_ARMED_TO_NIGHT_INSTANT:
		return "Armed to Night Instant"
	case ARMING_STATUS_ARMED_TO_VACATION:
		return "Armed to Vacation"
	case ARMING_STATUS_FORCE_AWAY:
		return "ForceArm to Away Mode"
	case ARMING_STATUS_FORCE_STAY:
		return "ForceArm to Stay Mode"
	}
	return "UNKNOWN"
}

type ArmUpState byte

const (
	ARM_UP_STATE_NOT_READY ArmUpState = '0'
	ARM_UP_STATE_READY ArmUpState = '1'
	ARM_UP_STATE_READY_FORCE ArmUpState = '2'
	ARM_UP_STATE_ARMED_EXIT_TIMER ArmUpState = '3'
	ARM_UP_STATE_ARMED ArmUpState = '4'
	ARM_UP_STATE_FORCE_ARMED ArmUpState = '5'
	ARM_UP_STATE_ARMED_BYPASS ArmUpState = '6'
)

func (a ArmUpState) String() string {
	switch a {
	case ARM_UP_STATE_NOT_READY:
		return "Not Ready To Arm"
	case ARM_UP_STATE_READY:
		return "Ready To Arm"
	case ARM_UP_STATE_READY_FORCE:
		return "Ready To Arm, but a zone is violated and can be Force Armed."
	case ARM_UP_STATE_ARMED_EXIT_TIMER:
		return "Armed with Exit Timer working"
	case ARM_UP_STATE_ARMED:
		return "Armed Fully"
	case ARM_UP_STATE_FORCE_ARMED:
		return "Force Armed with a force arm zone violated"
	case ARM_UP_STATE_ARMED_BYPASS:
		return "Armed with a bypass"
	}
	return "UNKNOWN"
}

type AlarmState byte

const (
	ALARM_STATE_NO_ALARM AlarmState = '0'
	ALARM_STATE_ENTRANCE_DELAY AlarmState = '1'
	ALARM_STATE_ABORT_DELAY AlarmState = '2'
	ALARM_STATE_ALARM_FIRE AlarmState = '3'
	ALARM_STATE_ALARM_MEDICAL AlarmState = '4'
	ALARM_STATE_ALARM_POLICE AlarmState = '5'
	ALARM_STATE_ALARM_BURGLAR AlarmState = '6'
	ALARM_STATE_ALARM_AUX1 AlarmState = '7'
	ALARM_STATE_ALARM_AUX2 AlarmState = '8'
	ALARM_STATE_ALARM_AUX3 AlarmState = '9'
	ALARM_STATE_ALARM_AUX4 AlarmState = ':'
	ALARM_STATE_ALARM_CARBON_MONOXIDE AlarmState = ';'
	ALARM_STATE_ALARM_EMERGENCY AlarmState = '<'
	ALARM_STATE_ALARM_FREEZE AlarmState = '='
	ALARM_STATE_ALARM_GAS AlarmState = '>'
	ALARM_STATE_ALARM_HEAT AlarmState = '?'
	ALARM_STATE_ALARM_WATER AlarmState = '@'
	ALARM_STATE_ALARM_FIRE_SUPERVISORY AlarmState = 'A'
	ALARM_STATE_ALARM_FIRE_VERIFY AlarmState = 'B'
)

func (a AlarmState) String() string {
	switch a{
	case ALARM_STATE_NO_ALARM:
		return "No Alarm Active"
	case ALARM_STATE_ENTRANCE_DELAY:
		return "Entrance Delay is Active"
	case ALARM_STATE_ABORT_DELAY:
		return "Alarm Abort Delay Active"
	case ALARM_STATE_ALARM_FIRE:
		return "Fire Alarm"
	case ALARM_STATE_ALARM_MEDICAL:
		return "Medical Alarm"
	case ALARM_STATE_ALARM_POLICE:
		return "Police Alarm"
	case ALARM_STATE_ALARM_BURGLAR:
		return "Burglar Alarm"
	case ALARM_STATE_ALARM_AUX1:
		return "Aux1 Alarm"
	case ALARM_STATE_ALARM_AUX2:
		return "Aux2 Alarm"
	case ALARM_STATE_ALARM_AUX3:
		return "Aux3 Alarm"
	case ALARM_STATE_ALARM_AUX4:
		return "Aux4 Alarm"
	case ALARM_STATE_ALARM_CARBON_MONOXIDE:
		return "CarbonMonoxide Alarm"
	case ALARM_STATE_ALARM_EMERGENCY:
		return "Emergency Alarm"
	case ALARM_STATE_ALARM_FREEZE:
		return "FreezeAlar"
	case ALARM_STATE_ALARM_GAS:
		return "Gas Alarm"
	case ALARM_STATE_ALARM_HEAT:
		return "Heat Alarm"
	case ALARM_STATE_ALARM_WATER:
		return "Water Alarm"
	case ALARM_STATE_ALARM_FIRE_SUPERVISORY:
		return "Fire Supervisory"
	case ALARM_STATE_ALARM_FIRE_VERIFY:
		return "Verify Fire"
	}
	return "UNKNOWN"
}