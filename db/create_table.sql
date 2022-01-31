CREATE TABLE rooms(
  id              VARCHAR(26)             NOT NULL,
  name	          VARCHAR(50)             NOT NULL,
  created_at      DATETIME                NOT NULL,
  updated_at      DATETIME                NOT NULL,
  PRIMARY KEY(id)
)ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE TABLE messages(
  id              VARCHAR(26)             NOT NULL,
  room_id		  INT					  NOT NULL,
  message         VARCHAR(100)            NOT NULL,
  created_at      DATETIME                NOT NULL,
  updated_at      DATETIME                NOT NULL,
  PRIMARY KEY(id)
)ENGINE=INNODB DEFAULT CHARSET=utf8;
