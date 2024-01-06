package tickets

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	filePath = "./tickets.csv"
)

var (
	periods = map[string][2]time.Time{
		"madrugada": {time.Date(0000, time.January, 01, 0, 0, 0, 0, time.UTC), time.Date(0000, time.January, 01, 7, 0, 0, 0, time.UTC)},
		"mañana":    {time.Date(0000, time.January, 01, 7, 0, 0, 0, time.UTC), time.Date(0000, time.January, 01, 13, 0, 0, 0, time.UTC)},
		"tarde":     {time.Date(0000, time.January, 01, 13, 0, 0, 0, time.UTC), time.Date(0000, time.January, 01, 20, 0, 0, 0, time.UTC)},
		"noche":     {time.Date(0000, time.January, 01, 20, 0, 0, 0, time.UTC), time.Date(0000, time.January, 01, 0, 0, 0, 0, time.UTC)},
	}
)

type Ticket struct {
	ID          uint64
	Name        string
	Email       string
	Destination string
	FlightTime  time.Time
	Price       uint64
}

type TicketRepository struct {
	Tickets []Ticket
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

func (repo *TicketRepository) LoadTickets(filePath string) error {
	file, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer file.Close() // Close the file when the function exits

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the CSV file row by row
	for {
		// Read one row from the CSV file
		row, err := reader.Read()

		// Check for the end of the file
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		id, err := strconv.ParseUint(row[0], 10, 64)

		if err != nil {
			return err
		}

		parsedTime, err := time.Parse("15:04", row[4])

		if err != nil {
			return err
		}

		price, err := strconv.ParseUint(row[5], 10, 64)

		if err != nil {
			return err
		}

		// Process the current row (in this example, just print it)
		ticket := Ticket{
			ID:          id,
			Name:        row[1],
			Email:       row[2],
			Destination: row[3],
			FlightTime:  parsedTime,
			Price:       price,
		}

		repo.Tickets = append(repo.Tickets, ticket)
	}

	return nil
}

// ejemplo 1
func (repo *TicketRepository) GetTotalTickets(destination string) (int, error) {

	if len(destination) == 0 {
		return 0, errors.New("destination is empty")
	}

	if len(repo.Tickets) == 0 {
		return 0, errors.New("there are no tickets in the repository")
	}

	count := repo.getDestinationTickets(destination)

	return count, nil

}

func (repo *TicketRepository) getDestinationTickets(destination string) int {
	count := 0

	for _, ticket := range repo.Tickets {
		if containsString(ticket.Destination, destination) {
			count++
		}
	}

	return count
}

// ejemplo 2
func (repo *TicketRepository) GetCountByPeriod(time string) (int, error) {

	time = strings.ToLower(time)

	err := validatePeriod(time)

	if err != nil {
		return 0, err
	}

	count := 0

	for _, ticket := range repo.Tickets {
		if ticket.FlightTime.After(periods[time][0]) && ticket.FlightTime.Before(periods[time][1]) {
			count++
		}
	}

	return count, nil
}

func validatePeriod(time string) error {
	switch time {
	case "madrugada", "mañana", "tarde", "noche":
		return nil
	default:
		return errors.New("invalid period")
	}
}

// ejemplo 3
func (repo *TicketRepository) AverageDestination(destination string) (float64, error) {

	destinationTickets, err := repo.GetTotalTickets(destination)

	if err != nil {
		return 0, err
	}

	return float64(destinationTickets) / float64(len(repo.Tickets)), nil
}

// containsString checks if a string is present in a slice of strings
func containsString(country string, target string) bool {
	return strings.EqualFold(country, target)
}
