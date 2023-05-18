

--PostgreSQL
CREATE TABLE Autos (
AutoID INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY(START WITH 1 INCREMENT BY 1),
Brand VARCHAR(50) not null,
Model VARCHAR(50) not null,
Year INT not null,
BodyType VARCHAR(50) not nul,
EngineType VARCHAR(50) not null,
Displacement INT,
Price DECIMAL(10,2) not null
);


--Test insert 
INSERT INTO Autos (Brand, Model, Year, BodyType, EngineType, Displacement, Price)
VALUES ('Toyota', 'Corolla', 2022, 'Sedan', 'Gasoline', 1600, 22000.00);