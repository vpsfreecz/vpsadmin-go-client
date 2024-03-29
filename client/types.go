package client

import (
	"strconv"
)

type ProgressCallbackReturn int

const (
	ContinueWatching = iota
	StopWatching     = iota
)

type OperationProgressCallback func(*ActionActionStatePollOutput) ProgressCallbackReturn

type BlockingOperationWatcher interface {
	IsBlocking() bool
	OperationStatus() (*ActionActionStateShowResponse, error)
	WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error)
	WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error)
}

func convertInt64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func convertFloat64ToString(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func convertBoolToString(v bool) string {
	if v {
		return "1"
	} else {
		return "0"
	}
}

func convertResourceToString(v int64) string {
	return convertInt64ToString(v)
}
