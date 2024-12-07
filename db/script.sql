CREATE DATABASE IF NOT EXISTS igaming_platform;

USE igaming_platform;

CREATE TABLE IF NOT EXISTS players
(
    player_id       INT PRIMARY KEY AUTO_INCREMENT,
    player_name     VARCHAR(255)        NOT NULL,
    player_email    VARCHAR(255) UNIQUE NOT NULL,
    account_balance INT DEFAULT 0 -- Store balance in cents
);

CREATE TABLE IF NOT EXISTS tournaments
(
    tournament_id   INT PRIMARY KEY AUTO_INCREMENT,
    tournament_name VARCHAR(255)   NOT NULL,
    prize_pool      INT NOT NULL, -- Store prize pool in cents
    start_date      DATETIME       NOT NULL,
    end_date        DATETIME       NOT NULL
);

CREATE TABLE IF NOT EXISTS bets
(
    bet_id        INT PRIMARY KEY AUTO_INCREMENT,
    player_id     INT            NOT NULL,
    amount        INT NOT NULL,  -- Store bet amount in cents
    bet_time      DATETIME       NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players (player_id)
);

CREATE INDEX idx_player_email ON players (player_email);
CREATE INDEX idx_bets_player_id ON bets (player_id);
CREATE INDEX idx_bets_player_tournament_time ON bets (player_id, bet_time);


DELIMITER //

CREATE PROCEDURE distribute_prizes()
BEGIN
    DECLARE active_tournament_id INT;
    DECLARE active_tournament_prize_pool INT;
    DECLARE active_tournament_start_date DATETIME;
    DECLARE active_tournament_end_date DATETIME;

SELECT
    tournament_id, prize_pool, start_date, end_date
INTO active_tournament_id, active_tournament_prize_pool, active_tournament_start_date, active_tournament_end_date
FROM tournaments
WHERE NOW() BETWEEN start_date AND end_date;

UPDATE players p
    JOIN (
    SELECT
    player_id,
    RANK() OVER (ORDER BY SUM(amount) DESC) AS player_rank
    FROM bets
    WHERE bet_time BETWEEN active_tournament_start_date AND active_tournament_end_date
    GROUP BY player_id
    ) ranked_bets ON p.player_id = ranked_bets.player_id
    SET p.account_balance = CASE
        WHEN ranked_bets.player_rank = 1 THEN p.account_balance + (active_tournament_prize_pool * 0.5)
        WHEN ranked_bets.player_rank = 2 THEN p.account_balance + (active_tournament_prize_pool * 0.3)
        WHEN ranked_bets.player_rank = 3 THEN p.account_balance + (active_tournament_prize_pool * 0.2)
        ELSE p.account_balance
END
WHERE ranked_bets.player_rank <= 3;

END //

DELIMITER ;