package exposureevents

import (
	"bytes"
	"io"
	"net/http"
	"time"

	pb "github.com/aau-tournaments/protos/exposure-events/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
)

var tournamentListUrl = "https://basketball.exposureevents.com/youth-basketball-events"

func GetTournamentList() {
	request := pb.ExposureEventTournamentListRequest{
		StartDate:  "01/01/2023",
		EndDate:    "12/31/2023",
		Gender:     "2",
		Page:       1,
		InviteType: "-1",
		SportType:  "1",
	}
	getPage(tournamentListUrl, &request)
}

func getPage(requestUrl string, requestPb *pb.ExposureEventTournamentListRequest) (*pb.ExposureEventTournamentListResponse, error) {
	log.Info().Msgf("Proto Message: %s", requestPb.String())
	jsonBody, _ := protojson.Marshal(requestPb)
	log.Info().Msgf("Send request to %s with %s", requestUrl, string(jsonBody))

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	if err != nil {
		log.Fatal().Msgf("Error constructing the request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal().Msgf("client: error making http request: %s", err)
	}

	if res.StatusCode != 200 {
		log.Fatal().Msgf("HTTP Error: %d", res.StatusCode)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal().Msgf("Error parsing the response body. %s", err)
	}
	// log.Debug().Msg(string(resBody))
	tournamentListResponse := pb.ExposureEventTournamentListResponse{}
	err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(resBody, &tournamentListResponse)
	// err = protojson.Unmarshal(resBody, &tournamentListResponse)
	if err != nil {
		log.Error().Msgf("Unmarshal tournament list response error: %s", err)
		return nil, err
	}

	// path := filepath.Join("data", "exposure_events", fmt.Sprintf("tournaments_%d.json", requestPb.Page))
	// err = os.WriteFile(path, resBody, 0644)
	// if err != nil {
	// 	log.Fatal().Msgf("Error writing file. %s", err)
	// }
	return &tournamentListResponse, nil
}
