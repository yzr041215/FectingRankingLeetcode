package fecthrating

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchRating(userid string) (float64, error) {
	url := "https://leetcode.cn/graphql/noj-go/"
	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"Content-Type":    "application/json",
		"Referer":         fmt.Sprintf("https://leetcode.cn/u/%s/", userid),
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0",
	}
	payload := map[string]interface{}{
		"query": `
        query ($userSlug: String!) {
            userContestRanking(userSlug: $userSlug) {
                rating
            }
        }
        `,
		"variables": map[string]string{
			"userSlug": userid,
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return 0, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Error: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	data := result["data"].(map[string]interface{})
	userContestRanking := data["userContestRanking"].(map[string]interface{})
	rating := userContestRanking["rating"].(float64)

	return rating, nil
}
