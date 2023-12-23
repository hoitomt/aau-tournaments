package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	pb "github.com/aau-tournaments/protos/exposure-events/v1"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var exposure_events = "exposure_events"

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	debug := flag.Bool("debug", false, "set the log level to debug")

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	os.MkdirAll(exposure_events, os.ModePerm)
}

func main() {
	url := "https://basketball.exposureevents.com/youth-basketball-events"
	request := pb.ExposureEventRequest{
		StartDate:  "01/01/2023",
		EndDate:    "12/31/2023",
		Gender:     "2",
		Page:       1,
		InviteType: "-1",
		SportType:  "1",
	}
	// get_page_json(url, "2", "01/01/2023", "12/31/2023", 1)
	get_page_json(url, &request)
}

func get_page_json(requestUrl string, requestPb *pb.ExposureEventRequest) {
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
	// log.Info().Msg(string(resBody))

	path := filepath.Join(exposure_events, fmt.Sprintf("tournaments_%d.json", requestPb.Page))
	err = os.WriteFile(path, resBody, 0644)
	if err != nil {
		log.Fatal().Msgf("Error writing file. %s", err)
	}
}
