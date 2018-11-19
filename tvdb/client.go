package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type LoginReq struct {
	Apikey   string `json:"apikey,omitempty"`
	Userkey  string `json:"userkey,omitempty"`
	Username string `json:"username,omitempty"`
}
type LoginRes struct {
	Token string `json:"token"`
}

type SearchResp struct {
	Data []Serie `json:data`
}
type Serie struct {
	Banner     string `json:"banner,omitempty"`
	ID         int    `json:"id,omitempty"`
	SeriesName string `json:"seriesName,omitempty"`
	Slug       string `json:"slug,omitempty"`
}
type GetEpisodeResp struct {
	Data []Episode `json:data`
}
type Episode struct {
	ID                 int    `json:"id,omitempty"`
	EpisodeName        string `json:"episodeName,omitempty"`
	AiredSeason        int    `json:"airedSeason,omitempty"`
	AiredEpisodeNumber int    `json:"AiredEpisodeNumber,omitempty"`
}

type Client struct {
	BaseURL       *url.URL
	Authorization string
	httpClient    *http.Client
}

func defaultHeaders(headers *http.Header) {
	headers.Add("Content-Type", "application/json")
	headers.Add("Accept", "application/json")
	headers.Add("Accept-Language", "en-US")
	headers.Add("User-Agent", "DDM")
}

func NewClient() *Client {
	url, _ := url.Parse("https://api.thetvdb.com")
	return &Client{
		BaseURL:    url,
		httpClient: &http.Client{},
	}
}

func (c *Client) makeRequest(response *interface{}) {

}

func (c *Client) Login(apikey string) error {
	path := &url.URL{Path: "/login"}
	finalURL := c.BaseURL.ResolveReference(path)
	loginReq, err := json.Marshal(LoginReq{Apikey: apikey})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", finalURL.String(), bytes.NewBuffer(loginReq))
	if err != nil {
		return err
	}
	defaultHeaders(&req.Header)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if 200 != resp.StatusCode {
		return fmt.Errorf("%s", body)
	}
	var loginRes LoginRes
	err = json.Unmarshal(body, &loginRes)
	if err != nil {
		return err
	}
	c.Authorization = "Bearer " + loginRes.Token
	return nil
}

func (c *Client) SearchSeries(name string) (SearchResp, error) {
	path := &url.URL{Path: "/search/series"}
	finalURL := c.BaseURL.ResolveReference(path)
	q := finalURL.Query()
	q.Set("name", name)
	finalURL.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", finalURL.String(), nil)
	var searchResp SearchResp
	if err != nil {
		return searchResp, err
	}
	defaultHeaders(&req.Header)
	req.Header.Add("Authorization", c.Authorization)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return searchResp, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return searchResp, err
	}
	if 200 != resp.StatusCode {
		return searchResp, fmt.Errorf("%s", body)
	}
	err = json.Unmarshal(body, &searchResp)
	if err != nil {
		return searchResp, err
	}
	return searchResp, nil
}

func (c *Client) GetEpisode(serieID int, season int, episode int) (GetEpisodeResp, error) {
	path := &url.URL{Path: fmt.Sprintf("/series/%d/episodes/query", serieID)}
	finalURL := c.BaseURL.ResolveReference(path)
	q := finalURL.Query()
	q.Set("airedSeason", strconv.Itoa(season))
	q.Set("airedEpisode", strconv.Itoa(episode))
	finalURL.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", finalURL.String(), nil)
	var getEpisodeResp GetEpisodeResp
	if err != nil {
		return getEpisodeResp, err
	}
	defaultHeaders(&req.Header)
	req.Header.Add("Authorization", c.Authorization)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return getEpisodeResp, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return getEpisodeResp, err
	}
	if 200 != resp.StatusCode {
		return getEpisodeResp, fmt.Errorf("%s", body)
	}
	err = json.Unmarshal(body, &getEpisodeResp)
	if err != nil {
		return getEpisodeResp, err
	}
	return getEpisodeResp, nil
}
