# SQL

```
CREATE TABLE guestbook (guestName VARCHAR(255) NOT NULL, content VARCHAR(255) NOT NULL, date TIMESTAMP NOT NULL, entryID SERIAL PRIMARY KEY);
INSERT INTO guestbook (guestName, content, date) values ('first guest', 'I got here!', '2017-08-06 12:00:00');
INSERT INTO guestbook (guestName, content, date) values ('second guest', 'Me too!', '2017-08-06 13:00:00');
```