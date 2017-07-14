package data

import (
	"fmt"
	"time"
)

type Activity struct {
	Type     string `json:"type"`
	Date     string `json:"date"`
	Activity string `json:"activity"`
}

func (n Activity) ToCSVEntry() string {
	return fmt.Sprintf("\t%s, %s, %s", n.Type, n.Date, n.Activity)
}

func NewActivity() *Activity {
	return &Activity{Date: time.Now().Format(time.RFC3339)}
}
