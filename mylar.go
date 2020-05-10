package mylar

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// New creates a new Mylar client instance.
func New(apiURL, apiKey string) (*Mylar, error) {
	if apiURL == "" {
		return &Mylar{}, errors.New("apiURL is required")
	}
	if apiKey == "" {
		return &Mylar{}, errors.New("apiKey is required")
	}
	baseURL, err := url.Parse(apiURL)
	if err != nil {
		return &Mylar{}, fmt.Errorf("failed to parse baseURL: %v", err)
	}
	return &Mylar{
		baseURL:    baseURL,
		apiKey:     apiKey,
		HTTPClient: http.Client{},
	}, nil
}

type Mylar struct {
	baseURL    *url.URL
	apiKey     string
	HTTPClient http.Client
	//Timeout in seconds -- default 5
	Timeout int
}

var commandValues = []string{
	"getIndex",
	"getComic",
	"getWanted",
	"getHistory",
}

type command int

func (c command) String() string {
	return commandValues[c]
}

const (
	GetIndexCommand command = iota
	GetComicCommand
	GetWantedCommand
	GetHistoryCommand
)

func (m Mylar) GetIndex() ([]Comic, error) {
	httpResponse, err := m.get(GetIndexCommand, url.Values{})
	if err != nil {
		return nil, err
	}
	response, err := handleStructuredResponse(httpResponse)
	if err != nil {
		return nil, err
	}
	var results []Comic
	err = json.Unmarshal(response, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (m Mylar) GetComic(id string) (*ComicDetail, error) {
	params := url.Values{}
	params.Set("id", id)
	httpResponse, err := m.get(GetComicCommand, params)
	if err != nil {
		return nil, err
	}
	response, err := handleStructuredResponse(httpResponse)
	if err != nil {
		return nil, err
	}
	var result ComicDetail
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m Mylar) GetWanted() ([]WantedIssue, error) {
	params := url.Values{}
	httpResponse, err := m.get(GetWantedCommand, params)
	if err != nil {
		return nil, err
	}
	var result []WantedIssue
	err = json.NewDecoder(httpResponse.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m Mylar) GetHistory() ([]History, error) {
	httpResponse, err := m.get(GetHistoryCommand, url.Values{})
	if err != nil {
		return nil, err
	}
	response, err := handleStructuredResponse(httpResponse)
	if err != nil {
		return nil, err
	}
	var results []History
	err = json.Unmarshal(response, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (m *Mylar) get(cmd command, params url.Values) (*http.Response, error) {
	params.Set("cmd", cmd.String())
	params.Set("apikey", m.apiKey)
	requestURL := m.baseURL
	requestURL.Path = "/api"
	requestURL.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return m.HTTPClient.Do(req)
}

func handleStructuredResponse(resp *http.Response) (json.RawMessage, error) {
	defer resp.Body.Close()
	var response Response
	err := json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, fmt.Errorf("error %d: %s", response.Error.Code, response.Error.Message)
	}
	return response.Data, nil
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success bool            `json:"success"`
	Error   Error           `json:"error"`
	Data    json.RawMessage `json:"data"`
}

type Comic struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"imageURL"`
	Status      string `json:"status"`
	Publisher   string `json:"publisher"`
	Year        string `json:"year"`
	LatestIssue string `json:"latestIssue"`
	Total       int    `json:"totalIssues"`
	DetailsURL  string `json:"detailsURL"`
}

type ComicDetail struct {
	Comic   []Comic `json:"comic"`
	Annuals []Issue `json:"annuals"`
	Issues  []Issue `json:"issues"`
}

type Issue struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"imageURL"`
	Number      string `json:"number"`
	ReleaseDate string `json:"releaseDate"`
	IssueDate   string `json:"issueDate"`
	Status      string `json:"status"`
	ComicName   string `json:"comicName"`
}

type WantedIssue struct {
	Status          string
	ComicName       string
	IssueID         string
	DigitalDate     string
	IssueDate       string
	ImageURL        string
	ReleaseDate     string
	IssueNumberText string `json:"Issue_Number"`
	IssueNumber     int    `json:"Int_IssueNumber"`
	IssueName       string
	ComicID         string
	DateAdded       string
}

type History struct {
	Status      string
	ComicName   string
	IssueID     string
	CheckSum    string `json:"crc"`
	IssueNumber string `json:"Issue_Number"`
	ComicID     string
	Provider    string
	DateAdded   string
}
