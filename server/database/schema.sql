-- Create table
CREATE TABLE IF NOT EXISTS exchange (
    id INTEGER NOT NULL PRIMARY KEY,
    code TEXT NOT NULL,
    codein TEXT NOT NULL,
    name TEXT NOT NULL,
    high TEXT NOT NULL,
    low TEXT NOT NULL,
    varBid TEXT NOT NULL,
    pctChange TEXT NOT NULL,
    bid TEXT NOT NULL,
    ask TEXT NOT NULL,
    timestamp TEXT NOT NULL,
    create_date TEXT NOT NULL
);
