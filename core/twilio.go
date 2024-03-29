package core

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

// SendSMS - sends sms using twilio
func SendSMS(done chan bool, toNumber string, message string) (string, error) {

	var accountSID = os.Getenv("TWILIO_SID")
	var authToken = os.Getenv("TWILIO_AUTH_TOKEN")
	var fromNumber = os.Getenv("FROM_TWILIO_NUMBER")

	var urlStr = "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json"

	v := url.Values{}

	v.Set("To", toNumber)
	v.Set("From", fromNumber)
	v.Set("Body", message)

	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)

	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	done <- true
	return string(resp.Status), nil

}
