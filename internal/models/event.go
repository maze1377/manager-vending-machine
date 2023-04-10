package models

type Event uint8

const (
	Payment Event = iota
	Dispensed
)
