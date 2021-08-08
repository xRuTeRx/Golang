CREATE TABLE "contacts" (
    "id" bigserial primary key,
    "name" VARCHAR not null,
	"phone" VARCHAR not null,
    "group_id" INT not null
);

CREATE TABLE "groups" (
    "id" bigserial primary key,
    "name" VARCHAR not null
);

ALTER TABLE "contacts" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

INSERT INTO "groups" ("name") VALUES ('inginers');
INSERT INTO "groups" ("name") VALUES ('boss');
INSERT INTO "groups" ("name") VALUES ('workers');