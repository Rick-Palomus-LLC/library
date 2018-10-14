package TwilioUtil

import (
	"os"

	twilio "github.com/saintpete/twilio-go"
)

const twilioUrlPart = "https://api.twilio.com/2010-04-01"
const post = "POST"

func SendMessage(body string) (string, error) {
	twilioSid := os.Getenv("TWILIO_SID")
	twilioToken := os.Getenv("TWILIO_TOKEN")

	client := twilio.NewClient(twilioSid, twilioToken)

	return client.Messages.SendMessage(os.Getenv("SMS_NUMBER"), os.Getenv("TWILIO_NUMBER"), body)
}
