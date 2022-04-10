package main

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"log"

	address_bookpb "github.com/harlesbayu/protobuf_exercise/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	ap := generatePerson()
	//readAndWriteDemo(ap)
	jsonDemo(ap)
}

func jsonDemo(pb proto.Message) {
	apAsString := toJSON(pb)
	fmt.Println(apAsString)

	ap2 := &address_bookpb.Person{}
	fromJSON(apAsString, ap2)
	fmt.Println("Successfully created proto struct", ap2)
}

func fromJSON(in string, pb proto.Message) {
	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		fmt.Println("Couldn't unmarshall the JSON into the pb document")
	}
}

func toJSON(pb proto.Message) string {
	jsonString := protojson.Format(pb)
	return jsonString
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("address_book.bin", sm)
	ap2 := &address_bookpb.Person{}
	readFromFile("address_book.bin", ap2)
	fmt.Println("Read the content:", ap2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise tp bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been writen!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}
	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffer struct", err)
		return err
	}
	return nil
}

func generatePerson() *address_bookpb.Person {
	p := &address_bookpb.Person{
		Id:    1,
		Name:  "Harles",
		Email: "harles@mail.com",
		Phones: []*address_bookpb.Person_PhoneNumber{
			&address_bookpb.Person_PhoneNumber{
				Number: "12345",
				Type:   address_bookpb.Person_PhoneType(2),
			},
			&address_bookpb.Person_PhoneNumber{
				Number: "123456",
				Type:   address_bookpb.Person_PhoneType(1),
			},
		},
		LastUpdated: timestamppb.Now(),
	}
	return p
}
