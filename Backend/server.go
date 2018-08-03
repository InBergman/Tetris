package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Type      string `json:"type"`
}

type Blog struct {
	ID      string `json:"id,omitempty"`
	Account string `json:"account"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func main() {

	fmt.Println("Starting application...")
	session, err := mgo.Dial("mongodb://Idem:!Vincentsuce123@ds135290.mlab.com:35290/cineman")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection done")
	}
	defer session.Close()
	c := session.DB("cineman").C("people")

	// err = c.Insert(
	// 	&Account{ID: "tmpID", Firstname: "Dude", Lastname: "DOUDOU", Type: "Private"},
	// 	&Blog{ID: "BLOG_ID", Account: "Germinator", Title: "CodeForCul", Content: "DuCulEtDuCode", Type: "Private"})

	//	err = c.Find.All(&results)
	//var results []Account

	accountType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Account",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"firstname": &graphql.Field{
				Type: graphql.String,
			},
			"lastname": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"accounts": &graphql.Field{
				Type: graphql.NewList(accountType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if err != nil {
						return nil, err
					}
					var Account []Account
					c.Find(bson.M{}).All(&Account)
					fmt.Println("QUERY === ", Account)
					return Account, nil
				},
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: graphql.Fields{},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Set("Accept", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Range, Content-Disposition, Content-Type")

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":7000", nil)
}
