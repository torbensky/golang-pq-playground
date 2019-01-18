CREATE TABLE users(
    id int NOT NULL,
    name text
);

INSERT INTO users (id,name) VALUES
    (1,'Sam'),
    (2,'Joe'),
    (3,'Robin');

CREATE TABLE user_things(
    thing_id int NOT NULL,
    user_id int NOT NULL,
    type text,
    description text
);

INSERT INTO user_things (thing_id, user_id, type, description) VALUES
    (1, 1, 'Mouse', 'Some super awesome mouse stuff for sure'),
    (2, 1, 'Laptop', 'A mint laptop of great value'),
    (3, 2, 'Hat', 'Keeps your head warm'),
    (4, 2, 'Backpack', 'Most useful thing in the universe'),
    (5, 3, 'Phone', 'A big waste of money');