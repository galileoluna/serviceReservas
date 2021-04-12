DROP DATABASE IF EXISTS reservas;
CREATE DATABASE reservas;
GO

/*
        DATA MODEL SERVICE SORTEO
*/
DROP TABLE if EXISTS Reserva;
CREATE TABLE Reserva (
  ID BIGINT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  ID_Evento BIGINT NOT NULL,
  ID_Cliente BIGINT NOT NULL,
  Nombre_Cliente NVARCHAR(50),
  Monto FLOAT
);
INSERT INTO Reserva(ID_Evento,ID_Cliente,Nombre_Cliente,Monto) VALUES (1,1,"Galileo Luna",10.5);
GO

select * from Reserva;


