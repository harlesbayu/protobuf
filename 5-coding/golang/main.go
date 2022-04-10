package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "github.com/harlesbayu/protobuf/complex"
	enumpb "github.com/harlesbayu/protobuf/enum_example"

	simplepb "github.com/harlesbayu/protobuf/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	//sm := doSimple()
	//readAndWriteDemo(sm)
	//jsonDemo(sm)

	//doEnum()

	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Dummy Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second Dummy Message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third Dummy Message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           1,
		DayOfTheWeek: enumpb.DayOfTheWeek_SATURDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_SUNDAY
	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct", sm2)
}

func toJSON(pb proto.Message) string {
	jsonString := protojson.Format(pb)
	return jsonString
}

func fromJSON(in string, pb proto.Message) {
	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		fmt.Println("Couldn't unmarshall the JSON into the pb document")
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)
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

func doSimple() *simplepb.SimpleMessage {
	sm := &simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "Harles",
		SampleList: []int32{1, 2, 3, 4},
	}

	sm.Name = "Harles Bayu"
	return sm
}
