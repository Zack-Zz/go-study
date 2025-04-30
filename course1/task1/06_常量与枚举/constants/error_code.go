package main

type ErrorCode int

const (
	// 用户模块错误码 1xx
	ErrUserNotFound ErrorCode = 100 + iota
	ErrInvalidUserInput
	ErrUserBanned

	// 订单模块错误码 2xx
	ErrOrderNotFound ErrorCode = 200 + iota
	ErrPaymentFailed
	ErrTimeout
)

var errorMessages = map[ErrorCode]string{
	ErrUserNotFound:     "User not found",
	ErrInvalidUserInput: "Invalid user input",
	ErrUserBanned:       "User has been banned",
	ErrOrderNotFound:    "Order not found",
	ErrPaymentFailed:    "Payment processing failed",
	ErrTimeout:          "Request timed out",
}

func (e ErrorCode) Message() string {
	if msg, ok := errorMessages[e]; ok {
		return msg
	}
	return "Unknown error"
}
