package wfhcounter

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

func getCurrentDate(now time.Time) string {
	return fmt.Sprint(now.Format("02/01/2006"))
}

func getCurrentDay(now time.Time) string {
	return fmt.Sprint(now.Format("Monday"))
}

func getCakeDay(now time.Time) string {
	switch now.Weekday() {
	case 1:
		return "in 2 days. :tada:"
	case 2:
		return "in 1 day. :mario_luigi_dance:"
	case 3:
		return "TODAY! :nyanparrot:"
	case 4:
		return "in 6 days. :facepalm:"
	case 5:
		return "in 5 days. :sad-panda:"
	default:
		return "no where! :fire:"
	}
}

func getWFHLine(now time.Time) string {
	holidays := getPublicHoliday()
	for _, h := range holidays {
		diff := math.Abs(h.Sub(now).Hours())
		if now.After(h) && diff < 24 {
			return "Today is public holiday. Happy holiday team! :blob-wobble-gif:"
		}
	}
	return fmt.Sprintf("WFH Day #%v", getWFHCount(now))
}

func getWFHCount(now time.Time) string {
	// WFH started on March 19th 2020
	start := time.Date(2020, 3, 22, 0, 0, 0, 0, time.Local)
	diffWeeks := math.Floor(now.Sub(start).Hours() / 24 / 7)
	days := diffWeeks*5 + 2 - getHolidayCount(now) + float64(now.Weekday())
	return formatDayToEmoji(fmt.Sprint(days))
}

func getHolidayCount(now time.Time) float64 {
	holidays := getPublicHoliday()
	count := 0.0
	for _, h := range holidays {
		if now.After(h) {
			count++
		}
	}
	return count
}

func getPublicHoliday() []time.Time {
	// Only included the public holidays in 2020. Really hope that we won't need 2021!
	EasterFriday := time.Date(2020, 4, 10, 0, 0, 0, 0, time.Local)
	EasterMonday := time.Date(2020, 4, 13, 0, 0, 0, 0, time.Local)
	QueensBirthday := time.Date(2020, 6, 8, 0, 0, 0, 0, time.Local)
	AFLGrandFinal := time.Date(2020, 9, 25, 0, 0, 0, 0, time.Local)
	MelbourneCup := time.Date(2020, 11, 3, 0, 0, 0, 0, time.Local)
	Christmas := time.Date(2020, 12, 25, 0, 0, 0, 0, time.Local)
	BoxingDay := time.Date(2020, 12, 28, 0, 0, 0, 0, time.Local)
	return []time.Time{
		EasterFriday,
		EasterMonday,
		QueensBirthday,
		AFLGrandFinal,
		MelbourneCup,
		Christmas,
		BoxingDay,
	}
}

func formatDayToEmoji(day string) string {
	result := []string{}
	for _, c := range strings.Split(day, "") {
		switch c {
		case "1":
			result = append(result, ":one:")
		case "2":
			result = append(result, ":two:")
		case "3":
			result = append(result, ":three:")
		case "4":
			result = append(result, ":four:")
		case "5":
			result = append(result, ":five:")
		case "6":
			result = append(result, ":six:")
		case "7":
			result = append(result, ":seven:")
		case "8":
			result = append(result, ":eight:")
		case "9":
			result = append(result, ":nine:")
		case "0":
			result = append(result, ":zero:")
		}
	}
	return strings.Join(result, "")
}

func formatCountDown(start time.Time, end time.Time) string {
	days := math.Ceil(end.Sub(start).Hours() / 24)
	switch {
	case days < 0:
		return ""
	case days == 0:
		return "TODAY! :blob-cheer-gif:"
	case days == 1:
		return "TMR! :blob-student:"
	default:
		return fmt.Sprintf("in %v days. :blob-wobble-gif:", days)
	}
}

func getCountdownLines(now time.Time) string {
	stage1SchoolBackDate := time.Date(2020, 5, 26, 0, 0, 0, 0, time.Local)
	stage2SchoolBackDate := time.Date(2020, 6, 9, 0, 0, 0, 0, time.Local)
	result := "School count downs: \n"
	var addition string
	st1 := formatCountDown(now, stage1SchoolBackDate)
	if st1 != "" {
		addition = addition + fmt.Sprintf("Stage 1 back-to-school is %s\n", st1)
	}
	st2 := formatCountDown(now, stage2SchoolBackDate)
	if st2 != "" {
		addition = addition + fmt.Sprintf("Stage 2 back-to-school is %s", st2)
	}
	if addition == "" {
		return ""
	}
	return result + addition
}

func getMessage() string {
	loc, err := time.LoadLocation("Australia/Melbourne")
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	//set timezone,
	current := time.Now().In(loc)
	return fmt.Sprintf(
		`
		Good morning team! :blob-sun-gif: Today is %v, %v. 
%v
Cake Day is %v
%s
		`,
		getCurrentDate(current),
		getCurrentDay(current),
		getWFHLine(current),
		getCakeDay(current),
		getCountdownLines(current))
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		isDebug = false
	}
	api := slack.New(os.Getenv("OAUTH_KEY"), slack.OptionDebug(isDebug))
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	channelID, timestamp, err := api.PostMessage(os.Getenv("CHANNEL_ID"), slack.MsgOptionText(getMessage(), false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Fprintf(w, "Message successfully sent to channel %s at %s", channelID, timestamp)
}
