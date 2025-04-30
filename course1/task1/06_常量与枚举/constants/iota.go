package main

type ConnState int

const (
	StateNew ConnState = iota
	StateActive
	StateIdle
	StateHijacked
	StateClosed
)

type Month int

const (
	January Month = 1 << iota // 左移，2的次方枚举
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
