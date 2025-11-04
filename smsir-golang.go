package smsir_golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SMSIR struct {
	apiToken   string
	httpClient http.Client
}

type SMSIRVerificationRequestParametersDTO struct {
	Name  string
	Value string
}

type SMSIRVerificationRequestDTO struct {
	PhoneNumber string
	TemplateID  string
	Parameters  []SMSIRVerificationRequestParametersDTO
}

type SMSIRResponseDataDTO struct {
	MessageID int64
	Cost      float32
}

type SMSIRResponseDTO struct {
	Status  int
	Message string
	Data    SMSIRResponseDataDTO
}

func Init(smsir *SMSIR, apiToken string) {
	smsir.apiToken = apiToken
	smsir.httpClient = http.Client{}
}

func SendVerificationCode(smsir *SMSIR, request SMSIRVerificationRequestDTO) (*SMSIRResponseDTO, error) {
	endpoint := "https://api.sms.ir/v1/send/verify"
	requestBodyJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("ACCEPT", "application/json")
	req.Header.Set("X-API-KEY", smsir.apiToken)
	response, err := smsir.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer response.Body.Close()
	var responseBody SMSIRResponseDTO
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return nil, err
	}
	return &responseBody, nil
}
