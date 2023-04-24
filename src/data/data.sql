CREATE TABLE ListQuestion (
    idListQuestion INT NOT NULL,
    question VARCHAR(255) UNIQUE,
    answer VARCHAR(255) UNIQUE,
    PRIMARY KEY (idListQuestion)
);

CREATE TABLE AskHistory(
    idHistory INT NOT NULL,
    idQuestion INT,
    idSession INT NOT NULL,
    question VARCHAR(255) NOT NULL,
    answer VARCHAR(255) NOT NULL,
    PRIMARY KEY (idHistory),
    FOREIGN KEY (idQuestion) REFERENCES ListQuestion(idListQuestion)
);

INSERT INTO ListQuestion (idListQuestion, question, answer) values (1, 'Matkul wajib paling seru?', 'udah pasti STIMA');
INSERT INTO ListQuestion (idListQuestion, question, answer) values (2, 'Matkul wajib paling susah?', 'udah pasti OS');
INSERT INTO ListQuestion (idListQuestion, question, answer) values (3, 'Matkul wajib paling mudah?', 'udah pasti PBO');
INSERT INTO ListQuestion (idListQuestion, question, answer) values (4, 'Anak K3 paling imba?', 'ZULLLL');
INSERT INTO ListQuestion (idListQuestion, question, answer) values (5, 'Siapa dosen K3 STIMA?', 'Pak Rila');
INSERT INTO ListQuestion (idListQuestion, question, answer) values (6, 'ITB ada dimana aja?', 'Bandung, Jatinangor, Cirebon');
INSERT INTO ListQuestion (idListQuestion, question, answer) values (7, 'Informatika berjiwa apa?', 'Informatika berjiwa satria');

INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (1, NULL, 1, '2+3', '5');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (2, NULL, 1, 'Sekarang tanggal berapa?', '24/02/2023');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (3, 6, 2, 'ITB ada dimana aja?', 'Bandung, Jatinangor, Cirebon');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (4, 6, 3, 'ITB ada dimana aja?', 'Bandung, Jatinangor, Cirebon');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (5, 5, 3, 'Siapa dosen K3 STIMA?', 'Pak Rila');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (6, 5, 3, 'Siapa dosen K3 STIMA?', 'Pak Rila');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (7, 2, 4, 'Matkul wajib paling susah?', 'udah pasti OS');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (8, 1, 4, 'Matkul wajib paling seru?', 'udah pasti STIMA');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (9, 3, 4, 'Matkul wajib paling mudah?', 'udah pasti PBO');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (10, 4, 5, 'Anak K3 paling imba?', 'ZULLLL');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (11, 4, 5, 'Anak K3 paling imba?', 'ZULLLL');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (12, 4, 5, 'Anak K3 paling imba?', 'ZULLLL');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (13, 7, 4, 'Informatika berjiwa apa?', 'Informatika berjiwa satria');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (14, NULL, 1, '10^2', '100');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (15, NULL, 1, '5*8+3', '43');
INSERT INTO AskHistory(idHistory, idQuestion, idSession, question, answer) values (16, NULL, 1, 'Tanggal 12 April 2023 itu hari apa?', 'Hari Rabu');