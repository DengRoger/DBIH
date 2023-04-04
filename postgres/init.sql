CREATE TABLE "UIDHead" (
  "UID" CHAR(64) PRIMARY KEY,
  "HeadKey" CHAR(64)
);

CREATE TABLE "Link" (
  "LID" CHAR(64) PRIMARY KEY,
  "listKey" CHAR(64),
  "NextKey" CHAR(64)
);

CREATE TABLE "List" (
  "LID" CHAR(64) PRIMARY KEY,
  "AIDList" CHAR(640)
  "NextKey" CHAR(64)
);

CREATE TABLE "AID" (
  "AID" CHAR(64) PRIMARY KEY,
  "Topic" TEXT,
  "Article" TEXT
);

