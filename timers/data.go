package timers

import "time"

type TimerTemplate struct {
	Name          string
	TimerType     int
	Duration      time.Duration
	TemplateLabel string
	ShortBreak    *string
	LongBreak     *string
}

type Timer struct {
	Template string
	Id       string
	Started  time.Time
	Ends     time.Time
	Label    string
	Duration time.Duration
}

type Data struct {
	Active    *Timer
	Completed []*Timer
	Labels    []string
	Paused    bool
}

type Timers struct {
	Data
	Templates          []*TimerTemplate
	AutoStartBreaks    bool
	AutoStartWork      bool
	DesiredPomsPerDay  int
	PomBeforeLongBreak int
}

func NewTemplate(name, templateLabel string, timerType int, duration time.Duration, sb, lb *string) *TimerTemplate {
	return &TimerTemplate{name, timerType, duration, templateLabel, sb, lb}
}

func (t *Timer) Completed() bool {
	return false // FIXME: Make it so
}
