package main

import "strconv"

type Contact struct {
	Id           int
	FirstName    string
	LastName     string
	EmailAddress string
	ZipCode      string
	Address      string
}

type Match struct {
	ContactIdSource int
	ContactIdMatch  int
	Accuracy        string
}

func LoadContact(record []string) *Contact {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		id = -1
	}
	return &Contact{
		Id:           id,
		FirstName:    record[1],
		LastName:     record[2],
		EmailAddress: record[3],
		ZipCode:      record[4],
		Address:      record[5],
	}
}
