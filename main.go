package main

import (
	"context"
	
	"log"

	"github.com/areaverua/Jane/clients/telegram"
	tgClient "github.com/areaverua/Jane/clients/telegram"
	
	event_consumer "github.com/areaverua/Jane/consumer/event-consumer"
	"github.com/areaverua/Jane/storage/sqlite"
)
const (
	tgBotHost   = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize   = 100
)

func main() {
	/*cfg := config.Musload()
	s := files.New(storagePath)

	 s, err := sqlite.New(sqliteStoragePath)
	 if err != nil{
		log.Fatal("can't connect to storage", err)
	 }


	 //err := s.Init(context.TODO())
	 if err != nil{
		log.Fatal("can't init storage: ", err)
	 }
	
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
	*/
}