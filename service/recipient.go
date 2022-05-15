package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	devMod "github.com/dhf0820./ds_model/device"
	recipMod "github.com/dhf0820./ds_model/recipMod"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRecipient(recipientId primitive.ObjectID) (*recipMod.Recipient, error) {
	url := fmt.Sprintf("%s/recipient/%v", RecipientURL(), recipientId.Hex())
	fmt.Printf("#####Final Recipient URL-17: %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = fmt.Errorf("Error in req: %s", err)
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("###error in Recipient request: %s", err)
		return nil, err
	}
	fmt.Printf("\n###Processsing the response\n")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b := string(body)
	fmt.Printf("b: %s\n", b)

	var response recipMod.RecipientResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("##### UnMarshal err: %v\n", err)
	}

	return response.Recipient, err
}

func RecipientURL() string {
	fmt.Printf("GetRecipientURL:51\n")
	ep := GetServiceEndPoint("recipient")
	if ep == nil {
		return ""
	}
	fmt.Printf("RecipienEP:55 -- %s\n", spew.Sdump(ep))
	//url := fmt.Sprintf("%s://%s", ep.Protocol, ep.Address )
	url := ep.Address
	fmt.Printf("Recipient URL:58 --  %s\n", url)
	return url
}

func RecipientDeviceSummaryByName(recip *recipMod.Recipient, name string) (*devMod.DeviceSummary, error) {
	devices := recip.Devices
	for _, dev := range devices {
		if dev.Name == strings.Trim(name, " ") {
			return dev, nil
		}
	}
	return nil, fmt.Errorf("device: %s Not Found", name)
}
