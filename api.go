package victronapipayg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiVictron = "https://payg.victronenergy.com/api/token"
)

//OpenVictronPaygAPI :
type OpenVictronPaygAPI struct {
}

//NewVictronAPI : Init API
func NewVictronAPI() (OpenVictronPaygAPI, error) {
	return OpenVictronPaygAPI{}, nil
}

type RequestVictron struct {
	Serial          string `json:"serial"`
	Counter         int    `json:"counter"`
	StartingCode    int    `json:"starting_code"`
	PrivateKey      string `json:"private_key"`
	Command         string `json:"command"`
	CommandArgument int    `json:"command_argument"`
	TimeGranularity int    `json:"time_granularity"`
}

type ResponseVictron struct {
	Counter int `json:"counter"`
	Token   int `json:"token"`
}

func (api *OpenVictronPaygAPI) GenerateCustomToken(serial string, startingCode int, privatekey string, counter int, command string, commandArgument, timeGranularity int) (ResponseVictron, error) {
	requestObj := RequestVictron{
		Serial:          serial,
		Counter:         counter,
		StartingCode:    startingCode,
		PrivateKey:      privatekey,
		Command:         command,
		CommandArgument: commandArgument,
		TimeGranularity: timeGranularity,
	}
	jsonStr, _ := json.Marshal(requestObj)

	req, _ := http.NewRequest("POST", apiVictron, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	var respObj ResponseVictron
	res, err := client.Do(req)
	if err != nil {
		return respObj, err
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	if !(res.StatusCode == 200) {
		err = fmt.Errorf(fmt.Sprintf("ERROR Status Code %v", res.StatusCode))
		return respObj, err
	}
	err = json.Unmarshal(b, &respObj)
	if err != nil {
		return respObj, err
	}
	return respObj, nil
}
