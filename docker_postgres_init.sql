CREATE TABLE patient(  
    id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255),
    message VARCHAR(255),
    order_id INT
);

INSERT INTO patient (name, message, order_id)
VALUES ('小人', 'Patient is experiencing chest pain', 1),
       ('小民', '超過120請施打8u', 2),
       ('小小', 'Patient has a follow-up appointment next week', 3),
       ('小異', 'Patient needs a refill on their medication', 4),
       ('小爆', 'Patient has an allergy to penicillin', 5);
