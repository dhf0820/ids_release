package service

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	delvMod "github.com/dhf0820./ds_model/delivery"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
)

func QueueDelivery(delivery *delvMod.Delivery) (*delvMod.Delivery, error) {
	url := fmt.Sprintf("%s/delivery/queue", DeliveryURL())

	delvBody, err := json.Marshal(delivery)
	if err != nil {
		return nil, err
	}
	fmt.Printf("#####Final QueueDelivery URL-21: %s\n", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(delvBody))
	if err != nil {
		err = fmt.Errorf("Error in req: %s", err)
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("###error in Delivery request: %s", err)
		return nil, err
	}
	fmt.Printf("\n###Processsing the response\n")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//b := string(body)
	//fmt.Printf("b: %s\n", b)

	var response delvMod.DeliveryResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("##### UnMarshal err: %v\n", err)
	} else {
		//fmt.Printf("QueueDelivery UnMarshal: %s\n", spew.Sdump(response))
	}
	//var release releasePkg.Release
	//if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
	//	err = fmt.Errorf("Decode - 87 err: %s", err)
	//	return nil, err
	//}
	fmt.Printf("submitRelease Decode:56 -- %s\n", spew.Sdump(response))
	if response.Status != 200 {
		err = fmt.Errorf("%s", response.Message)
		return nil, err
	}
	fmt.Printf("\n###Delivery Successful\n")
	return response.Delivery, nil
}

func DeliveryURL() string {
	fmt.Printf("GetDeliveryURL\n")
	ep := GetServiceEndPoint("delivery")
	if ep == nil {
		return ""
	}
	fmt.Printf("DeliveryEP: %s\n", spew.Sdump(ep))
	//url := fmt.Sprintf("%s://%s", ep.Protocol, ep.Address )
	url := ep.Address
	fmt.Printf("Delivery URL %s\n", url)
	return url
}
