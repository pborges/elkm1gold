package elkm1gold

func (d *Driver) OnZoneChange(f func(ZoneChangeUpdate)) {
	d.zoneChangeHandler = f
}

func (d *Driver) OnArmingStatusReport(f func(ArmingStatusReport)) {
	d.armingStatusReportHandler = f
}