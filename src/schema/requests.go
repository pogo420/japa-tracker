// Package contains all schema or response/request for endpoints
// All internal data structures required for communication
package schema

type DailyJapaCount struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}
