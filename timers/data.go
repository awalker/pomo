package timers

type TimerTemplate struct {
	Name          string
	TimerType     int
	Duration      int
	TemplateLabel string
	ShortBreak    *string
	LongBreak     *string
}

type Timer struct {
	Template string
	Started  int
	Ends     int
	Label    string
}

type Timers struct {
	Active             *Timer
	Completed          []*Timer
	Templates          []*TimerTemplate
	AutoStartBreaks    bool
	AutoStartWork      bool
	Paused             bool
	DesiredPomsPerDay  int
	PomBeforeLongBreak int
}

func NewTemplate(name, templateLabel string, timerType, duration int, sb, lb *string) *TimerTemplate {
	return &TimerTemplate{name, timerType, duration, templateLabel, sb, lb}
}

func (t *Timer) Completed() bool {
	return false // FIXME: Make it so
}
