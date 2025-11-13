package main

import (
	"fmt"
	"os"
	"encoding/json"
	"sync"
	"path/filepath"
	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warning(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string
		log logger
	}
)

type Options struct {
	Logger 
}


func New(dir string, options *Options)(*Driver, error){
	dir = file.path.Clean(dir)
	opts := Options{}

}

func (d *Driver) Write() error{

}

func (d *Driver) Read () error {

}

func (d *Driver) ReadAll()(){

}

func (d *Driver) Delete () error {

}

func (d *Driver) getOrCreateMutex () *synx.Mutex {
	
}

type User struct {
		Name string
		Age json.Number
		Contact string
		Company string
		Address Address
	}

type Address struct {
		City string
		State string
		Country string
		Pincode json.Number
	}

func main() {

	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("ERROR CREATING DB:", err)
	}

	employees := []User{
	
			{"John", "23", "2333341242", "CifCif", Address{"New York", "NY", "USA", "10001"}},
			{"Alice", "30", "9876543210", "TechCorp", Address{"San Francisco", "CA", "USA", "94105"}},
			{"Bob", "28", "5556667777", "InnovateX", Address{"Austin", "TX", "USA", "73301"}},
			{"Eve", "35", "4445556666", "SecureIT", Address{"Seattle", "WA", "USA", "98101"}},
			{"Charlie", "32", "1112223333", "DevSolutions", Address{"Boston", "MA", "USA", "02101"}},
	}

	for _, value := range employees {
		db.Write("Users", value.name, User{
			Name: value.Name,
			Age: value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: Address{ City: value.Address.City, State: value.Address.State, Country: value.Address.Country, Pincode: value.Address.Pincode
		}
	)
	}
	
	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{

	}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarchal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}

	fmt.Println((allusers))


	//if err := db.Delete("user," "john") err != nil {
	//	fmt.Println("Error", err)
	//}

//	if err := db.Delete("user", ""); err != nil {
//		fmt.Println("Error", err)
//	}




}