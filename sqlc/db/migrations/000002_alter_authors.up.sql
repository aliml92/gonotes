TRUNCATE TABLE authors;
ALTER TABLE authors
DROP COLUMN name,
ALTER COLUMN bio SET NOT NULL,
ADD COLUMN birth_year INT NOT NULL;