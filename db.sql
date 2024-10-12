CREATE DATABASE IF NOT EXISTS auth_service;
CREATE DATABASE IF NOT EXISTS document_service;
CREATE DATABASE IF NOT EXISTS timetable_service;


\c auth_service
CREATE TYPE IF NOT EXISTS account_role AS ENUM ('Admin', 'Manager', 'Doctor', 'User');

CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(30) NOT NULL CHECK(last_name != ''),
    first_name VARCHAR(30) NOT NULL CHECK(first_name != ''),
    username VARCHAR(30) NOT NULL UNIQUE CHECK(username != ''),
    password_hash VARCHAR(100) NOT NULL,
    roles account_role[] NOT NULL DEFAULT '{"User"}'
);

CREATE INDEX IF NOT EXISTS idx_username ON Account(username);

\c document_service
CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    pacient_id INTEGER NOT NULL,
    timetable_id INTEGER NOT NULL,
    time TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS history (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    pacient_id INTEGER NOT NULL,
    hospital_id INTEGER NOT NULL,
    doctor_id INTEGER NOT NULL,
    room VARCHAR(50) NOT NULL,
    data TEXT
);

\c timetable_service
CREATE TABLE IF NOT EXISTS timetables (
    id SERIAL PRIMARY KEY,
    hospital_id INTEGER NOT NULL,
    doctor_id INTEGER NOT NULL,
    from_time TIMESTAMP NOT NULL,
    to_time TIMESTAMP NOT NULL,
    room VARCHAR(50) NOT NULL
);

\c hospital_service

