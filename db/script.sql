CREATE DATABASE IF NOT EXISTS igaming_platform;

USE igaming_platform;

CREATE TABLE players
(
    player_id       INT PRIMARY KEY AUTO_INCREMENT,
    player_name     VARCHAR(255)        NOT NULL,
    player_email    VARCHAR(255) UNIQUE NOT NULL,
    account_balance DECIMAL(10, 2) DEFAULT 0.00
);

CREATE TABLE tournaments
(
    tournament_id   INT PRIMARY KEY AUTO_INCREMENT,
    tournament_name VARCHAR(255)   NOT NULL,
    prize_pool      DECIMAL(10, 2) NOT NULL,
    start_date      DATETIME       NOT NULL,
    end_date        DATETIME       NOT NULL
);

CREATE TABLE bets
(
    bet_id        INT PRIMARY KEY AUTO_INCREMENT,
    player_id     INT            NOT NULL,
    tournament_id INT            NOT NULL,
    amount        DECIMAL(10, 2) NOT NULL,
    bet_time      DATETIME       NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players (player_id),
    FOREIGN KEY (tournament_id) REFERENCES tournaments (tournament_id)
);

CREATE INDEX idx_player_email ON players (player_email);
CREATE INDEX idx_bets_player_id ON bets (player_id);
CREATE INDEX idx_bets_tournament_id ON bets (tournament_id);
CREATE INDEX idx_bets_player_tournament_time ON bets (player_id, tournament_id, bet_time);