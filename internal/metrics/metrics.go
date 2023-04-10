package metrics

import "time"

// Metric is an interface which defines metric handlers
// for example Prometheus is a metric handler which implements this interfaces.
type Metric interface {
	NewTimer(c TimerConfig, labels ...string) Timer
	NewCounter(c CounterConfig, labels ...string) Counter
	NewCommunicator(c CommunicatorConfig) Communicator
}

type Timer interface {
	Done(started time.Time, labels ...string)
}

type Counter interface {
	Inc(labels ...string)
}

type Communicator interface {
	Done(started time.Time, method, status string)
}

type TimerConfig struct {
	Name    string
	Help    string
	Buckets []float64
}

type CounterConfig struct {
	Name string
	Help string
}

type CommunicatorConfig struct {
	Name    string
	Help    string
	Buckets []float64
}
