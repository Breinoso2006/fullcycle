package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser

	jsonData := `{"name": "John", "age": 30, "city": "New York", "skills": ["Go", "Java", "Python"]}`
	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\n", v.GetStringBytes("name"))
	fmt.Printf("Age: %d\n", v.GetInt("age"))
	fmt.Printf("City: %s\n", v.GetStringBytes("city"))

	a := v.GetArray("skills")
	for i, skill := range a {
		fmt.Printf("Skill %d: %s\n", i+1, skill.GetStringBytes())
	}
	jsonData = `{"user": {"name": "Alice", "age": 25, "address": {"city": "Los Angeles", "zip": "90001"}}}`
	v, err = p.Parse(jsonData)
	if err != nil {
		panic(err)
	}
	user := v.GetObject("user")
	fmt.Printf("User Name: %s\n", user.Get("name"))
	fmt.Printf("User Age: %d\n", user.Get("age").GetInt())
	address := user.Get("address").GetObject()
	fmt.Printf("City: %s\n", address.Get("city"))
	fmt.Printf("ZIP: %s\n", address.Get("zip"))

	productJson := `{"id": 1, "name": "Laptop", "price": 999.99}`
	var product Product
	if err := json.Unmarshal([]byte(productJson), &product); err != nil {
		panic(err)
	}
	fmt.Printf("Product: %+v\n", product)

}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
