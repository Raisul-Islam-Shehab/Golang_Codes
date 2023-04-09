DROP TABLE IF EXISTS manager;
CREATE TABLE manager (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    blood_group VARCHAR(55) NOT NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO manager (name, phone, email, blood_group)
VALUES ('Shehab', '123456', 'abc@gmail.com', 'B+'),
    ('Adib', '234567', 'bcz@gmail.com', 'A+'),
    ('Rahat', '923456', 'abt@gmail.com', 'O-'),
    ('Raju', '323456', 'ghf@gmail.com', 'O+');