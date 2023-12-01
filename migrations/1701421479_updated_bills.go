package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("yu5bikdkno8ukkf")
		if err != nil {
			return err
		}

		// update
		edit_client := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "nyacearu",
			"name": "client",
			"type": "relation",
			"required": true,
			"presentable": true,
			"unique": false,
			"options": {
				"collectionId": "lsbyk2y566v640s",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_client)
		collection.Schema.AddField(edit_client)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("yu5bikdkno8ukkf")
		if err != nil {
			return err
		}

		// update
		edit_client := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "nyacearu",
			"name": "field",
			"type": "relation",
			"required": true,
			"presentable": true,
			"unique": false,
			"options": {
				"collectionId": "lsbyk2y566v640s",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_client)
		collection.Schema.AddField(edit_client)

		return dao.SaveCollection(collection)
	})
}
