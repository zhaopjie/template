/**
 * @Author: duke
 * @Description:
 * @File:  graphql
 * @Version: 1.0.0
 * @Date: 2020/6/22 2:32 下午
 */

package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

func S() {
	field := graphql.Fields{
		"hello": &graphql.Field{
			Name:              "",
			Type:              nil,
			Args:              nil,
			Resolve:           nil,
			DeprecationReason: "",
			Description:       "",
		},
	}

	Query := graphql.ObjectConfig{
		Name:        "QUERY",
		Interfaces:  nil,
		Fields:      field,
		IsTypeOf:    nil,
		Description: "测试接口",
	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(Query),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	// Query
	query :=`
    { hello
    }
    `
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // { “ data ” :{ “ hello ” : ” world ” }}
}
