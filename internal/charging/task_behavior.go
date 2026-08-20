package charging

import (
	"errors"
	"fmt"
)

var ErrStationNotFound = errors.New("station not found")
var ErrVersionConflict = errors.New("version conflict")
var ErrStorageFailure = errors.New("storage failure")

func ReportHazard(station string) error {
	if station == "deleted" {
		return ErrStationNotFound
	}
	if station == "stale" {
		return ErrVersionConflict
	}
	if station == "storage" {
		return fmt.Errorf("report hazard: %w", ErrStorageFailure)
	}
	return nil
}
