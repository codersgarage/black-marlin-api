package worker

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/codersgarage/black-marlin-api/config"
	"github.com/codersgarage/black-marlin-api/log"
	"github.com/codersgarage/black-marlin-api/mq"
)

var worker *machinery.Worker

func RunWorker() {
	go func() {
		worker = mq.MQServer().NewWorker(config.MQ().Worker.Name, config.MQ().Worker.Count)
		err := worker.Launch()
		if err != nil {
			log.Log().Errorln("Failed to launch worker :", err)
		}
	}()
}
