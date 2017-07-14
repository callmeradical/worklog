package data

import (
	"encoding/json"
	"strconv"
	"strings"
)

const (
	DefaultHeader = "type, date, activity"
)

type Log struct {
	Header     string      `json:"header,omitempty"`
	Activities []*Activity `json:"activities"`
}

func NewLog(header bool) *Log {
	return &Log{Header: DefaultHeader}
}

func NewLogFromJson(input []byte) (*Log, error) {

	newLog := &Log{}
	err := json.Unmarshal(input, newLog)
	if err != nil {
		return nil, err
	}

	return newLog, nil
}

func (l Log) ToCSV() string {
	csv := []string{"\t" + l.Header + "\n"}

	for i, a := range l.Activities {
		entry := strconv.Itoa(i + 1)
		csv = append(csv, entry+a.ToCSVEntry())
	}

	return strings.Join(csv, "")
}

func (l Log) ToJSON() (string, error) {
	b, err := json.Marshal(l)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
