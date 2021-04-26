package timers

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

const DATA = "data.directory"
const (
	AUTOSTART_BREAKS               = "autostart.breaks"
	AUTOSTART_WORK                 = "autostart.work"
	SHORT_BREAKS_BEFORE_LONG_BREAK = "default.short_breaks_before_long_break"
	GOAL                           = "default.goal"
	DEFAULT_BREAKS_LONG            = "default.breaks.long"
	DEFAULT_BREAKS_SHORT           = "default.breaks.short"
	DEFAULT_WORK                   = "default.work"
)

type LocalTimers struct {
	Timers
	filePath string
}

func Load() (*LocalTimers, error) {
	dataFolder := viper.GetString(DATA)
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
			return fmt.Errorf("timers: Failed to parse config file (%w)", err)
		}
		return nil
	} else if os.IsNotExist(err) {
		// Data file not found. Create a blank/default data file.
		t.AutoStartBreaks = viper.GetBool(AUTOSTART_BREAKS)
		t.DesiredPomsPerDay = viper.GetInt(GOAL)
		t.PomBeforeLongBreak = viper.GetInt(SHORT_BREAKS_BEFORE_LONG_BREAK)
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

func timerTemplateConfig(name string, duration int) map[string]interface{} {
	return map[string]interface{}{"name": name, "duration": duration}
}

func init() {
	viper.SetDefault(AUTOSTART_BREAKS, true)
	viper.SetDefault(AUTOSTART_WORK, false)
	viper.SetDefault(SHORT_BREAKS_BEFORE_LONG_BREAK, 4)
	viper.SetDefault(GOAL, 8)
	viper.SetDefault(DEFAULT_BREAKS_LONG, timerTemplateConfig("Long Break", 15*60))
	viper.SetDefault(DEFAULT_BREAKS_SHORT, timerTemplateConfig("Short Break", 5*60))
	viper.SetDefault(DEFAULT_WORK, timerTemplateConfig("Work", 25*60))
}
