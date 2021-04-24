package timers

type Starter interface {
	Start() error
}

type Stopper interface {
	Stop() error
}

type Pauser interface {
	Pause() error
}

type HasTimerInfo interface {
	Name() string
	TimerType() int
	Label() string
	Completed() bool
	TimeLeft() int
}

type StartStopper interface {
	Starter
	Stopper
	Pauser
}

type FullTimer interface {
	HasTimerInfo
	StartStopper
}
