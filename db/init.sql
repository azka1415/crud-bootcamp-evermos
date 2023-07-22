-- Active: 1690029910907@@127.0.0.1@3306
--  Execute this first if using sqlite
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS materials (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			teacher_id INTEGER NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
			updated_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
			FOREIGN KEY (teacher_id) REFERENCES teachers (id)
);

CREATE TABLE IF NOT EXISTS teachers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			position TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
			updated_at TIMESTAMP NOT NULL DEFAULT (datetime('now'))
);


