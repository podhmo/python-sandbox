INSERT INTO groups (id, name) VALUES(NULL, 'X');
INSERT INTO groups (id, name) VALUES(NULL, 'Y');
INSERT INTO groups (id, name) VALUES(NULL, 'Z');
INSERT INTO users (id, group_id, name) VALUES(NULL, 1, 'foo');
INSERT INTO users (id, group_id, name) VALUES(NULL, 1, 'bar');
INSERT INTO users (id, group_id, name) VALUES(NULL, 2, 'boo');
INSERT INTO skills (id, user_id, name) VALUES(NULL, 1, 'a');
INSERT INTO skills (id, user_id, name) VALUES(NULL, 1, 'b');
INSERT INTO skills (id, user_id, name) VALUES(NULL, 1, 'c');
INSERT INTO skills (id, user_id, name) VALUES(NULL, 2, 'a');
INSERT INTO skills (id, user_id, name) VALUES(NULL, 2, 'b');
INSERT INTO skills (id, user_id, name) VALUES(NULL, 2, 'c');
INSERT INTO skills (id, user_id, name) VALUES(NULL, 3, 'a');
INSERT INTO skills (id, user_id, name) VALUES(NULL, 3, 'b');
INSERT INTO skills (id, user_id, name) VALUES(NULL, 3, 'c');