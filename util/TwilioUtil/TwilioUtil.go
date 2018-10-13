package TwilioUtil

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type twilioApiData struct {
	twilioSID   string
	twilioToken string
	twilioUrl   string
}

const twilioUrlPart = "https://api.twilio.com/2010-04-01"
const post = "POST"

func SendMessage(body string) (string, error) {
	params := getParams(body)

	requestBody := *strings.NewReader(params.Encode())

	return sendRequest(requestBody)
}

func getParams(body string) url.Values {
	params := url.Values{}

	params.Set("To", os.Getenv("SMS_NUMBER"))
	params.Set("From", os.Getenv("TWILIO_NUMBER"))
	params.Set("Body", body)

	return params
}

func getApiData() twilioApiData {
	twilioSID := os.Getenv("TWILIO_SID")

	return twilioApiData{
		twilioSID:   twilioSID,
		twilioToken: os.Getenv("TWILIO_TOKEN"),
		twilioUrl:   fmt.Sprintf("%s/%s/Messages.json", twilioUrlPart, twilioSID),
	}
}

func sendRequest(requestBody strings.Reader) (string, error) {
	client := &http.Client{}

	apiData := getApiData()

	request, err := http.NewRequest(post, apiData.twilioUrl, &requestBody)

	if err != nil {
		return "", err
	}

	request.SetBasicAuth(apiData.twilioSID, apiData.twilioToken)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)

	return response.Status, nil
}
