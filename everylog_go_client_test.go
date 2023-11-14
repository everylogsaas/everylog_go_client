package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestCreateLogEntry(*testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder("POST", "https://everylog.io/api/v1/log-entries",
		func(req *http.Request) (*http.Response, error) {

			logEntry := PostBody{}
			if err := json.NewDecoder(req.Body).Decode(&logEntry); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			if len(logEntry.ProjectID) == 0 {
				return httpmock.NewStringResponse(422, "ProjectID is required"), nil
			}
			if len(logEntry.Body) == 0 {
				return httpmock.NewStringResponse(422, "Body is required"), nil
			}
			if len(logEntry.Summary) == 0 {
				return httpmock.NewStringResponse(422, "Summary is required"), nil
			}
			if len(logEntry.Title) == 0 {
				return httpmock.NewStringResponse(422, "Title is required"), nil
			}
			if len(logEntry.Icon) > 1 {
				return httpmock.NewStringResponse(422, "icon max length is 1"), nil
			}

			resp, err := httpmock.NewJsonResponse(200, logEntry)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			fmt.Printf("\n Resp: %v \n", resp)
			return resp, nil
		})

	postBody := PostBody{
		ProjectID:        "EverylogClient",
		Title:            "Test go client",
		Summary:          "Summary go client",
		Body:             "Body go client",
		Groups:           []string{"group-1"},
		Push:             false,
		Tags:             []string{"tag1, tag2"},
		Properties:       []map[string]interface{}{{"name": "mario", "surname": "rossi"}},
		Icon:             "😀",
		ExternalChannels: []string{"notify-email", "notify-slack"},
	}

	resp, err := CreateLogEntry("c7ac9f54-5074-4513-8086-4c6a5adea34c", "prova", "https://everylog.io/api/v1/log-entries", postBody)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("\n Resp: %v \n", resp)
}
