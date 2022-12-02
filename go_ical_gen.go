package go_ical_gen

import (
	"fmt"
	"strings"
	"time"
)

// Частота повторений
type Frequency string

const (
	Daily   = "DAILY"
	Weekly  = "WEEKLY"
	Monthly = "MONTHLY"
	Yearly  = "YEARLY"
)

// Правила повторения
type Recurrence_rule struct {
	Freq     Frequency
	Interval int
	Until    time.Time
}

type Event struct {
	Start    time.Time
	Rrules   Recurrence_rule // никогда или правила
	End      time.Time
	Summary  string
	Location string
}

func Generate_event(event_struct Event) string {

	event_ical_str := ""

	// определение повторяется ли событие
	if event_struct.Rrules.Freq != "" {

		// определение частоты повторения
		switch event_struct.Rrules.Freq {
		case Daily:
			if event_struct.Rrules.Interval == 0 {
				// каждый день
				start := event_struct.Start.Format("20060201T150405")
				end := event_struct.End.Format("20060201T150405")
				until := event_struct.Rrules.Until.Format("20060201T150405")
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=DAILY;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
END:VEVENT`)

			} else {
				// каждые несколько дней
				start := event_struct.Start.Format("20060102T150405")
				end := event_struct.End.Format("20060102T150405")
				until := event_struct.Rrules.Until.Format("20060102T150405")
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=DAILY;INTERVAL=`, event_struct.Rrules.Interval, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
END:VEVENT`)

			}

		case Weekly:
			if event_struct.Rrules.Interval == 0 {
				// каждую неделю
				start := event_struct.Start.Format("20060201T150405")
				end := event_struct.End.Format("20060201T150405")
				until := event_struct.Rrules.Until.Format("20060201T150405")
				day_of_week := strings.ToUpper(event_struct.Start.Weekday().String()[:2])
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=WEEKLY;BYDAY=`, day_of_week, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
END:VEVENT`)

			} else {
				// каждые несколько недель
				start := event_struct.Start.Format("20060102T150405")
				end := event_struct.End.Format("20060102T150405")
				until := event_struct.Rrules.Until.Format("20060102T150405")
				day_of_week := strings.ToUpper(event_struct.Start.Weekday().String()[:2])
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=WEEKLY;BYDAY=`, day_of_week, `;INTERVAL=`, event_struct.Rrules.Interval, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
END:VEVENT`)
			}

		case monthly:
			if event_struct.Rrules.Interval == 0 {
				// каждый месяц
				start := event_struct.Start.Format("20060201T150405")
				end := event_struct.End.Format("20060201T150405")
				until := event_struct.Rrules.Until.Format("20060201T150405")
				day_of_month := event_struct.Start.Day()
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=MONTHLY;BYMONTHDAY=`, day_of_month, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
END:VEVENT`)

			} else {
				// каждые несколько месяцев
				start := event_struct.Start.Format("20060102T150405")
				end := event_struct.End.Format("20060102T150405")
				until := event_struct.Rrules.Until.Format("20060102T150405")
				day_of_month := event_struct.Start.Day()
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=MONTHLY;BYMONTHDAY=`, day_of_month, `;INTERVAL=`, event_struct.Rrules.Interval, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
END:VEVENT`)

			}

		case yearly:
			if event_struct.Rrules.Interval == 0 {
				// каждый год
				start := event_struct.Start.Format("20060201T150405")
				end := event_struct.End.Format("20060201T150405")
				until := event_struct.Rrules.Until.Format("20060201T150405")
				day_of_year := event_struct.Start.YearDay()
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=YEARLY;BYYEARDAY=`, day_of_year, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
END:VEVENT`)

			} else {
				// каждые несколько лет
				start := event_struct.Start.Format("20060102T150405")
				end := event_struct.End.Format("20060102T150405")
				until := event_struct.Rrules.Until.Format("20060102T150405")
				day_of_year := event_struct.Start.YearDay()
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=YEARLY;BYYEARDAY=`, day_of_year, `;INTERVAL=`, event_struct.Rrules.Interval, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
END:VEVENT`)

			}

		}
	} else {
		// не повторяется
		start := event_struct.Start.Format("20060102T150405")
		end := event_struct.End.Format("20060102T150405")
		event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
END:VEVENT`)

	}

	return event_ical_str
}
