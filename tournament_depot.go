package main

// func get_url(path string) {
// 	url := "https://www.tournamentdepot.com"

// 	requestUrl := fmt.Sprintf("%s/%s", url, path)

// 	log.Info().Msgf("Send request to %s", requestUrl)
// 	res, err := http.Get(requestUrl)
// 	if err != nil {
// 		log.Fatal().Msgf("ERROR: %s", err)
// 	}

// 	if res.StatusCode != 200 {
// 		log.Fatal().Msgf("HTTP Error: %d", res.StatusCode)
// 	}

// 	resBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal().Msgf("Error parsing the response body. %s", err)
// 	}
// 	// log.Info().Msg(string(resBody))

// 	err = os.WriteFile("current_tournaments.json", resBody, 0644)
// 	if err != nil {
// 		log.Fatal().Msgf("Error writing file. %s", err)
// 	}
// }

// func search(year int) {
// 	url := "https://www.tournamentdepot.com"
// 	path := "servlet/rest/tournaments/search"
// 	requestUrl := fmt.Sprintf("%s/%s", url, path)

// 	// {
// 	// 	"title_desc":"",
// 	// 	"sport":"",
// 	// 	"address":"",
// 	// 	"city":"",
// 	// 	"state":"",
// 	// 	"zip":"",
// 	// 	"country":"",
// 	// 	"gender":"",
// 	// 	"age":"",
// 	// 	"month":"",
// 	// 	"year":2023,
// 	// 	"maximumcost":""
// 	// }

// 	jsonBody := []byte(`{"year": year}`)
// 	log.Info().Msgf("Send request to %s with %s", requestUrl, string(jsonBody))

// 	bodyReader := bytes.NewReader(jsonBody)

// 	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
// 	if err != nil {
// 		log.Fatal().Msgf("Error constructing the request: %s", err)
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	client := http.Client{
// 		Timeout: 30 * time.Second,
// 	}

// 	res, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal().Msgf("client: error making http request: %s", err)
// 	}

// 	if res.StatusCode != 200 {
// 		log.Fatal().Msgf("HTTP Error: %d", res.StatusCode)
// 	}

// 	resBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal().Msgf("Error parsing the response body. %s", err)
// 	}
// 	// log.Info().Msg(string(resBody))

// 	err = os.WriteFile("current_tournaments.json", resBody, 0644)
// 	if err != nil {
// 		log.Fatal().Msgf("Error writing file. %s", err)
// 	}
// }
