package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	devMod "github.com/dhf0820./ds_model/device"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDevice(deviceId primitive.ObjectID) (*devMod.Device, error) {
	url := fmt.Sprintf("%s/device/%v", DeviceURL(), deviceId.Hex())
	fmt.Printf("#####Final  GetDevice-13 URL: %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = fmt.Errorf("Error in req: %s", err)
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("###error in Device request: %s", err)
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

	var response devMod.DeviceResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("##### UnMarshal err: %v\n", err)
	} else {
		fmt.Printf("GetDevice UnMarshal: %s\n", spew.Sdump(response))
	}
	//if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
	//	err = fmt.Errorf("Decode - 87 err: %s", err)
	//	return nil, err
	//}
	//fmt.Printf("GetRelease Decode: %s\n", spew.Sdump(release))
	return response.Device, nil

	return nil, nil
}

func DeviceURL() string {
	fmt.Printf("GetDeviceURL\n")
	ep := GetServiceEndPoint("device")
	if ep == nil {
		return ""
	}
	fmt.Printf("DeviceEP: %s\n", spew.Sdump(ep))
	//url := fmt.Sprintf("%s://%s", ep.Protocol, ep.Address )
	url := ep.Address
	fmt.Printf("Device URL: %s\n", url)
	return url
}
