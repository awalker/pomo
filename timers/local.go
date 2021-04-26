package timers

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

const (
	DATA                           = "data.directory"
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
	filePath         string
	templateFilePath string
}

func Load() (*LocalTimers, error) {
	dataFolder := viper.GetString(DATA)
	jsonFileName := path.Join(dataFolder, "timers.json")
	templateFileName := path.Join(dataFolder, "templates.json")
	timers := LocalTimers{}
	timers.filePath = jsonFileName
	timers.templateFilePath = templateFileName
	return &timers, timers.Load()
}

func (t *LocalTimers) Load() error {
	t.AutoStartBreaks = viper.GetBool(AUTOSTART_BREAKS)
	t.DesiredPomsPerDay = viper.GetInt(GOAL)
	t.PomBeforeLongBreak = viper.GetInt(SHORT_BREAKS_BEFORE_LONG_BREAK)

	if jsonFile, err := os.Open(t.templateFilePath); err == nil {
		defer jsonFile.Close()
		decoder := json.NewDecoder(jsonFile)
		err = decoder.Decode(&t.Timers.Templates)
		if err != nil {
			return fmt.Errorf("timers: Failed to parse templates file (%w)", err)
		}
	} else if os.IsNotExist(err) {
		// Data file not found. Create a blank/default data file.
		const (
			NAME = ".name"
			DUR  = ".duration"
		)
		ssb, slb := viper.GetString(DEFAULT_BREAKS_SHORT+NAME), viper.GetString(DEFAULT_BREAKS_LONG+NAME)
		dssb, dslb := viper.GetDuration(DEFAULT_BREAKS_SHORT+DUR), viper.GetDuration(DEFAULT_BREAKS_LONG+DUR)
		def_name, def_dur := viper.GetString(DEFAULT_WORK+NAME), viper.GetDuration(DEFAULT_WORK+DUR)
		w := NewTemplate(def_name, "", 1, def_dur, &ssb, &slb)
		sb := NewTemplate(ssb, "", 2, dssb, nil, nil)
		lb := NewTemplate(slb, "", 3, dslb, nil, nil)
		t.Templates = []*TimerTemplate{w, sb, lb}
		// Should probably save the config now.
		if err := t.SaveTemplates(); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("timers: Could not load templates file (%w)", err)
	}

	if jsonFile, err := os.Open(t.filePath); err == nil {
		defer jsonFile.Close()
		decoder := json.NewDecoder(jsonFile)
		err = decoder.Decode(&t.Data)
		if err != nil {
			return fmt.Errorf("timers: Failed to parse config file (%w)", err)
		}
		t.Normalize()
		return nil
	} else if os.IsNotExist(err) {
		// Data file not found. Create a blank/default data file.
		// Should probably save the config now.
		return t.SaveData()
	} else {
		return fmt.Errorf("timers: Could not load templates file (%w)", err)
	}
}

func (t *LocalTimers) Normalize() error {
	// Check if active timer has expired
	if t.Active != nil {
		if time.Now().After(t.Active.Ends) {
			fmt.Println("afet")
			t.Completed = append(t.Completed, t.Active)
			t.Active = nil
			// TODO: check auto starts
			return t.SaveData()
		} else {
			fmt.Println("not after")
		}
	}
	return nil
}

func (t *LocalTimers) SaveTemplates() error {
	// json.NewEncoder(os.Stdout).Encode(t.Timers.Templates)
	if wFile, err := os.Create(t.templateFilePath); err == nil {
		defer wFile.Close()
		enc := json.NewEncoder(wFile)
		enc.Encode(t.Templates)
	} else {
		fmt.Println("Could not write templates file. ", t.templateFilePath, err)
	}
	return nil
}

func (t *LocalTimers) SaveData() error {
	if wFile, err := os.Create(t.filePath); err == nil {
		defer wFile.Close()
		enc := json.NewEncoder(wFile)
		enc.Encode(t.Data)
	} else {
		fmt.Println("Could not write data file. ", t.filePath, err)
	}
	return nil
}

func (d *LocalTimers) FindTemplate(name string) *TimerTemplate {
	for _, t := range d.Templates {
		if t.Name == name {
			return t
		}
	}
	return nil
}

func (d *LocalTimers) Start(name string) error {
	if d.Active == nil {
		template := d.FindTemplate(name)
		if template == nil {
			return fmt.Errorf("localTimer.Start: Template not found")
		}
		// We have a template, create the Active Timer
		timer := new(Timer)
		timer.Id = uuid.New().String()
		timer.Template = name
		timer.Duration = template.Duration
		timer.Started = time.Now()
		timer.Ends = time.Now().Add(template.Duration)
		d.Active = timer
		return nil
	} else {
		return fmt.Errorf("localTimer.Start: A timer is already active")
	}
}

func timerTemplateConfig(name string, duration string) map[string]interface{} {
	return map[string]interface{}{"name": name, "duration": duration}
}

func init() {
	viper.SetDefault(AUTOSTART_BREAKS, true)
	viper.SetDefault(AUTOSTART_WORK, false)
	viper.SetDefault(SHORT_BREAKS_BEFORE_LONG_BREAK, 4)
	viper.SetDefault(GOAL, 8)
	viper.SetDefault(DEFAULT_BREAKS_LONG, timerTemplateConfig("Long Break", "15m"))
	viper.SetDefault(DEFAULT_BREAKS_SHORT, timerTemplateConfig("Short Break", "5m"))
	viper.SetDefault(DEFAULT_WORK, timerTemplateConfig("Work", "25m"))
}
