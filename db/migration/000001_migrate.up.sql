CREATE TABLE "client" (
  "id" serial PRIMARY KEY,
  "name_client" VARCHAR(255),
  "email" VARCHAR(255),
  "phone_number" VARCHAR(20),
  "pwd" VARCHAR(20),
  "address_client" TEXT,
  "created_at" timestamptz
);

CREATE TABLE "booker" (
  "id" serial PRIMARY KEY,
  "name_booker" VARCHAR(255),
  "email" VARCHAR(255),
  "phone_number" VARCHAR(20),
  "pwd" VARCHAR(20),
  "address_booker" TEXT,
  "created_at" timestamptz
);

CREATE TABLE "restaurant" (
  "id" serial PRIMARY KEY,
  "name_restaurant" VARCHAR(255),
  "location_restaurant" VARCHAR(255),
  "contact_number" VARCHAR(20),
  "created_at" timestamptz
);

CREATE TABLE "booking" (
  "id" serial PRIMARY KEY,
  "client_id" INT REFERENCES "client" ("id"),
  "booker_id" INT REFERENCES "booker" ("id"),
  "restaurant_id" INT REFERENCES "restaurant" ("id"),
  "date_and_time" timestamptz,
  "status_booking" VARCHAR(20)
);

CREATE TABLE "chat_message" (
  "id" serial PRIMARY KEY,
  "sender_id" INT REFERENCES "client" ("id"),
  "receiver_id" INT REFERENCES "client" ("id"),
  "message_text" TEXT,
  "timestamp" timestamptz
);