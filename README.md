# Vyking trial 🛡🪓

## Initialization 🏁

- Build images: `docker-compose up -d --build`
- Copy SQL file to database container: `docker cp db/script.sql vyking-trial-db:/docker-entrypoint-initdb.d/script.sql`
- Run the SQL file: `docker exec -i vyking-trial-db bash -c 'mysql -h 127.0.0.1 -u root -p"${DB_ROOT_PASSWORD}" < /docker-entrypoint-initdb.d/script.sql'`

## Seeding the database 🌱

- Run the seeder: `sudo docker exec -it vyking-trial-web /usr/src/app/seed`

## Endpoints 🎯

- All user ranked by their current balance: http://localhost:3000/ranking
- All user ranked by their amount they spent in the current tournament: http://localhost:3000/ranking

## Util 🛠

- Enter mysql: `docker exec -it vyking-trial-db mysql -u root -p` (password in `.env`)

## Disclaimers 📃

- I know it is bad practice to push your `.env` file, but to make things easier when running the project, I left it in.
- In order to shorten the initialization step, the app will be connecting through root instead of having its own MySQL user.
- I assumed there would be one tournament ending daily since I made a cron job that triggers the stored procedure daily, and the stored procedure checks if there is a tournament ending today.

## Questions and Answers 🤔

### What did you learn and if you encountered any challenges, how did you overcome them?

- **Building the seeder:**  Initially, I had trouble building the seeder executable within the Docker container due to issues with static linking and the `scratch` image. I overcame this by using a multi-stage build with the same base image for both building and running the seeder, and by carefully managing dependencies and build flags.
- **Executing the SQL script:** I encountered difficulties in getting the SQL script to execute correctly during the MySQL container startup. I explored various approaches, including using the `/docker-entrypoint-initdb.d/` directory, mounting volumes, and manually executing the script. I ultimately settled on a manual execution approach for better control and flexibility.

### What did you take into consideration to ensure the query and/or stored procedure is efficient and handles edge cases?

* **Efficient queries:** I used window functions instead of CTEs for ranking, as they are generally more efficient for simple ranking scenarios. I also avoided unnecessary joins and indexes in the database schema.
* **Edge cases:** In the stored procedure, I handled cases where there might be no active tournament or no bets within the active tournament's date range.
* **Data integrity:**  I stored monetary values as integers in cents to avoid potential floating-point inaccuracies.

### CTEs and Window functions: If you used CTEs or Window Functions, what did you learn about their power and flexibility? How might you apply the technique in more complex scenarios?

I utilized window functions for ranking players based on their bet amounts. I find them to be a valuable tool for performing calculations across rows, especially when dealing with rankings or aggregations within partitions. In this specific case, they provided an efficient way to rank players without the need for self-joins or subqueries.
* Calculate running totals or averages.
* Segment data into groups based on specific criteria.
* Perform time-series analysis, such as calculating moving averages or identifying trends.

### Optimization: How did you optimize the queries and stored procedures?

* **Indexing:** I created indexes on relevant columns to speed up query execution.
* **Appropriate data types:** I used appropriate data types for efficient storage and retrieval.
* **Avoiding unnecessary joins:** I removed unnecessary joins in the database schema and the stored procedure.
* **Window functions:** I used window functions for efficient ranking.
* **Limiting result sets:** I used `LIMIT` clauses where appropriate to reduce the amount of data processed.
* **Error handling:** I added error handling to the stored procedure to handle potential issues gracefully.