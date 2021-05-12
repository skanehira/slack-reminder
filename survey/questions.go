package survey

import (
	"fmt"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
)

// common question
var qsHour = &survey.Question{
	Name: "Hour",
	Prompt: &survey.Input{
		Message: "Hour(HH:MM)",
	},
	Validate: ValidateHour,
}

var qsDate = &survey.Question{
	Name: "Date",
	Prompt: &survey.Input{
		Message: "Date(YYYY-MM-DD)",
	},
	Validate: ValidateDate,
}

var qsDay = &survey.Question{
	Name: "Day",
	Prompt: &survey.Input{
		Message: "Day",
	},
	Validate: ValidateDay,
}

var qsMonth = &survey.Question{
	Name: "Month",
	Prompt: &survey.Input{
		Message: "Month",
	},
	Validate: ValidateMonth,
}

func ValidateDate(ans interface{}) error {
	if !regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`).MatchString(ans.(string)) {
		return fmt.Errorf("invalid date: %s", ans)
	}
	return nil
}

func ValidateHour(ans interface{}) error {
	if !regexp.MustCompile(`^\d{2}:\d{2}$`).MatchString(ans.(string)) {
		return fmt.Errorf("invalid hour: %s", ans)
	}
	return nil
}

func ValidateMonth(ans interface{}) error {
	if !regexp.MustCompile(`^\d{2}$`).MatchString(ans.(string)) {
		return fmt.Errorf("invalid month: %s", ans)
	}
	return nil
}

func ValidateDay(ans interface{}) error {
	if !regexp.MustCompile(`^\d{2}$`).MatchString(ans.(string)) {
		return fmt.Errorf("invalid day: %s", ans)
	}
	return nil
}
