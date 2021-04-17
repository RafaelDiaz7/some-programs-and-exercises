package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	SendMsg *SendMsg `json:"sendMsg,omitempty"`
	Say     *string  `json:"say,omitempty"`
}

type SendMsg struct {
	User *string `json:"user,omitempty"`
	Msg  *string `json:"msg,omitempty"`
	Tag *string `json:"tag1"`
}

func main() {
	data := []byte(`{
					 "sendMsg":{
						"user":"ANisus",
						"msg":"Trying to send a message"
						},
					 "say":"Hello"
					}`)

	data2 := []byte(`{
					 "sendMsg":{
						"user":"ANisus",
						"msg":"Trying to send a message that is different",
						"tag1":"one-tag"
						},
					 "say":"Hello"
					}`)


	welcome,err := UnmarshalWelcome(data)
	if err != nil {
		log.Fatal(err)
	}
	// This will print welcome.SendMsg.Msd address because its not deferenced. That field was defined as a pointer to string type
	fmt.Printf("%v\n",welcome.SendMsg.Msg)
	fmt.Printf("%v\n\n",welcome.SendMsg.Tag)

	welcome2, err := UnmarshalWelcome(data2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n",*welcome2.SendMsg.Msg)
	fmt.Printf("%v",*welcome2.SendMsg.Tag)
}
