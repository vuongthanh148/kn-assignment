package miscrepo

import (
	"time"

	"github.com/centraldigital/cfw-core-lib/pkg/constant/datetimeconst"
	"github.com/google/uuid"
)

func (mr *miscRepo) GetCurrentDateString() string {
	return time.Now().Format(datetimeconst.DATE_FORMAT)
}

func (mr *miscRepo) NewUUID() string {
	return uuid.New().String()
}

func (mr *miscRepo) GetCurrentDate() time.Time {
	return time.Now()
}
