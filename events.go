package elkm1gold

func (d *Driver) OnZoneChange(f func(ZoneChangeUpdate)) {
	d.zoneChangeHandler = f
}