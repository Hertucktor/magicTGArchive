package importer

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

/*
RequestAllCards receives a response with type *http.Response from
the mtg api containing 100 cards.
Returning the response and an error
*/
func RequestAllCards(page string) (APIResponseForMultipleCards, error) {
	var response APIResponseForMultipleCards
	var resp *http.Response
	var err error
	var body []byte
	//GET request to URL with page param
	if resp, err = http.Get("https://api.magicthegathering.io/v1/cards?page="+page); err != nil{
		log.Error().Timestamp().Err(err).Msg("Error: problem with http GET request\n")
		return response, err
	}

	log.Info().Timestamp().Msgf("HTTP GET REQUEST TO https://api.magicthegathering.io/v1/cards?page=\n",page)

	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't close response body\n")
		}
	}()
	//checks if there is an http status code other than 200 in the response
	if resp.StatusCode != 200{
		err = errors.New("Http statuscode != 200")
		log.Error().Timestamp().Err(err).Msgf("Error: Http status code:\n", resp.StatusCode)
		return response, err
	}
	//reads response body into []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't read from response body\n")
		return response, err
	}
	//parses response body []byte values into response
	if err = json.Unmarshal(body, &response); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: couldn't unmarshal body into MTGDevAPIResponse struct\n")
		return response, err
	}

	return response, err
}