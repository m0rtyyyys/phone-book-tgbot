package main

import (
	"context"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FullInfo struct {
	Name        string
	Surname     string
	Pantronymic string
	City        string
	Street      string
	HouseNum    string
	Building    string
	Flat        int
	TNumber     string
}

func ParseInput(input string) (FullInfo, error) {
	info := make([]string, 0, 9)
	strings.Split(input, " ")
	person := FullInfo{}
	person.Surname = info[0]
	person.Name = info[1]
	if info[2] == "_" {
		person.Pantronymic = ""
	} else {
		person.Pantronymic = info[2]
	}
	person.Street = info[3]
	person.HouseNum = info[4]
	if info[5] == "_" {
		person.Building = ""
	} else {
		person.Building = info[5]
	}
	if info[6] == "_" {
		person.Flat = 0
	} else {
		var err error
		person.Flat, err = strconv.Atoi(info[6])
		if err != nil {
			return FullInfo{}, err
		}
	}
	person.TNumber = info[7]
	return person, nil
}

func AddUser(ctx context.Context, pool *pgxpool.Pool, person FullInfo) {

}
