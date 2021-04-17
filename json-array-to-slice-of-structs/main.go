package main

import (
	"encoding/json"
	"fmt"
)


type Welcome struct {
	SendMsg *SendMsg `json:"sendMsg,omitempty"`
	ExampleStruct *ExampleStruct  `json:"example-struct,omitempty"`
}

type SendMsg struct {
	User *string `json:"user,omitempty"`
	Msg  *string `json:"msg,omitempty"`
	Tag *string `json:"tag1,omitempty"`
}

type ExampleStruct struct {
	ExampleField *float64 `json:"example"`
	SecondExampleField *float64 `json:"second-example"`
}

func main() {
	data := []byte(`{
					 "sendMsg":{
						"user":"ANisus",
						"msg":"Trying to send a message"
						},
					 "example-struct":{
					 	"example":45.0,
						"second-example":6.3
					 }
					}`)
	welcome,err := UnmarshalWelcome(data)
	if err != nil {
		panic(err)
	}
	// This will print welcome.SendMsg.Msd address because its not deferenced.
	// That field was defined as a pointer to string type
	//fmt.Printf("%v\n",welcome.SendMsg.Msg)
	//fmt.Printf("%v\n\n",welcome.SendMsg.Tag)
	fmt.Printf("%v\n", *welcome.ExampleStruct.ExampleField)
	fmt.Printf("%v\n\n", *welcome.ExampleStruct.SecondExampleField)

	data2 := []byte(`[{
					 	"sendMsg":{
							"user":"ANisus",
							"msg":"Trying to send a message"
						 },
					 	"example-struct":{
					 		"example":45.0,
							"second-example":6.3
						}
					  },
					  {
					 	"sendMsg":{
							"user":"OTRAPELSONA",
							"msg":"Trying asdadsxage"
						},
					 	"example-struct":{
					 		"example":5.0,
							"second-example":8.03
						}
					  } 
					]`)


	structsSlice, err := UnmarshalWelcomeElements(data2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", structsSlice[1].SendMsg.AnExampleMethod())
	fmt.Printf("%v\n", structsSlice[1].ExampleStruct.AnExampleMethod())

	callElementsMethods := func (w []Welcome) {
		for _ , element := range w {
			element.ExampleStruct.AnExampleMethod()
			element.SendMsg.AnExampleMethod()
		}
	}

	callElementsMethods(structsSlice)
}

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalWelcomeElements(data []byte) ([]Welcome, error) {
	var r = []Welcome{}
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (receiver *SendMsg) AnExampleMethod() string {
	*receiver.Msg = "Changing fields content"
	return *receiver.Msg
}

func (receiver *ExampleStruct) AnExampleMethod() string {
	*receiver.ExampleField = 9
	return "Changing value of ExampleField of ExampleStruct to 9"
}