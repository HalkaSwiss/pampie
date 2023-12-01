package populate

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func CreateClients(app *pocketbase.PocketBase) error {
	clients, err := app.Dao().FindCollectionByNameOrId("clients")
	if err != nil {
		return err
	}

	for i := 0; i < 5000; i++ {
		record := models.NewRecord(clients)
		record.Set("name", gofakeit.Name())
		record.Set("contact", gofakeit.Email())
		record.Set("type", gofakeit.RandomString([]string{"Client", "Prospect"}))
		_ = app.Dao().SaveRecord(record)
	}

	return nil
}
