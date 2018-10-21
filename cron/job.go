package cron

import (
	"log"

	"github.com/robfig/cron"
	"nativepub.net/idrequest/csv"
)

func ReloadMapEvery30Minutes(synchronizedMap *csv.SynchronizedMap) {
	records, err := csv.GetRecordsFromCsvFile("ids.csv")
	if err != nil {
		log.Fatal(err)
	}
	synchronizedMap.UpdateMap(records)
	log.Printf("loaded synchronized map")

	cron := cron.New()
	cron.AddFunc("@every 30m", func() {
		records, err := csv.GetRecordsFromCsvFile("ids.csv")
		synchronizedMap.UpdateMap(records)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("reloaded synchronized map")

	})
	cron.Start()
}
