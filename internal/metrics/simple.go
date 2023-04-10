package metrics

import (
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type simple struct {
	listener net.Listener
	handler  http.Handler
	addr     string
}

type simpleCommunicator struct {
	timer   Timer
	counter Counter
}

func (s *simpleCommunicator) Done(started time.Time, method, status string) {
	s.timer.Done(started, method, status)
	s.counter.Inc(method, status)
}

func (p *simple) NewCommunicator(c CommunicatorConfig) Communicator {
	timer := p.NewTimer(TimerConfig{
		Name:    c.Name + "_timer",
		Help:    c.Help,
		Buckets: c.Buckets,
	}, "method", "status")
	counter := p.NewCounter(CounterConfig{
		Name: c.Name + "_counter",
		Help: c.Help,
	}, "method", "status")
	return &simpleCommunicator{
		timer:   timer,
		counter: counter,
	}
}

type simpleTimer struct {
	watcher *prometheus.HistogramVec
}

func (s simpleTimer) Done(started time.Time, labels ...string) {
	end := time.Now()
	duration := end.Sub(started)
	s.watcher.WithLabelValues(labels...).Observe(duration.Seconds())
}

func (p *simple) NewTimer(c TimerConfig, labels ...string) Timer {
	if len(c.Buckets) == 0 {
		c.Buckets = prometheus.DefBuckets
	}
	collector := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    c.Name,
		Help:    c.Help,
		Buckets: c.Buckets,
	}, labels)
	prometheus.MustRegister(collector)
	return &simpleTimer{collector}
}

type simpleCounter struct {
	watcher *prometheus.CounterVec
}

func (s *simpleCounter) Inc(labels ...string) {
	s.watcher.WithLabelValues(labels...).Inc()
}

func (p *simple) NewCounter(c CounterConfig, labels ...string) Counter {
	collector := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: c.Name,
		Help: c.Help,
	}, labels)
	prometheus.MustRegister(collector)
	return &simpleCounter{collector}
}

func NewSimpleMetric(addr string) (Metric, error) {
	p := &simple{addr: addr}
	newListener, err := net.Listen("tcp", p.addr)
	if err != nil {
		return nil, err
	}
	p.listener = newListener
	p.handler = promhttp.Handler()
	return p, nil
}
