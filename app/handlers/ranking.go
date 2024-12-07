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
	if len(ranking) == 0 {
		// If no rows, return an empty JSON array
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
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

type RankByTournamentRow struct {
	Id                        int    `json:"player_id"`
	UserName                  string `json:"user_name"`
	AccountBalance            int64  `json:"account_balance"`
	TotalBetsMadeInTournament int64  `json:"total_bets_made"`
	Rank                      int    `json:"rank"`
}

// GetTournamentRanking returns the ranking list for the current tournament as JSON
func GetTournamentRanking(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
        SELECT p.player_id,
			   p.player_name,
			   p.account_balance,
			   SUM(b.amount)                             AS total_bets_made,
			   RANK() OVER (ORDER BY SUM(b.amount) DESC) AS player_rank
		FROM players p
				 JOIN bets b ON p.player_id = b.player_id
		WHERE b.bet_time
				  BETWEEN (SELECT start_date FROM tournaments WHERE DATE(end_date) = CURDATE() LIMIT 1)
				  AND (SELECT end_date FROM tournaments WHERE DATE(end_date) = CURDATE() LIMIT 1)
		GROUP BY p.player_id;
	`)
	if err != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ranking []RankByTournamentRow

	for rows.Next() {
		var row RankByTournamentRow

		err := rows.Scan(
			&row.Id,
			&row.UserName,
			&row.AccountBalance,
			&row.TotalBetsMadeInTournament,
			&row.Rank)
		if err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}

		ranking = append(ranking, row)
	}

	if len(ranking) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
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
