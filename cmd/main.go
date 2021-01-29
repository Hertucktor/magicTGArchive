package main

import (
	"github.com/rs/zerolog/log"
	"magicTGArchive/internal/pkg/mongodb"
)

func main() {
/*
	language, cardName, err := cli.ReadFromCLI()
	if err != nil {
		log.Fatal().Err(err)
	}

	URL := importer.URLGenerator(language, cardName)

	fmt.Println(URL)

	cardInfo, err := importer.RequestCardInfo(URL)
	if err != nil {
		log.Fatal().Err(err)
	}*/

	client, ctx, err := mongodb.CreateClient()
	if err != nil {
		log.Fatal().Err(err)
	}

	/*for _,cards := range cardInfo.Cards{
		if err = mongodb.InsertCardInfo(cards, client, ctx); err != nil {
			log.Fatal().Err(err)
		}
	}*/

	/*
	if err = mongodb.GetAllCardInfo(client, ctx); err != nil{
		log.Fatal().Err(err)
	}*/

	cardName := "Quicksand"
	if err = mongodb.GetFilteredSingleCardInfo(cardName, client, ctx); err != nil{
		log.Fatal().Err(err)
	}
}