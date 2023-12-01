package populate

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func CreateBills(app *pocketbase.PocketBase) error {
	clients, err := app.Dao().FindRecordsByExpr("clients", dbx.HashExp{"type": "Client"})
	clientsId := make([]string, len(clients))
	for _, client := range clients {
		clientsId = append(clientsId, client.Id)
	}

	bills, err := app.Dao().FindCollectionByNameOrId("bills")
	if err != nil {
		return err
	}

	for i := 0; i < 100000; i++ {
		record := models.NewRecord(bills)
		record.Set("amount", gofakeit.Number(0, 100000))
		record.Set("to_be_paid_at", gofakeit.Date())
		record.Set("paid", gofakeit.Bool())
		record.Set("client", gofakeit.RandomString(clientsId))
		_ = app.Dao().SaveRecord(record)
	}

	return nil
}
