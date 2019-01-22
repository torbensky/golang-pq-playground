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
    description text,
    map jsonb DEFAULT '{}'::jsonb NOT NULL
);

INSERT INTO user_things (thing_id, user_id, type, description, map) VALUES
    (1, 1, 'Mouse', 'Some super awesome mouse stuff for sure', '{}'::jsonb),
    (2, 1, 'Laptop', 'A mint laptop of great value', '{}'::jsonb),
    (3, 2, 'Hat', 'Keeps your head warm', '{}'::jsonb),
    (4, 2, 'Backpack', 'Most useful thing in the universe','{"test": "foo"}'::jsonb),
    (5, 3, 'Phone', 'A big waste of money','{}'::jsonb);
