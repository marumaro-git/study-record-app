SET CHARSET UTF8;
SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;

CREATE TABLE IF NOT EXISTS record (
        id INT AUTO_INCREMENT PRIMARY KEY,
        study_type VARCHAR(10) NOT NULL,
        item VARCHAR(50) NOT NULL,
        content VARCHAR(255) NOT NULL,
        study_date DATE NOT NULL,
        create_datetime DATETIME DEFAULT CURRENT_TIMESTAMP,
        update_datetime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO record (study_type, item, content, study_date, create_datetime, update_datetime)
VALUES
    ('IT', 'golang', 'API開発しました', '2024-01-01', NOW(), NOW()),
    ('ENGLISH', 'Biology', 'あいさつについて', '2022-01-02', NOW(), NOW()),
    ('IT', 'java', 'Causes and consequences', '2022-01-03', NOW(), NOW());