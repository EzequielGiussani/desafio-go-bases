package tickets_test

import (
	"testing"
	"time"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
)

func TestLoadTickets(t *testing.T) {
	//arrange
	repo := tickets.NewTicketRepository()

	//act/assert
	t.Run("Should return error when file does not exist", func(t *testing.T) {
		err := repo.LoadTickets("fail.csv")
		require.Error(t, err)
		require.EqualError(t, err, "open fail.csv: no such file or directory")
	})

	t.Run("Should return error when file is unreadable by reader", func(t *testing.T) {
		err := repo.LoadTickets("./testing/reader_error.csv")
		require.Error(t, err)
		require.Contains(t, err.Error(), "ParseUint")
	})

	t.Run("Should return error when file doesn't have uint in row[0]", func(t *testing.T) {
		err := repo.LoadTickets("./testing/row_0_error.csv")
		require.Error(t, err)
		require.Contains(t, err.Error(), "ParseUint")
	})

	t.Run("Should return error when file doesn't have correct time format in row[4]", func(t *testing.T) {
		err := repo.LoadTickets("./testing/row_4_error.csv")
		require.Error(t, err)
		require.Contains(t, err.Error(), "parsing time")
	})

	t.Run("Should return error when file doesn't have uint in row[5]", func(t *testing.T) {
		err := repo.LoadTickets("./testing/row_5_error.csv")
		require.Error(t, err)
		require.Contains(t, err.Error(), "ParseUint")
	})

	t.Run("Should successfully load tickets from file", func(t *testing.T) {
		err := repo.LoadTickets("./testing/tickets_test.csv")
		require.NoError(t, err)
		require.Equal(t, 5, len(repo.Tickets))
		require.Equal(t, uint64(1), repo.Tickets[0].ID)
		require.Equal(t, "Tait Mc Caughan", repo.Tickets[0].Name)
		require.Equal(t, "tmc0@scribd.com", repo.Tickets[0].Email)
		require.Equal(t, "Finland", repo.Tickets[0].Destination)
		require.Equal(t, "Finland", repo.Tickets[0].Destination)
		require.Equal(t, time.Date(0000, time.January, 01, 17, 11, 0, 0, time.UTC), repo.Tickets[0].FlightTime)
		require.Equal(t, uint64(785), repo.Tickets[0].Price)
	})

}
func TestGetTotalTickets(t *testing.T) {
	//arrange
	var emptyRepo tickets.TicketRepository = tickets.TicketRepository{}

	//act/assert
	t.Run("Should return error when repo is empty", func(t *testing.T) {
		_, err := emptyRepo.GetTotalTickets("test")
		require.Error(t, err)
		require.Equal(t, 0, len(emptyRepo.Tickets))
		require.EqualError(t, err, "there are no tickets in the repository")
	})

	//arrange
	repository := tickets.NewTicketRepository()
	repository.LoadTickets("./testing/tickets_test.csv")

	//act/assert
	t.Run("Should return error when destination is empty", func(t *testing.T) {
		_, err := repository.GetTotalTickets("")
		require.Error(t, err)
		require.EqualError(t, err, "destination is empty")
	})

	t.Run("Should be successful", func(t *testing.T) {
		total, err := repository.GetTotalTickets("China")
		require.NoError(t, err)
		require.Equal(t, 2, total)
	})

}

func TestGetCountByPeriod(t *testing.T) {
	//arrange
	repo := tickets.NewTicketRepository()
	repo.LoadTickets("./testing/tickets_test.csv")

	//act/assert

	t.Run("Should return error invalid period", func(t *testing.T) {
		count, err := repo.GetCountByPeriod("test")
		require.Error(t, err)
		require.Equal(t, 0, count)
		require.EqualError(t, err, "invalid period")
	})

	t.Run("Should return successfully", func(t *testing.T) {
		count, err := repo.GetCountByPeriod("tarde")
		require.NoError(t, err)
		require.Equal(t, 2, count)
	})
}

func TestAverageDestination(t *testing.T) {

	//arrange
	repo := tickets.NewTicketRepository()

	//act/assert
	t.Run("Should return error when repo is empty", func(t *testing.T) {
		_, err := repo.AverageDestination("test")
		require.Error(t, err)
		require.Equal(t, 0, len(repo.Tickets))
		require.EqualError(t, err, "there are no tickets in the repository")
	})

	//arrange
	repo.LoadTickets("./testing/tickets_test.csv")

	//act/assert
	t.Run("Should return error when destination is empty", func(t *testing.T) {
		_, err := repo.AverageDestination("")
		require.Error(t, err)
		require.EqualError(t, err, "destination is empty")
	})

	t.Run("Should be successful", func(t *testing.T) {
		total, err := repo.AverageDestination("China")
		require.NoError(t, err)
		require.Equal(t, 0.4, total)
	})

}
