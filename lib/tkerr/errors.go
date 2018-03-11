package tkerr

import "fmt"

const (
	InvalidAccount = iota
	InvalidUsername
	UntrackedAccount
	OSAPIError
	RecentUpdate
	InvalidPeriod
	NoPointsInPeriod
)

type TKError struct {
	Code    int
	Message string
}

func (e *TKError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

var errorMessages = map[int]string{
	InvalidAccount:   "Account does not exist",
	InvalidUsername:  "Invalid username",
	UntrackedAccount: "Account has not been tracked",
	OSAPIError:       "Problem with OSRS API",
	RecentUpdate:     "Account was recently updated",
	InvalidPeriod:    "Invalid period",
	NoPointsInPeriod: "No datapoints for account in specified period",
}

func Create(code int) *TKError {
	return &TKError{Code: code, Message: errorMessages[code]}
}
