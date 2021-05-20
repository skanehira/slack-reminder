package survey

import (
	"fmt"
)

type When interface {
	Time() string
	String() string
}

type Onetime struct {
	Date string
	Hour string
}

func (o Onetime) Time() string {
	return "at " + o.Hour
}

func (o Onetime) String() string {
	return fmt.Sprintf("%s on %s", o.Time(), o.Date)
}

const (
	EveryDay       = "every day"
	EveryWeek      = "every week"
	EveryOtherWeek = "every other week"
	EveryMonth     = "every month"
	EveryYear      = "every year"
)

type RepeatEveryDay struct {
	Hour string
}

func (rd RepeatEveryDay) Time() string {
	return "at " + rd.Hour
}

func (rd RepeatEveryDay) String() string {
	return rd.Time() + " everyday"
}

// kind of day of week
const (
	Weekday = "weekday"
	Weekend = "Saturday, Sunday"
	Choice  = "Choice"
)

var (
	Days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
)

type RepeatEveryWeek struct {
	Hour string
	Day  string
}

func (rw RepeatEveryWeek) Time() string {
	return "at " + rw.Hour
}

func (rw RepeatEveryWeek) String() string {
	return fmt.Sprintf("%s on every %s", rw.Time(), rw.Day)
}

type RepeatEveryOtherWeek struct {
	Hour string
	Day  string
}

func (rw RepeatEveryOtherWeek) Time() string {
	return "at " + rw.Hour
}

func (rw RepeatEveryOtherWeek) String() string {
	return fmt.Sprintf("%s on every other %s", rw.Time(), rw.Day)
}

type RepeatEveryMonth struct {
	Hour string
	Day  string
}

func (rw RepeatEveryMonth) Time() string {
	return "at " + rw.Hour
}

func (rw RepeatEveryMonth) String() string {
	return fmt.Sprintf("%s on %s every month", rw.Time(), rw.Day)
}

type RepeatEveryYear struct {
	Month string
	Day   string
	Hour  string
}

func (ry RepeatEveryYear) Time() string {
	return "at " + ry.Hour
}

func (ry RepeatEveryYear) String() string {
	return fmt.Sprintf("%s on %s-%s every year", ry.Time(), ry.Month, ry.Day)
}

type Answer struct {
	Clipboard   bool
	Destination string
	Message     string
	When        When
}

func (a Answer) String() string {
	return fmt.Sprintf(`/remind %s "%s" %s`, a.Destination, a.Message, a.When)
}
