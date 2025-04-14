package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Course struct {
	Name             string  `json:"courseName"`
	Price            float64 `json:"coursePrice"`
	CreatedDate      time.Time
	TotalChapters    int `json:"-"` // this field will be ignored while doing json conversions
	Tags             []string
	CourseInstructor Instructor
	DummyData        []Data `json:"dummyData,omitempty"` // ignores when list is empty
}

type Instructor struct {
	Name string
	Id   int
}

type Data struct {
	Dummy string
	Arr   [3]int
}

func main() {

	fmt.Println("-----> Dealing with Json in GO Lang <----- ")

	EncodeJson()

	DecodeJson()

}

// encode json
func EncodeJson() {
	courses := []Course{
		{Name: "GO Lang", Price: 344.89, CreatedDate: time.Now(), TotalChapters: 67, Tags: []string{"backend", "concurrency"}, CourseInstructor: Instructor{Name: "google", Id: 690}, DummyData: []Data{
			{Dummy: "fake string", Arr: [...]int{1, 2, 3}},
			{Dummy: "fake2 string2", Arr: [...]int{4, 5, 6}},
		}},
		{Name: "Flutter", Price: 500, CreatedDate: time.Now(), TotalChapters: 45, Tags: nil, CourseInstructor: Instructor{Name: "Dart", Id: 435}, DummyData: nil},
	}

	fmt.Printf("Courses Data :%+v\n", courses)

	jsonEncodedCourses, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Json Encoded Data : \n%s\n", jsonEncodedCourses)
	}

	/*
		⚠️ Avoid marshalling runes
		rune is basically a number, a unicode point
	*/
	dataToEncode := []interface{}{"1", '1'}

	for _, eachData := range dataToEncode {
		if bytes, err := json.Marshal(eachData); err == nil {
			fmt.Println("bytes: ", bytes)
			fmt.Println("string(bytes): ", string(bytes))
		}
	}
}

// decodes json
func DecodeJson() {
	fmt.Println("-----> DECODE JSON <-----")
	var jsonCourse Course
	jsonDataFromApi := []byte(`
	{
		"courseName": "Flutter",
		"coursePrice": 500,
		"CreatedDate": "2024-02-26T07:15:58.454359+05:30",
		"Tags": ["tag1", "tag2"],
		"CourseInstructor": {
				"Name": "Dart", 
				"Id": 435
		}
	}	
	`)

	if json.Valid(jsonDataFromApi) {
		json.Unmarshal(jsonDataFromApi, &jsonCourse)
		fmt.Printf("Parsed Course Json : \n %#v\n", jsonCourse)
	} else {
		fmt.Println("Invalid Json")
	}

	/// decode without using any struct
	var decodedJsonMap map[string]interface{}
	json.Unmarshal(jsonDataFromApi, &decodedJsonMap)
	fmt.Printf("Parsed Json Map : \n %#v\n", jsonCourse)
	fmt.Println()
	fmt.Println(decodedJsonMap["courseName"])

	for key, value := range decodedJsonMap {
		fmt.Printf("%v : %v (type %T)\n", key, value, value)
	}

}
