package main

import (
	"github.com/go-oas/docs"
)

func main() {
	apiDoc := docs.New()

	apiSetInfo(&apiDoc)

	if err := apiDoc.BuildDocs(); err != nil {
		panic(err)
	}

	if err := docs.ServeSwaggerUI(&docs.ConfigSwaggerUI{
		Route: "/docs/api/",
		Port:  "3005",
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
	apiInfo.SetContact("aleksandar.nesovic@protonmail.com") // mixed usage of setters ->
	apiInfo.License = docs.License{                         // and direct struct usage.
		Name: "MIT",
		URL:  "https://github.com/go-oas/docs/blob/main/LICENSE",
	}
	apiInfo.Version = "1.0.1"
}
