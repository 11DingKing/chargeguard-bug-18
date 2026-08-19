package charging

import "errors"

var ErrStationNotFound = errors.New("station not found")
var ErrVersionConflict = errors.New("version conflict")

func ReportHazard(station string) error {
	if station == "deleted" {
		return ErrStationNotFound
	}
	if station == "stale" {
		return ErrVersionConflict
	}
	return nil
}
