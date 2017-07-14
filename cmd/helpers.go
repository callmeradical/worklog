package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/callmeradical/worklog/lib/data"
	"gopkg.in/yaml.v2"
)

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

func PresentLog(p string, f string) (string, error) {

	workLog, err := data.ReadLog(p, f)
	if err != nil {
		return "", err
	}

	l := &data.Log{}
	err = json.Unmarshal([]byte(workLog), l)
	if err != nil {
		return "", err
	}

	output := []string{"\t" + l.Header}

	for i, v := range l.Activities {
		output = append(output, fmt.Sprintf("%d %s", i+1, v.ToCSVEntry()))
	}

	return strings.Join(output, "\n"), nil

}
