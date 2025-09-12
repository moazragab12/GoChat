package main

import "encoding/json"

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"Content"`
}

func (m *Message) ToJSON() []byte {
	data, _ := json.Marshal(m)
	return append(data, '\n')

}
func (m *Message) ToString(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return &msg, err
}
