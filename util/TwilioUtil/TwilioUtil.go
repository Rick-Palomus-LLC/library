package TwilioUtil

import (
	"os"

	twilio "github.com/kevinburke/twilio-go"
)

const twilioUrlPart = "https://api.twilio.com/2010-04-01"
const post = "POST"

func SendMessage(body string) (*twilio.Message, error) {
	twilioSid := os.Getenv("TWILIO_SID")
	twilioToken := os.Getenv("TWILIO_TOKEN")

	client := twilio.NewClient(twilioSid, twilioToken, nil)

	return client.Messages.SendMessage(os.Getenv("TWILIO_NUMBER"), os.Getenv("SMS_NUMBER"), body, nil)
}
