package main

import (
	"bufio"
	"fmt"
	"os"
	"petstore/cmd/connection"
	"petstore/manage/client"
	"petstore/manage/pet"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions() {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose an option (a - add pet, c - add client, s - add tip): ", reader)

	switch opt {
	case "a":
		newPet := pet.Pet{}
		animal, _ := getInput("Pet - Animal: ", reader)
		price, _ := getInput("Pet - Price: ", reader)
		age, _ := getInput("Pet - Age: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions()
		}

		a, err := strconv.ParseInt(age, 10, 0)

		if err != nil {
			fmt.Println("Age must be a number")
			promptOptions()
		}
		newPet.Animal = animal
		newPet.Price = p
		newPet.Age = a

		connection.SavePet(newPet)

	case "c":
		newClient := client.Client{}
		name, _ := getInput("Client - Name: ", reader)
		lastName, _ := getInput("Client - Last Name: ", reader)
		email, _ := getInput("Client - Email: ", reader)
		phone, _ := getInput("Client - Phone Number: ", reader)
		address, _ := getInput("Client - Address: ", reader)
		gender, _ := getInput("Client - gender: ", reader)
		age, _ := getInput("Client - Age: ", reader)

		a, err := strconv.ParseInt(age, 10, 0)
		if err != nil {
			fmt.Println("Age must be a number")
			promptOptions()
		}

		newClient.Name = name
		newClient.LastName = lastName
		newClient.Email = email
		newClient.Phone = phone
		newClient.Address = address
		newClient.Gender = gender
		newClient.Age = a

		connection.SaveClient(newClient)
	}
}

func main() {
	promptOptions()
}
