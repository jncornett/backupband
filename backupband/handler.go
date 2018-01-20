package backupband

import (
	"log"

	"github.com/jncornett/backupband/backupband/alexa"
	"github.com/mitchellh/mapstructure"
)

// Handler is the main entry point for the BackupBand Alexa skill service.
func Handler(rawEvent map[string]interface{}) {
	event := new(alexa.Event)
	err := mapstructure.Decode(rawEvent, event)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", event)
}
