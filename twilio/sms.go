package twilio

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

// SendSMS - sends sms using twilio
func SendSMS(toNumber string, message string) string {

	var accountSID = os.Getenv("TWILIO_SID")
	var authToken = os.Getenv("TWILIO_AUTH_TOKEN")
	var fromNumber = os.Getenv("FROM_TWILIO_NUMBER")

	var urlStr = "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json"

	v := url.Values{}

	v.Set("To", toNumber)
	v.Set("From", fromNumber)
	v.Set("Body", "Brooklyn's in the house!")

	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)

	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)

	return string(resp.Status)

}
