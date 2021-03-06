package chaika

import "encoding/json"

type Log struct {
	Service string
	Catalog string
	Message string
	LogType string
	Level   string
	Time    string
}

func ParseLog(logInput []byte) (Log, error) {
	var log Log
	err := json.Unmarshal(logInput, &log)
	return log, err
}
