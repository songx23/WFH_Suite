package wfhcounter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatCountDown(t *testing.T) {
	endDate := time.Date(2020, 7, 14, 0, 0, 0, 0, time.Local)
	startDate := time.Date(2020, 7, 14, 8, 0, 0, 0, time.Local)

	res := formatCountDown(startDate, endDate)
	assert.Equal(t, "TODAY! :blob-cheer-gif:", res)

	startDate = time.Date(2020, 7, 15, 8, 0, 0, 0, time.Local)
	res = formatCountDown(startDate, endDate)
	assert.Equal(t, "", res)

	startDate = time.Date(2020, 7, 13, 8, 0, 0, 0, time.Local)
	res = formatCountDown(startDate, endDate)
	assert.Equal(t, "TMR! :blob-student:", res)

	startDate = time.Date(2020, 7, 12, 8, 0, 0, 0, time.Local)
	res = formatCountDown(startDate, endDate)
	assert.Equal(t, "in :two: days.", res)
}

func TestGetCountDownLines(t *testing.T) {
	loc, err := time.LoadLocation("Australia/Melbourne")
	assert.NoError(t, err)
	now := time.Date(2020, 7, 22, 8, 0, 0, 0, time.Local)
	res := getCountdownLines(now, loc)
	assert.Equal(t, "",
		res)

	now = time.Date(2020, 6, 23, 8, 0, 0, 0, time.Local)
	res = getCountdownLines(now, loc)
	assert.Equal(t, "School count downs: \nTerm 3 back-to-school is in :two::one: days.\n",
		res)

	now = time.Date(2020, 7, 14, 8, 0, 0, 0, time.Local)
	res = getCountdownLines(now, loc)
	assert.Equal(t, "School count downs: \nTerm 3 back-to-school is TODAY! :blob-cheer-gif:\n",
		res)
}
