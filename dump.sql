CREATE TABLE IF NOT EXISTS tasks (
    id integer PRIMARY KEY,
    description text NOT NULL,
    completed integer NOT NULL DEFAULT 0
);
