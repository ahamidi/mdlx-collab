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
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-08-25 19:11:12.405Z",
				"updated": "2023-08-25 19:15:26.060Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "uugurw0e",
						"name": "roles",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "vyuri5l2fr91gly",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": false
				}
			},
			{
				"id": "vyuri5l2fr91gly",
				"created": "2023-08-25 19:15:26.060Z",
				"updated": "2023-08-25 19:15:26.060Z",
				"name": "roles",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "grhzfuhx",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "smdcchyt",
						"name": "permission_set",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "rhrg6kxx8u35mml",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
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
			},
			{
				"id": "ywag7fj3itq9nxm",
				"created": "2023-08-25 19:15:26.060Z",
				"updated": "2023-08-25 19:15:26.060Z",
				"name": "permission_granular",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "neuuhrai",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "afnpj0v0",
						"name": "description",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
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
			},
			{
				"id": "rhrg6kxx8u35mml",
				"created": "2023-08-25 19:15:26.060Z",
				"updated": "2023-08-25 19:15:26.060Z",
				"name": "permission_set",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "wukpdyph",
						"name": "entity_type",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "ru4ozbvd",
						"name": "entity_id",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "yhe7yx2k",
						"name": "permissions",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "ywag7fj3itq9nxm",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
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
			},
			{
				"id": "leae1atb799bxn7",
				"created": "2023-08-25 19:15:26.060Z",
				"updated": "2023-08-25 19:15:26.060Z",
				"name": "resources",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "75c4jwlt",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "3n9ihahh",
						"name": "type",
						"type": "select",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"postgres",
								"mysql",
								"s3"
							]
						}
					},
					{
						"system": false,
						"id": "38jzuvje",
						"name": "credentials",
						"type": "json",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\" &&\n@request.auth.roles.permission_set.entity_type = 'resources' &&\n@request.auth.roles.permission_set.permissions.name ~ \"list\"",
				"viewRule": "@request.auth.id != \"\" &&\n@request.auth.roles.permission_set.entity_type = 'resources' &&\n@request.auth.roles.permission_set.permissions.name ~ \"read\"",
				"createRule": "@request.auth.id != \"\" &&\n@request.auth.roles.permission_set.entity_type = 'resources' &&\n@request.auth.roles.permission_set.permissions.name ~ \"create\"",
				"updateRule": "@request.auth.id != \"\" &&\n@request.auth.roles.permission_set.entity_type = 'resources' &&\n@request.auth.roles.permission_set.permissions.name ~ \"update\"",
				"deleteRule": "@request.auth.id != \"\" &&\n@request.auth.roles.permission_set.entity_type = 'resources' &&\n@request.auth.roles.permission_set.permissions.name ~ \"delete\"",
				"options": {}
			},
			{
				"id": "mnsxayc16ipntlw",
				"created": "2023-08-25 20:57:09.082Z",
				"updated": "2023-08-25 20:57:09.082Z",
				"name": "invites",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ub9ay5ok",
						"name": "inviter",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "k6uhiawm",
						"name": "invited_user",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "dtrzvu8i",
						"name": "token",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "upmqocgj",
						"name": "accepted",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
