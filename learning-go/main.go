package main

import (
	"errors"
	"log"
)

/****************** Variables and Functions ************************
func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string ){
	return "something", "else"
} */

/****************** Pointers ************************
func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
	} */

/****************** Recievers: Structs with functions ************************
type User struct{
	FirstName string
	LastName string
	PhoneNumber string
	Age string
	BirthDate time.Time
}

func main() {
	user := User {
		FirstName: "Todd",
		LastName: "Cole",
		PhoneNumber: "440-781-5639",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
} */

/****************** Other Data Structures: Maps and Slices ************************
// type User struct {
// 	FirstName string
// 	LastName string
// }

func main() {
	// A map for strings
	// myMap := make(map[string]string)

	// myMap["dog"] = "Penny"

	// myMap["other-dog"] = "Fido"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	// A map for ints
	// myMap := make(map[string]int)

	// myMap["First"] = 1
	// myMap["Second"] = 2

	// log.Println(myMap["First"], myMap["Second"])

	// A custom map
	// myMap := make(map[string]User)

	// me := User {
	// 	FirstName: "Todd",
	// 	LastName: "Cole",
	// }

	// myMap["me"] = me

	// log.Println(myMap["me"].FirstName)

	// Slices
	// var mySlice []int

	// mySlice = append(mySlice, 2)
	// mySlice = append(mySlice, 1)
	// mySlice = append(mySlice, 3)

	// sort.Ints(mySlice)

	// log.Println(mySlice)

	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	log.Println(numbers)
	log.Println(numbers[6:9])
} */

/****************** Decision Structures ************************
func main() {
	// isTrue := false

	// if isTrue {
	// 	log.Println("isTrue is", isTrue)
	// } else {
	// 	log.Println("isTrue is", isTrue)
	// }
	// cat := "cat"

	// if cat == "cat" {
	// 	log.Println("Cat is cat")
	// } else {
	// 	log.Println("Cat is not cat")
	// }

	// myNum := 100
	// isTrue := false

	// if myNum > 99 && !isTrue {
	// 	log.Println("myNum is greater than 99 and isTrue is set to true")
	// } else if myNum < 100 && isTrue {
	// 	log.Println("1")
	// } else if myNum == 101 || isTrue {
	// 	log.Println("2")
	// } else if myNum > 1000 && isTrue == false {
	// 	log.Println("3")
	// }

	myVar := "horse"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is something else")
	}
} */

/****************** Loops and ranging over data ************************
func main() {
	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}

	// animals := make(map[string]string)
	// animals["dog"] = "Fido"
	// animals["cat"] = "Fluffy"

	// 	for animalType, animal := range animals {
	// 		log.Println(animalType, animal)
	// 	}

	// firstLine := "Once upon a time."

	// for i, l := range firstLine {
	// 	log.Println(i, ":", l)
	// }

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Smith", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 17})

	for _, u := range users {
		log.Println(u.FirstName, u.LastName, u.Email, u.Age)
	}
} */

/****************** Interfaces ************************
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name: "Penny",
		Breed: "Boxer",
	}

	PrintInfo(&dog)

	gorilla := Gorilla {
		Name: "Jock",
		Color: "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Ugh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
} */

/****************** Packages ************************
func main() {
	log.Println("Hello")
}

// go mod init github.com/trcwrx78/projectname
// Used to create your own packages */

/****************** Channels ************************
const numPool = 1000
func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumbers(numPool)
	intChan <- randomNumber
}

func main() {
	intChan  := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
} */

/****************** Reading and Writing JSON ************************
type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Dick"
	m1.LastName = "Grayson"
	m1.HairColor = "brown"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Kory"
	m2.LastName = "Anders"
	m2.HairColor = "red"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
} */

/****************** Writing Tests ************************/
func main() {
	result, err := divide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result of division is", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}

	result = x / y
	return result, nil
}