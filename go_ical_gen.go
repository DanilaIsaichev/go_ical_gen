package go_ical_gen

import (
	"fmt"
	"strings"
	"time"
)

// Частота повторений
type frequency string

const (
	daily   = "DAILY"
	weekly  = "WEEKLY"
	monthly = "MONTHLY"
	yearly  = "YEARLY"
)

// Правила повторения
type recurrence_rule struct {
	freq     frequency
	interval int
	until    time.Time
}

type event struct {
	start    time.Time
	rrules   recurrence_rule // никогда или правила
	end      time.Time
	summary  string
	location string
}

func generate_event(event_struct event) string {

	event_ical_str := ""

	// определение повторяется ли событие
	if event_struct.rrules.freq != "" {

		// определение частоты повторения
		switch event_struct.rrules.freq {
		case daily:
			if event_struct.rrules.interval == 0 {
				// каждый день
				start := event_struct.start.Format("20060201T150405")
				end := event_struct.end.Format("20060201T150405")
				until := event_struct.rrules.until.Format("20060201T150405")
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=DAILY;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.summary, `
LOCATION:`, event_struct.location, `
END:VEVENT`)

			} else {
				// каждые несколько дней
				start := event_struct.start.Format("20060102T150405")
				end := event_struct.end.Format("20060102T150405")
				until := event_struct.rrules.until.Format("20060102T150405")
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=DAILY;INTERVAL=`, event_struct.rrules.interval, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.summary, `
LOCATION:`, event_struct.location, `
END:VEVENT`)

			}

		case weekly:
			if event_struct.rrules.interval == 0 {
				// каждую неделю
				start := event_struct.start.Format("20060201T150405")
				end := event_struct.end.Format("20060201T150405")
				until := event_struct.rrules.until.Format("20060201T150405")
				day_of_week := strings.ToUpper(event_struct.start.Weekday().String()[:2])
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=WEEKLY;BYDAY=`, day_of_week, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.summary, `
LOCATION:`, event_struct.location, `
END:VEVENT`)

			} else {
				// каждые несколько недель
				start := event_struct.start.Format("20060102T150405")
				end := event_struct.end.Format("20060102T150405")
				until := event_struct.rrules.until.Format("20060102T150405")
				day_of_week := strings.ToUpper(event_struct.start.Weekday().String()[:2])
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=WEEKLY;BYDAY=`, day_of_week, `;INTERVAL=`, event_struct.rrules.interval, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.summary, `
LOCATION:`, event_struct.location, `
END:VEVENT`)
			}

		case monthly:
			if event_struct.rrules.interval == 0 {
				// каждый месяц
				start := event_struct.start.Format("20060201T150405")
				end := event_struct.end.Format("20060201T150405")
				until := event_struct.rrules.until.Format("20060201T150405")
				day_of_month := event_struct.start.Day()
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=MONTHLY;BYMONTHDAY=`, day_of_month, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.summary, `
LOCATION:`, event_struct.location, `
END:VEVENT`)

			} else {
				// каждые несколько месяцев
				start := event_struct.start.Format("20060102T150405")
				end := event_struct.end.Format("20060102T150405")
				until := event_struct.rrules.until.Format("20060102T150405")
				day_of_month := event_struct.start.Day()
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=MONTHLY;BYMONTHDAY=`, day_of_month, `;INTERVAL=`, event_struct.rrules.interval, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.summary, `
LOCATION:`, event_struct.location, `
END:VEVENT`)

			}

		case yearly:
			if event_struct.rrules.interval == 0 {
				// каждый год
				start := event_struct.start.Format("20060201T150405")
				end := event_struct.end.Format("20060201T150405")
				until := event_struct.rrules.until.Format("20060201T150405")
				day_of_year := event_struct.start.YearDay()
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=YEARLY;BYYEARDAY=`, day_of_year, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.summary, `
LOCATION:`, event_struct.location, `
END:VEVENT`)

			} else {
				// каждые несколько лет
				start := event_struct.start.Format("20060102T150405")
				end := event_struct.end.Format("20060102T150405")
				until := event_struct.rrules.until.Format("20060102T150405")
				day_of_year := event_struct.start.YearDay()
				event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
RRULE:FREQ=YEARLY;BYYEARDAY=`, day_of_year, `;INTERVAL=`, event_struct.rrules.interval, `;UNTIL=`, until, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.summary, `
LOCATION:`, event_struct.location, `
END:VEVENT`)

			}

		}
	} else {
		// не повторяется
		start := event_struct.start.Format("20060102T150405")
		end := event_struct.end.Format("20060102T150405")
		event_ical_str = fmt.Sprint(`BEGIN:VEVENT
DTSTART;TZID=Europe/Moscow:`, start, `
DTEND;TZID=Europe/Moscow:`, end, `
SUMMARY:`, event_struct.summary, `
LOCATION:`, event_struct.location, `
END:VEVENT`)

	}

	return event_ical_str
}
