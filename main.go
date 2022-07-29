package main

import (
	"github.com/go-oas/docs"
)

func main() {
	apiDoc := docs.New()

	apiSetInfo(&apiDoc)
	apiSetTags(&apiDoc)
	apiSetServers(&apiDoc)
	apiSetComponents(&apiDoc)

	apiDoc.AttachRoutes([]docs.RouteFn{
		handleCreateUserRoute,
		// handleGetUserRoute,
	})

	apiDoc.Paths = docs.Paths{
		docs.Path{
			Route:       "/user",
			HTTPMethod:  "GET",
			OperationID: "GetUser",
			Responses: docs.Responses{
				getResponseOK(),
				getResponseNotFound(),
			},
			HandlerFuncName: "handleCreateUser",
			RequestBody: docs.RequestBody{
				Description: "Create a new User",
				Content: docs.ContentTypes{
					getContentApplicationJSON("#/components/schemas/User"),
				},
				Required: true,
			},
		},
	}

	if err := apiDoc.BuildDocs(); err != nil {
		panic(err)
	}

	if err := docs.ServeSwaggerUI(&docs.ConfigSwaggerUI{
		Route: "/docs/api/",
		Port:  "3006",
	}); err != nil {
		panic(err)
	}
}

func apiSetInfo(apiDoc *docs.OAS) {
	apiDoc.SetOASVersion("3.0.1")
	apiInfo := apiDoc.GetInfo()
	apiInfo.Title = "Build OAS3.0.1"
	apiInfo.Description = "Description - Builder Testing for OAS3.0.1"
	apiInfo.TermsOfService = "https://smartbear.com/terms-of-use/"
	apiInfo.SetContact("padiazg@gmail.com") // mixed usage of setters ->
	apiInfo.License = docs.License{         // and direct struct usage.
		Name: "MIT",
		URL:  "https://github.com/go-oas/docs/blob/main/LICENSE",
	}
	apiInfo.Version = "1.0.1"
}

func apiSetTags(apiDoc *docs.OAS) {
	// With Tags example you can see usage of direct struct modifications, setter and appender as well.
	apiDoc.Tags.AppendTag(&docs.Tag{
		Name:        "user",
		Description: "Operations about the User",
		ExternalDocs: docs.ExternalDocs{
			Description: "User from the Petstore example",
			URL:         "http://swagger.io",
		},
	})
}
func apiSetServers(apiDoc *docs.OAS) {
	apiDoc.Servers = docs.Servers{
		docs.Server{
			URL: "https://petstore.swagger.io/v2",
		},
		docs.Server{
			URL: "http://httpbin.org",
		},
	}
}

func apiSetComponents(apiDoc *docs.OAS) {
	apiDoc.Components = docs.Components{
		docs.Component{
			Schemas: docs.Schemas{
				docs.Schema{
					Name: "User",
					Type: "object",
					Properties: docs.SchemaProperties{
						docs.SchemaProperty{
							Name:        "id",
							Type:        "integer",
							Format:      "int64",
							Description: "UserID",
						},
						docs.SchemaProperty{
							Name: "username",
							Type: "string",
						},
						docs.SchemaProperty{
							Name: "email",
							Type: "string",
						},
						docs.SchemaProperty{
							Name:        "userStatus",
							Type:        "integer",
							Description: "User Status",
							Format:      "int32",
						},
						docs.SchemaProperty{
							Name: "phForEnums",
							Type: "enum",
							Enum: []string{"placed", "approved"},
						},
					},
					// Ref: "#/components/schemas/User",
					XML: docs.XMLEntry{Name: "User"},
				},
				docs.Schema{
					Name: "Tag",
					Type: "object",
					Properties: docs.SchemaProperties{
						docs.SchemaProperty{
							Name:   "id",
							Type:   "integer",
							Format: "int64",
						},
						docs.SchemaProperty{
							Name: "name",
							Type: "string",
						},
					},
					// Ref: "#/ref",
					XML: docs.XMLEntry{Name: "Tag"},
				},
			},
		},
	}
}

func handleCreateUserRoute(oasPathIndex int, oas *docs.OAS) {
	path := oas.GetPathByIndex(oasPathIndex)

	path.Summary = "Create a new User"
	path.OperationID = "createUser"

	path.RequestBody = docs.RequestBody{
		Description: "Create a new User",
		Content: docs.ContentTypes{
			getContentApplicationJSON("#/components/schemas/User"),
		},
		Required: true,
	}

	path.Responses = docs.Responses{
		getResponseNotFound(),
		getResponseOK(),
	}

	path.Security = docs.SecurityEntities{
		docs.Security{
			AuthName:  "petstore_auth",
			PermTypes: []string{"write:users", "read:users"},
		},
	}

	path.Tags = append(path.Tags, "user")
}
