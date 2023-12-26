package exposureevents

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	pb "github.com/aau-tournaments/protos/exposure-events/v1"
)

func TestExposureEventsGetTournamentList(t *testing.T) {
	testResponse, err := os.ReadFile(filepath.Join("..", "data", "test", "exposure_events", "tournament_list.json"))
	if err != nil {
		t.Errorf("Read test data failed %s", err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(testResponse)
	}))
	defer server.Close()

	request := pb.ExposureEventTournamentListRequest{
		StartDate:  "01/01/2023",
		EndDate:    "12/31/2023",
		Gender:     "2",
		Page:       1,
		InviteType: "-1",
		SportType:  "2",
	}
	tournamentList, err := getPage(server.URL, &request)
	t.Logf("Tournament list: %v", tournamentList)
	if err != nil {
		t.Errorf("getPage exception: %s", err)
	}

	expectedLength := 50
	if len(tournamentList.Results) != expectedLength {
		t.Errorf("Expected tournament list length to be %d but was %d", expectedLength, len(tournamentList.Results))
	}
}
