package wfhcounter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatCountDown(t *testing.T) {
	endDate := time.Date(2020, 5, 26, 0, 0, 0, 0, time.Local)
	startDate := time.Date(2020, 5, 26, 8, 0, 0, 0, time.Local)

	res := formatCountDown(startDate, endDate)
	assert.Equal(t, "TODAY! :blob-cheer-gif:", res)

	startDate = time.Date(2020, 5, 27, 8, 0, 0, 0, time.Local)
	res = formatCountDown(startDate, endDate)
	assert.Equal(t, "", res)

	startDate = time.Date(2020, 5, 25, 8, 0, 0, 0, time.Local)
	res = formatCountDown(startDate, endDate)
	assert.Equal(t, "TMR! :blob-student:", res)

	startDate = time.Date(2020, 5, 24, 8, 0, 0, 0, time.Local)
	res = formatCountDown(startDate, endDate)
	assert.Equal(t, "in 2 days. :blob-wobble-gif:", res)
}

func TestGetCountDownLines(t *testing.T) {
	now := time.Date(2020, 6, 10, 8, 0, 0, 0, time.Local)
	res := getCountdownLines(now)
	assert.Equal(t, "",
		res)

	now = time.Date(2020, 5, 27, 8, 0, 0, 0, time.Local)
	res = getCountdownLines(now)
	assert.Equal(t, "School count downs: \nStage 2 back-to-school is in 13 days. :blob-wobble-gif:",
		res)

	now = time.Date(2020, 5, 26, 8, 0, 0, 0, time.Local)
	res = getCountdownLines(now)
	assert.Equal(t, "School count downs: \nStage 1 back-to-school is TODAY! :blob-cheer-gif:\nStage 2 back-to-school is in 14 days. :blob-wobble-gif:",
		res)
}
