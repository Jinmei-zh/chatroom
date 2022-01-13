package utils

import (
	"chatroom/common/message"
	"encoding/json"
)

func MessageDecode(resData *message.Message, data interface{}) (err error) {
	err = json.Unmarshal([]byte(resData.Data), &data)
	if err != nil {
		return
	}
	return err

}
func MessageEncode(mesType string, mesData interface{}) (dataMessage []byte, err error) {
	data, err := json.Marshal(mesData)
	if err != nil {
		return
	}

	message := message.Message{
		Type: mesType,
		Data: string(data),
	}

	dataMessage, err = json.Marshal(message)
	if err != nil {
		return
	}
	return dataMessage, nil
}
