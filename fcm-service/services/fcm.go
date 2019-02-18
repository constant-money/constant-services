package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/constant-money/constant-mvp/fcm-service/config"
)

// FCMService :
type FCMService struct{}

// Send :
func (s FCMService) Send(jsonData map[string]interface{}) (bool, error) {
	conf := config.GetConfig()
	endpoint := conf.GcmEndpoint
	serverKey := conf.GcmServerKey

	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("key=%s", serverKey))

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return false, err
	}

	b, _ := ioutil.ReadAll(response.Body)

	var data map[string]interface{}
	json.Unmarshal(b, &data)

	success, ok := data["success"]

	if ok && (float64(1) == success) {
		return true, nil
	} else {
		return false, errors.New(string(b[:]))
	}
}
