package handlers

import (
	"encoding/json"
	"net/http"
	"vyking-trial/app/database"
)

type RankRow struct {
	UserName       string `json:"user_name"`
	AccountBalance int64  `json:"account_balance"`
	Rank           int    `json:"rank"`
}

// GetRanking returns the ranking list as JSON, fetching data from the database
func GetRanking(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT
		    p.player_name,
		    p.account_balance,
		    RANK() OVER (ORDER BY p.account_balance DESC) AS player_rank
		FROM
		    players p;
	`)
	if err != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if !rows.Next() {
		// If no rows, return an empty JSON array
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
	}

	var ranking []RankRow

	for rows.Next() {
		var row RankRow
		err := rows.Scan(&row.UserName, &row.AccountBalance, &row.Rank)
		if err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		ranking = append(ranking, row)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(ranking)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}
