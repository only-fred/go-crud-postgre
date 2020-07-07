CREATEDB db_user;

CREATE TABLE tbl_user(
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(50) NOT NULL
)

INSERT INTO tbl_user (username, password)
  VALUES ('T_601', 'Yojimbo');