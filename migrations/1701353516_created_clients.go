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
			"id": "lsbyk2y566v640s",
			"created": "2023-11-30 14:11:56.747Z",
			"updated": "2023-11-30 14:11:56.747Z",
			"name": "clients",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "3goxqivu",
					"name": "name",
					"type": "text",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
						"min": 3,
						"max": 25,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "aq8zfjnf",
					"name": "contact",
					"type": "email",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"exceptDomains": [],
						"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "cmiprfhk",
					"name": "type",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"Client",
							"Prospect"
						]
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

		collection, err := dao.FindCollectionByNameOrId("lsbyk2y566v640s")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
