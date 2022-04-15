DROP TABLE IF EXISTS links;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS faskes;
DROP TABLE IF EXISTS kota_kabupaten;
DROP TABLE IF EXISTS provinsi;
DROP TABLE IF EXISTS official_email;
DROP TABLE IF EXISTS jwt;
DROP TABLE IF EXISTS register_confirmation;

CREATE TABLE IF NOT EXISTS users (
    ID SERIAL UNIQUE,
    username VARCHAR (127) NOT NULL UNIQUE,
    password VARCHAR (127) NOT NULL
);

CREATE TABLE IF NOT EXISTS links (
    ID SERIAL UNIQUE,
    title VARCHAR (255),
    address VARCHAR (255),
    UserID INT,
    FOREIGN KEY (UserID) REFERENCES Users(ID)
);


CREATE TABLE IF NOT EXISTS provinsi (
    id_provinsi VARCHAR (255) NOT NULL UNIQUE,
    nama VARCHAR (255),
    PRIMARY KEY (id_provinsi), 
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS kota_kabupaten (
    id_kota_kabupaten VARCHAR (255) NOT NULL UNIQUE,
    nama VARCHAR (255),
    id_provinsi VARCHAR(255),
    FOREIGN KEY (id_provinsi) REFERENCES Provinsi(id_provinsi),
    PRIMARY KEY (id_kota_kabupaten),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS faskes (
    id_faskes VARCHAR(255) NOT NULL UNIQUE,
    nama_faskes VARCHAR(255),
    jenis_faskes VARCHAR(255),
    total_nakes INT,
    id_kota_kabupaten VARCHAR(255),
    FOREIGN KEY (id_kota_kabupaten) REFERENCES Kota_Kabupaten(id_kota_kabupaten),
    PRIMARY KEY (id_faskes),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS official_email (
    email VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (email),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS jwt (
    signed_token VARCHAR(1000) NOT NULL UNIQUE,
    is_logout BOOLEAN,
    PRIMARY KEY (signed_token),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS register_confirmation (
    email VARCHAR (255) NOT NULL UNIQUE,
    hashed_password VARCHAR (255),
    registration_token VARCHAR (1000),
    PRIMARY KEY (email), 
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

