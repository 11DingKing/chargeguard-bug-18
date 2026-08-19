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
func ErrorCode(err error) string {
	switch {
	case errors.Is(err, ErrStationNotFound):
		return "STATION_NOT_FOUND"
	case errors.Is(err, ErrVersionConflict):
		return "VERSION_CONFLICT"
	case err != nil:
		return "STORAGE_ERROR"
	default:
		return ""
	}
}
