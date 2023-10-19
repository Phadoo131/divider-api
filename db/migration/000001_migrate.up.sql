CREATE TABLE "Client" (
  "ClientID" serial PRIMARY KEY,
  "Name" VARCHAR(255),
  "Email" VARCHAR(255),
  "Phone_Number" VARCHAR(20),
  "Password" VARCHAR(20),
  "Address" TEXT
);

CREATE TABLE "Booker" (
  "BookerID" serial PRIMARY KEY,
  "Name" VARCHAR(255),
  "Email" VARCHAR(255),
  "Phone_Number" VARCHAR(20),
  "Password" VARCHAR(20),
  "Address" TEXT
);

CREATE TABLE "Restaurant" (
  "RestaurantID" serial PRIMARY KEY,
  "Name" VARCHAR(255),
  "Location" VARCHAR(255),
  "Contact_Number" VARCHAR(20)
);

CREATE TABLE "Booking" (
  "BookingID" serial PRIMARY KEY,
  "ClientID" INT REFERENCES "Client" ("ClientID"),
  "BookerID" INT REFERENCES "Booker" ("BookerID"),
  "RestaurantID" INT REFERENCES "Restaurant" ("RestaurantID"),
  "Date_and_Time" timestamptz,
  "Status" VARCHAR(20)
);

CREATE TABLE "ChatMessage" (
  "MessageID" serial PRIMARY KEY,
  "SenderID" INT REFERENCES "Client" ("ClientID"),
  "ReceiverID" INT REFERENCES "Client" ("ClientID"),
  "MessageText" TEXT,
  "Timestamp" timestamptz
);