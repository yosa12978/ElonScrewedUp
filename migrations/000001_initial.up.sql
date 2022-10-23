CREATE TABLE IF NOT EXISTS Tweets (
    Id TEXT primary key,
    Text TEXT,
    Url TEXT not null,
    Date BIGINT not null
);