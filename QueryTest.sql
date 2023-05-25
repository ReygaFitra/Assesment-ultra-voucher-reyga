-- buat query tabelnya terlebih dahulu
CREATE TABLE parent_table (
  id INT,
  name VARCHAR,
  parent_id INT,
  PRIMARY KEY (id),
  FOREIGN KEY (parent_id) REFERENCES parent_table(id)
);

-- masukan data sesuai soal
INSERT INTO parent_table (id, name, parent_id)
VALUES
  (1, 'Zaki', 2),
  (2, 'Ilham', NULL),
  (3, 'Irwan', 2),
  (4, 'Arka', 3);

-- Query untuk menghasilkan data sesuai soal yang diberikan
SELECT t1.id, t1.name, t2.name AS parent_name
FROM parent_table AS t1
LEFT JOIN parent_table AS t2 ON t1.parent_id = t2.id;