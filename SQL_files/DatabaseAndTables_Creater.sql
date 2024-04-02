-- To drop table if made wrong
DROP TABLE IF EXISTS "Clients";
CREATE TABLE IF NOT EXISTS "Clients" (
	"id" serial NOT NULL UNIQUE,
	"Owner_Name" varchar(256) NOT NULL,
	"Email" varchar(256) NOT NULL,
	"Phone_Number" int NOT NULL,
	"Dog_Name" varchar(256) NOT NULL,
	"Dog_Breed" varchar(256),
	"Dog_Age" int NOT NULL,
	PRIMARY KEY ("id")
);
-- Test records for table "Clients"
INSERT INTO "Clients" ("id", "Owner_Name", "Email", "Phone_Number", "Dog_Name", "Dog_Breed", "Dog_Age") VALUES (1, 'John Doe', 'JohnDoe@example.com', 1234567890, 'Fido', 'Golden Retriever', 5);
INSERT INTO "Clients" ("id", "Owner_Name", "Email", "Phone_Number", "Dog_Name", "Dog_Breed", "Dog_Age") VALUES (2, 'Jane Doe', 'JaneDoe@example.com', 1234567890, 'Fido', 'Blood Hound', 5);
INSERT INTO "Clients" ("id", "Owner_Name", "Email", "Phone_Number", "Dog_Name", "Dog_Breed", "Dog_Age") VALUES (3, 'John Smith', 'JohnSmith@example.com', 1234567890, 'Fido', 'Retriever', 5);

-- To drop table if made wrong
DROP TABLE IF EXISTS "Appointment";

CREATE TABLE IF NOT EXISTS "Appointment" (
	"id" serial NOT NULL UNIQUE,
	"Client_id" int NOT NULL,
	"appointment_date" date NOT NULL,
	"appointment_Time" time NOT NULL UNIQUE,
	"Service" int NOT NULL,
	PRIMARY KEY ("id")
);
-- Test records for table "Appointment"
INSERT INTO "Appointment" ("id", "Client_id", "appointment_date", "appointment_Time", "Service") VALUES (1, 1, '2022-01-01', '12:00:00', 1);
INSERT INTO "Appointment" ("id", "Client_id", "appointment_date", "appointment_Time", "Service") VALUES (2, 2, '2022-01-01', '12:30:00', 2);
INSERT INTO "Appointment" ("id", "Client_id", "appointment_date", "appointment_Time", "Service") VALUES (3, 3, '2022-01-01', '1:00:00', 3);
-- To drop table if made wrong
DROP TABLE IF EXISTS "Service";
CREATE TABLE IF NOT EXISTS "Service" (
	"id" serial NOT NULL UNIQUE,
	"service_name" varchar(256) NOT NULL,
	"service_description" text NOT NULL,
	"price" decimal(10) NOT NULL DEFAULT '2',
	PRIMARY KEY ("id")
);
-- Test records for table "Service"
INSERT INTO "Service" ("id", "service_name", "service_description", "price") VALUES (1, 'Bath', 'Bath and blow dry', 20);
INSERT INTO "Service" ("id", "service_name", "service_description", "price") VALUES (2, 'Haircut', 'Haircut and blow dry', 30);
INSERT INTO "Service" ("id", "service_name", "service_description", "price") VALUES (3, 'Nail Trim', 'Nail trim and file', 10);
INSERT INTO "Service" ("id", "service_name", "service_description", "price") VALUES (4, 'Full Grooming', 'This is a Full Grooming. It includes Bath, Haircut, and Nail Trim', 50);


-- To drop table if made wrong
DROP TABLE IF EXISTS "Service_Provider";
CREATE TABLE IF NOT EXISTS "Service_Provider" (
	"id" serial NOT NULL UNIQUE,
	"name" VARCHAR(256) NOT NULL,
	"job" VARCHAR(256) NOT NULL,
	"service_offered" int NOT NULL,
	PRIMARY KEY ("id")
);
-- Test records for table "service_provider"
INSERT INTO "Service_Provider" ("id", "name", "job", "service_offered") VALUES (1, 'Aramis', 'Groomer', 1);

-- Foreign Key Constraints
ALTER TABLE "Appointment" ADD CONSTRAINT "Appointment_fk1" FOREIGN KEY ("Client_id") REFERENCES "Clients"("id");
ALTER TABLE "Appointment" ADD CONSTRAINT "Appointment_fk4" FOREIGN KEY ("Service") REFERENCES "Service"("id");
ALTER TABLE "service_provider" ADD CONSTRAINT "service_provider_fk3" FOREIGN KEY ("service_offered") REFERENCES "Service"("id");

