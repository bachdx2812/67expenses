package initializers

import (
	"log"
	"os"
	"server/app/gql/resolvers"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"gorm.io/gorm"
)

func GqlHandler(db *gorm.DB) gin.HandlerFunc {
	schema, err := fetchSchema("./gql/schemas/")

	if err != nil {
		log.Fatalf("failed to get schema: %v", err)
	}
	opts := []graphql.SchemaOpt{graphql.UseStringDescriptions(), graphql.UseFieldResolvers()}
	gqlSchema := graphql.MustParseSchema(schema, &resolvers.Resolver{Db: db}, opts...)

	return ginSchemaHandler(gqlSchema)
}

func ginSchemaHandler(gqlSchema *graphql.Schema) gin.HandlerFunc {
	handler := &relay.Handler{Schema: gqlSchema}

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func fetchSchema(schemaPath string) (string, error) {
	entries, err := os.ReadDir(schemaPath)
	if err != nil {
		return "", err
	}

	var schemaContent []byte

	for _, entry := range entries {
		filePath := schemaPath + entry.Name()
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("failed to read file %s: %v", filePath, err)
			continue
		}
		schemaContent = append(schemaContent, content...)
	}

	return string(schemaContent), nil
}
