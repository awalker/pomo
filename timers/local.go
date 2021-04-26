package timers

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"pomo/cmd"

	"github.com/spf13/viper"
)

type LocalTimers struct {
	Timers
	filePath string
}

func Load() (*LocalTimers, error) {
	dataFolder := viper.GetString(cmd.DATA)
	jsonFileName := path.Join(dataFolder, "timers.json")
	timers := LocalTimers{}
	timers.filePath = jsonFileName
	return &timers, timers.Load()
}

func (t *LocalTimers) Load() error {
	if jsonFile, err := os.Open(t.filePath); err == nil {
		defer jsonFile.Close()
		decoder := json.NewDecoder(jsonFile)
		err = decoder.Decode(&t.Timers)
		if err != nil {
			return fmt.Errorf("timers:Failed to parse config file (%w)", err)
		}
		return nil
	} else if os.IsNotExist(err) {
		// Data file not found. Create a blank/default data file.
		t.AutoStartBreaks = true
		t.DesiredPomsPerDay = 8
		t.PomBeforeLongBreak = 4
		ssb, slb := "Short Break", "Long Break"
		w := NewTemplate("Pom", "", 1, 60*25, &ssb, &slb)
		sb := NewTemplate(ssb, "", 2, 60*5, nil, nil)
		lb := NewTemplate(slb, "", 3, 60*15, nil, nil)
		t.Templates = []*TimerTemplate{w, sb, lb}
		// Should probably save the config now.
		return t.Save()
	} else {
		return fmt.Errorf("timers:Could not load config file (%w)", err)
	}
}

func (t *LocalTimers) Save() error {
	json.NewEncoder(os.Stdout).Encode(t.Timers)
	/*if wFile, err := os.Create(jsonFileName); err == nil {
		defer wFile.Close()
		enc := json.NewEncoder(wFile)
		enc.Encode(timers)
	} else {
		fmt.Println("Could not right configuration file. ", jsonFileName, err)
	}*/
	return nil
}
