CREATE TABLE banks (
  id  INT NOT NULL  AUTO_INCREMENT  PRIMARY KEY,
  name varchar(255) NOT NULL,
  code varchar(255) NOT NULL,
  ussd varchar(255) DEFAULT NULL,
  logo varchar(255) DEFAULT NULL,
  slug varchar(255) NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
);
