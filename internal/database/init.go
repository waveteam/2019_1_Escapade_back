package database

import (
	"database/sql"
	"fmt"
	"os"

	//
	_ "github.com/lib/pq"
)

// Init try to connect to DataBase.
// If success - return instance of DataBase
// if failed - return error
func Init() (db *DataBase, err error) {
	//connStr := "user=unemuzhregdywt password=5d9ae1059a39b0a8838b5653854adc7fb266deb7da1dc35de729a4836ba27b65 dbname=dd1f3dqgsuq1k5 sslmode=disable"

	//connStr := "user=rolepade password=escapade dbname=escabase sslmode=disable"

	//var database *sql.DB
	//database, err = sql.Open("postgres", connStr)
	var database *sql.DB
	database, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		return
	}

	db = &DataBase{Db: database}
	db.Db.SetMaxOpenConns(20)

	err = db.Db.Ping()
	if err != nil {
		return
	}

	if err = db.CreateTables(); err != nil {
		return
	}

	return
}

func (db *DataBase) CreateTables() error {
	sqlStatement := `
	DROP TABLE IF EXISTS Player;
	DROP TABLE IF EXISTS Photo;
	DROP TABLE IF EXISTS Session;
	DROP TABLE IF EXISTS Game;

	CREATE TABLE Photo (
    id    SERIAL PRIMARY KEY,--integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    title varchar(30) NOT NULL
);

CREATE TABLE Player (
    id SERIAL PRIMARY KEY,--integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(30) NOT NULL,
    password varchar(30) NOT NULL,
    email varchar(30) NOT NULL,
    photo_id int,
    best_score int,
    best_time int,
    FOREIGN KEY (photo_id) REFERENCES Photo (id)
);

CREATE Table Session (
    id SERIAL PRIMARY KEY,
    player_id int NOT NULL,
    session_code varchar(30) NOT NULL,
    expiration timestamp without time zone NOT NULL
    --FOREIGN KEY (player_id) REFERENCES Player (id)
);

ALTER TABLE Session
ADD CONSTRAINT session_player
   FOREIGN KEY (player_id)
   REFERENCES Player(id)
   ON DELETE CASCADE;

CREATE Table Game (
    id SERIAL PRIMARY KEY,
    player_id int NOT NULL,
    FieldWidth int CHECK (FieldWidth > -1),
    FieldHeight int CHECK (FieldHeight > -1),
    MinsTotal int CHECK (MinsTotal > -1),
    MinsFound int CHECK (MinsFound > -1),
    Finished bool NOT NULL,
    Exploded bool NOT NULL,
    Date timestamp without time zone NOT NULL,
    FOREIGN KEY (player_id) REFERENCES Player (id)
);

/*
CREATE Table PlayerStatistics (
    id SERIAL PRIMARY KEY,
    player_id int NOT NULL,
    GamesTotal  int CHECK (MinsTotal > -1),
	SingleTotal int CHECK (MinsTotal > -1),
	OnlineTotal int CHECK (MinsTotal > -1),
	SingleWin   int CHECK (SingleWin > -1),
	OnlineWin   int CHECK (OnlineWin > -1),
	MinsFound   int CHECK (MinsFound > -1),
	FirstSeen   timestamp without time zone NOT NULL,
	LastSeen    timestamp without time zone NOT NULL,
    FOREIGN KEY (player_id) REFERENCES Player (id)
);
*/

INSERT INTO Player(name, password, email, best_score, best_time) VALUES
    ('tiger', 'Bananas', 'tinan@mail.ru', 1000, 10),
    ('panda', 'apple', 'today@mail.ru', 2323, 20),
    ('catmate', 'juice', 'allday@mail.ru', 10000, 5),
    ('hotdog', 'where', 'three@mail.ru', 88, 1000);

    /*
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(30) NOT NULL,
    password varchar(30) NOT NULL,
    email varchar(30) NOT NULL,
    photo_id int,
    best_score int,
    FOREIGN KEY (photo_id) REFERENCES Photo (id)
    */

INSERT INTO Game(player_id, FieldWidth, FieldHeight,
MinsTotal, MinsFound, Finished, Exploded, Date) VALUES
    (1, 50, 50, 100, 20, true, true, date '2001-09-28'),
    (1, 50, 50, 80, 30, false, false, date '2018-09-27'),
    (1, 50, 50, 70, 70, true, false, date '2018-09-26'),
    (1, 50, 50, 60, 30, true, true, date '2018-09-23'),
    (1, 50, 50, 50, 50, true, false, date '2018-09-24'),
    (1, 50, 50, 40, 30, true, false, date '2018-09-25'),
    (2, 25, 25, 80, 30, false, false, date '2018-08-27'),
    (2, 25, 25, 70, 70, true, false, date '2018-08-26'),
    (2, 25, 25, 60, 30, true, true, date '2018-08-23'),
    (3, 10, 10, 10, 10, true, false, date '2018-10-26'),
    (3, 10, 10, 20, 19, true, true, date '2018-10-23'),
    (3, 10, 10, 30, 30, true, false, date '2018-10-24'),
    (3, 10, 10, 40, 5, true, false, date '2018-10-25');

    /*
CREATE Table Game (
    id SERIAL PRIMARY KEY,
    player_id int NOT NULL,
    FieldWidth int CHECK (FieldWidth > -1),
    FieldHeight int CHECK (FieldHeight > -1),
    MinsTotal int CHECK (MinsTotal > -1),
    MinsFound int CHECK (MinsFound > -1),
    Finished bool NOT NULL,
    Exploded bool NOT NULL,
    Date timestamp without time zone NOT NULL,
    FOREIGN KEY (player_id) REFERENCES Player (id)
);
    */
	`
	_, err := db.Db.Exec(sqlStatement)

	if err != nil {
		fmt.Println("database/init - fail:" + err.Error())

	}
	return err
}
