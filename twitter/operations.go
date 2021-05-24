package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ====================================Twitter API v2 Operations==================================== //

// GetTweetByID Twitter API v2 endpoint - retrieve a single Tweet by the provided Tweet ID
//
// API Reference: https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
//
func (op *operatorV2) GetTweetByID(tweetID string, optionalQueryParams map[string]string) (GetTweetByIDResponse, error) {
	if op.http == nil {
		return GetTweetByIDResponse{}, fmt.Errorf("%w the operator's HTTP client was nil", ErrTwitterClient)
	}

	tweetID = strings.TrimSpace(tweetID)
	if tweetID == "" {
		return GetTweetByIDResponse{}, fmt.Errorf("%w tweetID cannot be empty or whitespace-only", ErrTwitterClient)
	}

	if optionalQueryParams == nil {
		optionalQueryParams = make(map[string]string)
	}

	url, err := appendQueryParams(fmt.Sprintf("%s%s/%s", op.cfgV2.BaseURL, endpointV2TweetLookup, tweetID), optionalQueryParams)
	if err != nil {
		return GetTweetByIDResponse{}, fmt.Errorf("%s failed to append query parameters to GetTweetByID HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return GetTweetByIDResponse{}, fmt.Errorf("%s failed to create GetTweetByID HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", op.cfgV2.BearerToken))

	res, err := op.http.Do(req)
	if err != nil {
		return GetTweetByIDResponse{}, fmt.Errorf("%s failed to send GetTweetByID HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	defer res.Body.Close()

	data := GetTweetByIDResponse{}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return GetTweetByIDResponse{}, fmt.Errorf("%s failed to decode GetTweetByID HTTP response body; %w", ErrTwitterClient.Error(), err)
	}

	return data, nil
}

// GetTweets Twitter API v2 endpoint - retrieve Tweets by the provided Tweet IDs
//
// tweetIDs should be a comma delimited list of Tweet IDs with no spaces; max number of IDs is 100
//
// API Reference: https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
//
func (op *operatorV2) GetTweets(tweetIDs []string, optionalQueryParams map[string]string) (GetTweetsResponse, error) {
	if op.http == nil {
		return GetTweetsResponse{}, fmt.Errorf("%w the operator's HTTP client was nil", ErrTwitterClient)
	}

	if len(tweetIDs) == 0 || tweetIDs == nil {
		return GetTweetsResponse{}, fmt.Errorf("%w at least one Tweet ID must be provided to get Tweets", ErrTwitterClient)
	}

	var b strings.Builder

	for _, id := range tweetIDs {
		if _, err := fmt.Fprintf(&b, "%s,", id); err != nil {
			return GetTweetsResponse{}, fmt.Errorf("%s failed to build GetTweets 'id' query parameter; %w", ErrTwitterClient.Error(), err)
		}
	}

	if optionalQueryParams == nil {
		optionalQueryParams = make(map[string]string)
	}

	// Remove trailing comma from tweetIDs query parameter value and assign to query parameter map
	optionalQueryParams["ids"] = b.String()[:b.Len()-1]

	url, err := appendQueryParams(op.cfgV2.BaseURL+endpointV2TweetLookup, optionalQueryParams)
	if err != nil {
		return GetTweetsResponse{}, fmt.Errorf("%s failed to append query parameters to GetTweets HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return GetTweetsResponse{}, fmt.Errorf("%s failed to create GetTweets HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", op.cfgV2.BearerToken))

	res, err := op.http.Do(req)
	if err != nil {
		return GetTweetsResponse{}, fmt.Errorf("%s failed to send GetTweets HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	defer res.Body.Close()

	data := GetTweetsResponse{}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return GetTweetsResponse{}, fmt.Errorf("%s failed to decode GetTweets HTTP response body; %w", ErrTwitterClient.Error(), err)
	}

	return data, nil
}

func (op *operatorV2) GetUserByUsername(username string, optionalQueryParams map[string]string) (GetUserByUsernameResponse, error) {
	if op.http == nil {
		return GetUserByUsernameResponse{}, fmt.Errorf("%w the operator's HTTP client was nil", ErrTwitterClient)
	}

	username = strings.TrimSpace(username)
	if username == "" {
		return GetUserByUsernameResponse{}, fmt.Errorf("%w username must be provided", ErrTwitterClient)
	}

	if optionalQueryParams == nil {
		optionalQueryParams = make(map[string]string)
	}

	url, err := appendQueryParams(fmt.Sprintf("%s%s/%s", op.cfgV2.BaseURL, endpointV2UserLookupByUsernameSingle, username), optionalQueryParams)
	if err != nil {
		return GetUserByUsernameResponse{}, fmt.Errorf("%s failed to append query parameters to GetUserByUsername HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return GetUserByUsernameResponse{}, fmt.Errorf("%s failed to create GetUserByUsername HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", op.cfgV2.BearerToken))

	res, err := op.http.Do(req)
	if err != nil {
		return GetUserByUsernameResponse{}, fmt.Errorf("%s failed to send GetUserByUsername HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	defer res.Body.Close()

	data := GetUserByUsernameResponse{}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return GetUserByUsernameResponse{}, fmt.Errorf("%s failed to decode GetUserByUsername HTTP response body; %w", ErrTwitterClient.Error(), err)
	}

	return data, nil
}

// SearchTweetsRecent Twitter API v2 endpoint - // TODO: add additional documentation
//
// https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent
//
func (op *operatorV2) SearchTweetsRecent(queryParams map[string]string) (SearchTweetsResponse, error) {
	if op.http == nil {
		return SearchTweetsResponse{}, fmt.Errorf("%w the operator's HTTP client was nil", ErrTwitterClient)
	}

	if queryParams == nil {
		queryParams = make(map[string]string, 0)
	}

	url, err := appendQueryParams(op.cfgV2.BaseURL+endpointV2TweetsSearchRecent, queryParams)
	if err != nil {
		return SearchTweetsResponse{},
			fmt.Errorf("%s failed to append queryParams to SearchTweetsRecent HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return SearchTweetsResponse{},
			fmt.Errorf("%s failed to create SearchTweetsRecent HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", op.cfgV2.BearerToken))

	res, err := op.http.Do(req)
	if err != nil {
		return SearchTweetsResponse{},
			fmt.Errorf("%s failed to send SearchTweetsRecent HTTP request; %w", ErrTwitterClient.Error(), err)
	}

	defer res.Body.Close()

	tweets := SearchTweetsResponse{}

	if err := json.NewDecoder(res.Body).Decode(&tweets); err != nil {
		return SearchTweetsResponse{},
			fmt.Errorf("%s failed to decode SearchTweetsRecent HTTP response body; %w", ErrTwitterClient.Error(), err)
	}

	return tweets, nil
}
