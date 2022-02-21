CREATE TABLE snippets(
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(100) NOT NULL,
  context TEXT NOT NULL,
  created DATETIME NOT NULL,
  expires DATETIME NOT NULL
);

CREATE INDEX idx_snippetx_created ON snippets(created)

CREATE TABLE users(
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  hashed_password CHAR(60) NOT NULL,
  created DATETIME NOT NULL
);


INSERT INTO users(name,email,hashed_password,created) VALUES(
  'Indiana Jones',
  'indiana@example.com',
  '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
  '2022-02-21 08:45:00'
);




