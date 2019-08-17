package utils

import (
	"bytes"
	"encoding/json"
	"gopkg.in/resty.v1"
	"istio-service-mesh/constants"
	"istio-service-mesh/types"
	"net/http"
	"reflect"
)

func SendLog(msg, message_type, env_id string) (int, error) {

	var data types.LoggingRequest

	data.Id = env_id
	data.Service = constants.SERVICE_NAME
	data.Environment = "environment"
	data.Level = message_type
	data.Message = msg

	response := PostNotify(constants.LoggingURL+constants.LOGGING_ENDPOINT, data)
	if response.Error != nil {
		Info.Println(response.Error)
		return 400, response.Error
	}
	return response.StatusCode, response.Error

}

func Notify_Generic(state interface{}, path string) {

	url := path
	Info.Println("Notifying front end: URL: ", url)
	b, err1 := json.Marshal(state)
	if err1 != nil {
		Info.Println(err1)
	}

	_ = json.Unmarshal(b, &state)
	b1, _ := json.Marshal(state)
	Info.Println("notification payload:\n", string(b1))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))

	req.Header.Set("Content-Type", "application/json")

	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	client := &http.Client{}
	//client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Info.Println(err)
		Info.Println(reflect.TypeOf(resp))

	} else {
		statusCode := resp.StatusCode
		Info.Printf("notification status code %d\n", statusCode)

		resp.Body.Close()

	}

}
func PostNotify(url string, data interface{}) types.ResponseData {
	b, err1 := json.Marshal(data)
	if err1 != nil {
		Info.Println(err1)
	}
	Info.Println("notification endpoint:", url)
	Info.Println("notification payload:", string(b))
	req := resty.New()

	resp, err := req.R().SetBody(data).SetHeader("Content-Type", "application/json").Post(url)
	if err != nil {
		Error.Println(err)
		return types.ResponseData{Error: err}
	}
	Info.Println("responseCode: ", resp.StatusCode(), "\n response Body", string(resp.Body()))
	return types.ResponseData{StatusCode: resp.StatusCode(), Status: resp.Status(), Body: string(resp.Body())}
}
