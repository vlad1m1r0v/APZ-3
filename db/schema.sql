DROP TABLE IF EXISTS machines CASCADE;
CREATE TABLE machines (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE,
    cpuCount SMALLINT,
    totalDiskSpace  INTEGER
);

DROP TABLE IF EXISTS disks CASCADE;
CREATE TABLE  disks (
    id SERIAL PRIMARY KEY,
    machines_id INTEGER,
    FOREIGN KEY (machines_id) REFERENCES machines(id),
    name VARCHAR(20) NOT NULL UNIQUE,
    capacity INTEGER
);

