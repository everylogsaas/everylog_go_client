package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PostBody struct {
	ProjectID        string                   `json:"projectId"`
	Title            string                   `json:"title"`
	Summary          string                   `json:"summary"`
	Body             string                   `json:"body"`
	Link             string                   `json:"link"`
	Push             bool                     `json:"push"`
	Tags             []string                 `json:"tags"`
	Groups           []string                 `json:"groups"`
	ExternalChannels []string                 `json:"externalChannels"`
	Properties       []map[string]interface{} `json:"properties"`
	Icon             string                   `json:"icon"`
}

func CreateLogEntry(key string, project string, url string, postBody PostBody) (responseBody interface{}, err error) {

	jsonFormat, err := json.Marshal(postBody)
	if err != nil {
		fmt.Printf("\n MarshalError: %v \n", err)
	}

	body := []byte(fmt.Sprintf(`%v`, string(jsonFormat)))

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %v", key))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		fmt.Printf("\n Error: %s \n", err)
	}

	defer res.Body.Close()

	respBody, _ := io.ReadAll(res.Body)
	responseBody = string(respBody)

	if res.StatusCode != http.StatusCreated {
		fmt.Printf("\n Status: %s \n", res.Status)
	}

	return responseBody, err
}
