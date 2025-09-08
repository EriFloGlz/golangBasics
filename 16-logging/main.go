package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Contact struct {
	Name        string
	IsAvailable bool
}
type Contacts []Contact

func printMenu() {
	fmt.Println("****MENU****")
	fmt.Println("1.Add user")
	fmt.Println("2.Print users")
	fmt.Println("3. Exit")
}
func readLine(promp string) (string, error) {
	fmt.Println(promp)
	s, error := reader.ReadString('\n')
	if error != nil {
		return "", fmt.Errorf("Error leyendo input")
	}
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return "", errors.New("Insert a valid input")
	}
	return s, nil
}
func mustReadLine(promp string) string {
	for {
		s, e := readLine(promp)
		if e == nil {
			return s
		}
		fmt.Println(e)
	}
}
func (c *Contacts) isAlreadyRegister(contactName string) (bool, error) {
	for _, c := range *c {
		if c.Name == contactName {
			return false, errors.New("User is already register")
		}
	}
	return false, nil
}
func (c *Contacts) createContact() (Contact, error) {
	name := mustReadLine("Ingresa nombre")

	_, e := c.isAlreadyRegister(name)

	if e != nil {
		return Contact{}, errors.New("Error registrado previante")
	}

	return Contact{
		Name:        name,
		IsAvailable: true,
	}, nil
}

func (c *Contacts) Add(newContact Contact) bool {
	*c = append(*c, newContact)
	return true
}

var nameFile = "contacts.json"

// Guarda los datos en un json
func saveToFile(contacts Contacts) error {
	file, err := os.Create(nameFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(contacts)

	if err != nil {
		return err
	}
	return nil

}

// lee los datos del archivo
func loadFromFile(contacts *Contacts) error {

	file, error := os.Open(nameFile)
	if error != nil {
		return error
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err := decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil

}
func (c *Contacts) print() {
	for _, contact := range *c {
		fmt.Println(contact.Name)
	}
}
func main() {
	var contact Contacts
	slog.Info("Server is starting...")

	error := loadFromFile(&contact)

	if error != nil {
		slog.Error("Error leyendo file")
	}
	slog.Warn("Success reading file")
	for {
		printMenu()
		option := mustReadLine("Insert option: ")
		switch option {
		case "1":
			newContact, error := contact.createContact()
			if error != nil {
				slog.Error("Error creando contacto")
			}
			contact.Add(newContact)
			e := saveToFile(contact)
			if e != nil {
				slog.Error("Error guardando en file")
			}
			slog.Info("Save, success\n")
		case "2":
			contact.print()
		case "3":
			return
		default:

			slog.Info("No valid option")
		}
	}
}
