CREATE TABLE recommendations (
  listKey  VARCHAR(64) NOT NULL,
  entryKey VARCHAR(64) NOT NULL,
  PRIMARY KEY (listKey, entryKey)
);

CREATE TABLE entryList (
  listKey  VARCHAR(64) NOT NULL,
  AID  VARCHAR(64)[] ,
  PRIMARY KEY (listKey)
);

