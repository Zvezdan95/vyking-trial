# Vyking trial 🛡🪓
## Initialization 🏁
- Build images: `docker-compose up -d --build`
- Copy sql file to database container: `docker cp script.sql vyking-trial-db-1:/script.sql`
- Run the sql file: `docker exec -i vyking-trial-db-1 bash -c 'mysql -h 127.0.0.1 -u root -p"${MYSQL_ROOT_PASSWORD}" < /docker-entrypoint-initdb.d/script.sql'`

## Seeding the database 🌱
- Run the seeder: `sudo docker exec -it vyking-trial-web-1 /usr/src/app/seed`

## Util 🛠
- Enter mysql :`docker exec -it vyking-trial-db-1 mysql -u root -p` password in `.env`

## Disclaimers 📃
- I know it is bad practise to push your `.env` file, but to make things easier when running the project. I left it in.
- In order to shorten the initialization step the app will be connecting though root instead of having its own mysql user 