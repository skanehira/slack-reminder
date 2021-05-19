package survey

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
)

func askOnetime() (When, error) {
	qs := []*survey.Question{
		qsDate,
		qsHour,
	}

	var onetime Onetime
	err := survey.Ask(qs, &onetime)
	return onetime, err
}

func askEveryDay() When {
	return &RepeatEveryDay{}
}

func askEveryWeek() (When, error) {
	everyWeek := &RepeatEveryWeek{}
	var dayOfWeek string
	s := &survey.Select{
		Message: "What day of week",
		Options: []string{
			Weekday, Weekend, Choice,
		},
	}

	if err := survey.AskOne(s, &dayOfWeek, survey.WithValidator(survey.Required)); err != nil {
		return nil, err
	}

	switch dayOfWeek {
	case Weekday, Weekend:
		everyWeek.Day = dayOfWeek
	case Choice:
		var days []string
		s := &survey.MultiSelect{
			Message: "Choice days",
			Options: Days,
		}

		if err := survey.AskOne(s, &days, survey.WithValidator(survey.Required)); err != nil {
			return nil, err
		}
		everyWeek.Day = strings.Join(days, ", ")
	}
	return everyWeek, nil
}

func askEveryOtherWeek() (When, error) {
	everyOtherWeek := &RepeatEveryOtherWeek{}
	s := &survey.Select{
		Message: "Choice day",
		Options: Days,
	}

	if err := survey.AskOne(s, &everyOtherWeek.Day, survey.WithValidator(survey.Required)); err != nil {
		return nil, err
	}

	return everyOtherWeek, nil
}

func askEveryMonth() (When, error) {
	everyMonth := &RepeatEveryMonth{}
	if err := survey.Ask([]*survey.Question{qsDay}, &everyMonth.Day, survey.WithValidator(survey.Required)); err != nil {
		return nil, err
	}
	return everyMonth, nil
}

func askEveryYear() (When, error) {
	qs := []*survey.Question{
		qsMonth,
		qsDay,
	}

	everyYear := &RepeatEveryYear{}
	err := survey.Ask(qs, everyYear)
	return everyYear, err
}

func askRepetition() (When, error) {
	var (
		RepetitionKind string
	)

	s := &survey.Select{
		Message: "What kind of repetition",
		Options: []string{
			EveryDay, EveryWeek, EveryOtherWeek, EveryMonth, EveryYear,
		},
	}

	if err := survey.AskOne(s, &RepetitionKind, survey.WithValidator(survey.Required)); err != nil {
		return nil, err
	}

	var (
		err  error
		when When
	)
	switch RepetitionKind {
	case EveryDay:
		when = askEveryDay()
	case EveryWeek:
		when, err = askEveryWeek()
	case EveryOtherWeek:
		when, err = askEveryOtherWeek()
	case EveryMonth:
		when, err = askEveryMonth()
	case EveryYear:
		when, err = askEveryYear()
	}

	if err != nil {
		return nil, err
	}

	if err := survey.Ask([]*survey.Question{qsHour}, when); err != nil {
		return nil, err
	}
	return when, nil
}

func Ask() error {
	var (
		answer Answer
		times  string
	)

	s := &survey.Select{
		Message: "Kind of remind",
		Options: []string{"repetition", "onetime"},
	}
	if err := survey.AskOne(s, &times, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	if times == "onetime" {
		onetime, err := askOnetime()
		if err != nil {
			return err
		}
		answer.When = onetime
	} else {
		repetition, err := askRepetition()
		if err != nil {
			return err
		}
		answer.When = repetition
	}

	qs := []*survey.Question{
		{
			Name: "Destination",
			Prompt: &survey.Input{
				Message: "@someone or #channel or me",
				Default: "me",
			},
		},
		{
			Name: "Message",
			Prompt: &survey.Input{
				Message: "Message",
			},
			Validate: survey.Required,
		},
		{
			Name: "Clipboard",
			Prompt: &survey.Confirm{
				Message: "Copy reminder to clipboard?",
				Default: true,
			},
		},
	}

	err := survey.Ask(qs, &answer)
	if err != nil {
		return err
	}

	if answer.Clipboard {
		clipboard.WriteAll(answer.String())
		fmt.Printf("The following reminder was copied to the clipboard: %s\n", answer)
	} else {
		fmt.Println(answer)
	}

	return nil
}
