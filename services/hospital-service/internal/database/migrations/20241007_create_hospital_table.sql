CREATE TABLE IF NOT EXISTS hospitals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL CHECK(name != ''),
    address VARCHAR(150) NOT NULL UNIQUE CHECK(address != ''),
    contact_phone VARCHAR(20) NOT NULL UNIQUE CHECK(contact_phone != ''),
    rooms VARCHAR(50)[] NOT NULL
);