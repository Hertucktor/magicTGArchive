package importer

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

/*
RequestCardInfo receives a response with type *http.Response from
the official mtg api containing all available card detail for one card
specified with the multiverseID
Returning the response and an error
*/
func RequestCardInfo(URL string, filter string) (MTGResponse, error) {
	var card MTGResponse

	resp, err := http.Get(URL+filter)
	if err != nil {
		log.Print(err)
		return card, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	if resp.StatusCode != 200{
		log.Error().Msg(resp.Status)
		return card, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err(err)
		return card, err
	}

	err = json.Unmarshal(body, &card)
	if err != nil {
		log.Err(err)
		return card, err
	}

	return card, err
}