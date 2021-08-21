CREATE TABLE IF NOT EXISTS money (
   id   	    SERIAL        PRIMARY KEY,
   user_id      INT           NOT NULL, 
   date_time    DATE          NOT NULL,
   amount       INT           NOT NULL,
   is_expense   BOOLEAN       DEFAULT false NOT NULL,
   title        VARCHAR(100)  NOT NULL,
   description  VARCHAR(150),
   CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);