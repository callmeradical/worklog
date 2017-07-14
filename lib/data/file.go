package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ArchiveLog(p string, f string) error {
	archive := filepath.Join(os.Getenv("HOME"), p, time.Now().Format(time.RFC3339)+"_"+f)
	filename := WorkLogPath(p, f)
	return os.Rename(filename, archive)
}

func CreateLogFile(f string) error {

	file, err := os.Create(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	l := NewLog(false)
	output, err := l.ToJSON()
	if err != nil {
		return err
	}
	_, err = file.WriteString(output)
	if err != nil {
		return err
	}

	return nil
}

func LogNotExists(f string) bool {
	_, err := os.Stat(f)
	return os.IsNotExist(err)
}
func LogExists(f string) bool {
	return !LogNotExists(f)
}

func ReadLog(p string, f string) (string, error) {
	b, err := ioutil.ReadFile(WorkLogPath(p, f))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ReadJsonLog(path string) ([]byte, error) {
	if LogNotExists(path) {
		return nil, errors.New("CurrentLog does not exist!")
	}

	return ioutil.ReadFile(path)
}

func UpdateActivity(path string, file string, t string, i int) error {
	t = strings.ToUpper(t)
	path = WorkLogPath(path, file)
	currentLog, err := ReadJsonLog(path)
	if err != nil {
		return err
	}

	newLog, err := NewLogFromJson([]byte(currentLog))
	if err != nil {
		return err
	}

	if i < 1 || i > len(newLog.Activities) {
		return errors.New("Index out of range, unable to update activity")
	}

	newLog.Activities[i-1].Type = t

	b, _ := json.Marshal(newLog)

	return WriteJsonLog(path, b)
}

func WorkLogPath(p string, f string) string {
	return filepath.Join(os.Getenv("HOME"), p, f)
}

func WriteJsonLog(path string, currentLog []byte) error {
	return ioutil.WriteFile(path, currentLog, 0600)
}

func WriteToWorkLog(path string, t string, s string) error {
	currentLog, err := ReadJsonLog(path)
	if err != nil {
		return err
	}

	newLog, err := NewLogFromJson([]byte(currentLog))
	if err != nil {
		return err
	}

	a := &Activity{
		Type:     strings.ToUpper(t),
		Date:     time.Now().Format(time.RFC3339),
		Activity: s,
	}

	newLog.Activities = append(newLog.Activities, a)

	b, _ := json.Marshal(newLog)

	return WriteJsonLog(path, b)
}
