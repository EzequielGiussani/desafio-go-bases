package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	t := tickets.Repository

	err := t.LoadTickets("tickets.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	count, err := tickets.GetTotalTickets("Brazil")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Brazil count: ", count)

	c, err := tickets.GetCountByPeriod("madrugada")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Period count: ", c)

	tot, err := tickets.AverageDestination("Brazil")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Average: ", tot)

}
