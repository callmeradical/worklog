package cmd

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func WorkLogPath(p string, f string) string {
	return filepath.Join(os.Getenv("HOME"), p, f)
}
func ArchiveLog(p string, f string) error {
	archive := filepath.Join(os.Getenv("HOME"), p, "worklog_"+time.Now().Format(time.RFC3339)+".csv")
	filename := WorkLogPath(p, f)
	return os.Rename(filename, archive)
}

func CreateLog(f string) error {

	file, err := os.Create(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write([]byte("date, action\n"))

	return err
}

func LogNotExists(f string) bool {
	_, err := os.Stat(f)
	return os.IsNotExist(err)
}
func LogExists(f string) bool {
	return !LogNotExists(f)
}
func UpdateConfig(key string, value string) error {
	filename := filepath.Join(os.Getenv("HOME"), ".worklog.yaml")
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	c := map[string]string{}
	err = yaml.Unmarshal(f, c)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	c[key] = value

	b, _ := yaml.Marshal(c)

	return ioutil.WriteFile(filename, b, 0666)

	return nil
}

func WriteToWorkLog(path string, s string) error {
	if LogNotExists(path) {
		return errors.New("CurrentLog does not exist!")
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(TimeStampActivity(s)); err != nil {
		return err
	}

	return nil

}

func TimeStampActivity(s string) string {
	t := time.Now().Format(time.RFC3339)
	activity := strings.Join([]string{t, s}, ",")
	return activity + "\n"
}

func ReadLog(p string, f string) (string, error) {
	path := WorkLogPath(p, f)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil

}
