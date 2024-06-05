package initializers

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"server/app/gql/resolvers"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"gorm.io/gorm"
)

func GqlHandler(db *gorm.DB) gin.HandlerFunc {
	schema, err := fetchSchema("./app/gql/schemas/")

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

// fetchSchema reads all files in the specified directory and its subdirectories
// and returns their concatenated contents as a string.
func fetchSchema(schemaPath string) (string, error) {
	var schemaContent []byte

	err := filepath.WalkDir(schemaPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, we only want to read files
		if d.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			log.Printf("failed to read file %s: %v", path, err)
			return nil // Continue walking the directory tree
		}
		schemaContent = append(schemaContent, content...)

		return nil
	})

	if err != nil {
		return "", err
	}

	return string(schemaContent), nil
}
