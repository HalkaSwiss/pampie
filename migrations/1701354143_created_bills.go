package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "yu5bikdkno8ukkf",
			"created": "2023-11-30 14:22:23.913Z",
			"updated": "2023-11-30 14:22:23.913Z",
			"name": "bills",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "weqtlvnn",
					"name": "amount",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "m8ofjote",
					"name": "to_be_paid_at",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "puy42l6j",
					"name": "paid",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
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
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("yu5bikdkno8ukkf")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
