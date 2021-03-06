package inputs

import (
	"fmt"
	"sync"

	"github.com/mmcquillan/hex/models"
	"github.com/robfig/cron"
)

//Exec struct
type Scheduler struct {
}

//Input function
func (x Scheduler) Read(inputMsgs chan<- models.Message, rules *map[string]models.Rule, config models.Config) {

	// find schedules
	var schedules = make(map[string]int)
	for _, rule := range *rules {
		if rule.Active && rule.Schedule != "" {
			schedules[rule.Schedule] = 0
		}
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	cron := cron.New()
	for schedule, _ := range schedules {
		config.Logger.Debug("Scheduler Input - Adding Schedule '" + schedule + "'")
		cron.AddFunc(schedule, func() {
			message := models.NewMessage()
			message.Attributes["hex.botname"] = config.BotName
			message.Attributes["hex.service"] = "scheduler"
			message.Attributes["hex.schedule"] = schedule
			message.Attributes["hex.input"] = ""
			config.Logger.Debug("Scheduler Input - ID:" + message.Attributes["hex.id"])
			config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
			inputMsgs <- message
		})
	}
	cron.Start()
	defer cron.Stop()
	wg.Wait()

	config.Logger.Warn("Scheduler Ending")
}
