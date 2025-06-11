package main

import (
	"cache_API/db"
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request, ctx context.Context, database *sql.DB) {
	if len(r.URL.Query()) > 0 {
		handleGet(w, r, ctx, database)
	} else {
		handleLanding(w, ctx, database)
	}	
}

func handleLanding(w http.ResponseWriter, ctx context.Context, database *sql.DB ) {
	cacheValues := GetAllCached(ctx)
	dbValues, err := db.GetAll(ctx, database)
	if err != nil {
		http.Error(w, "Database error: "+ err.Error(), http.StatusInternalServerError)
		return
	}


	fmt.Println("db records: ")
	    for key, value := range dbValues {
        fmt.Printf("%s: %s\n", key, value)
    }

		fmt.Println("redis records: ")
	data := map[string]interface{}{
		"Title":   "Results",
		"Header":  "Records",
		"RecordsCache": cacheValues,
		"RecordsDB": dbValues,
	}

	parseDataOnPage(w, data)
}

func handleGet(w http.ResponseWriter, r *http.Request, ctx context.Context, database *sql.DB) {
	cacheValues, _ := getResultsFromQuery(r, ctx, database)

	data := map[string]interface{}{
		"Title":   "Results",
		"Header":  "Cached records from GET",
		"RecordsCache": cacheValues,
	}

	parseDataOnPage(w, data)
}

func getResultsFromQuery(r *http.Request, ctx context.Context, database *sql.DB) (map[string]string, string) {
	queryMap := make(map[string]string)
	var err error
	var errMessage string
	queryParams := r.URL.Query()

	for key, values := range queryParams {
		for _, value := range values {
			queryMap[key], err = GetOrSaveRecord(ctx, database, key, value)
			if err != nil {
				errMessage = "Error Occured"
				break
			}
		}
	}

	return queryMap, errMessage
}

func parseDataOnPage(w http.ResponseWriter, data map[string]interface{}) {
	tmpl := template.Must(template.New("page").Parse(tpl))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
