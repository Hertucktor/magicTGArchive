package importer

/*
autogenerated JSON thanks to https://mholt.github.io/json-to-go/
 */
type MTGResponse struct {
	Cards []struct {
		Name          string   `json:"name"`
		Names         []string `json:"names"`
		ManaCost      string   `json:"manaCost"`
		Cmc           float64  `json:"cmc"`
		Colors        []string `json:"colors"`
		ColorIdentity []string `json:"colorIdentity"`
		Type          string   `json:"type"`
		Supertypes    []string `json:"supertypes"`
		Types         []string `json:"types"`
		Subtypes      []string `json:"subtypes"`
		Rarity        string   `json:"rarity"`
		Set           string   `json:"set"`
		SetName       string   `json:"setName"`
		Text          string   `json:"text"`
		Artist        string   `json:"artist"`
		Number        string   `json:"number"`
		Power         string   `json:"power"`
		Toughness     string   `json:"toughness"`
		Layout        string   `json:"layout"`
		Rulings       []struct {
			Date string `json:"date"`
			Text string `json:"text"`
		} `json:"rulings"`
		ForeignNames []interface{} `json:"foreignNames"`
		Printings    []string      `json:"printings"`
		Legalities   []struct {
			Format   string `json:"format"`
			Legality string `json:"legality"`
		} `json:"legalities"`
		ID           string `json:"id"`
		Multiverseid int    `json:"multiverseid,omitempty"`
		ImageURL     string `json:"imageUrl,omitempty"`
		OriginalText string `json:"originalText,omitempty"`
		OriginalType string `json:"originalType,omitempty"`
	} `json:"cards"`
}