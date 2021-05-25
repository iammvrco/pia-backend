
create database Manhattam
use Manhattam


create table Users(
	UserName varchar(30) PRIMARY KEY not null,
	Password varchar(50) not null
)


create table PersonalInfo( 
	CURP varchar(19) primary key not null,
	RFC varchar(14),
	NSS varchar(12),
	Names varchar(30) not null,
	LastNames varchar(30) not null,
	BirthDate date not null
);


create table Phones(
    PhoneNumber varchar(11) not null,
	Place varchar(20) not null,
	CURP varchar(19) not null
		foreign key (CURP)
		references PersonalInfo (CURP)
		on delete cascade 
);


create table Emails(
	Email varchar(100) primary key not null,
	CURP varchar(19) not null
		foreign key (CURP)
		references PersonalInfo (CURP)
		on delete cascade 
);


create table Addresses(
	ID int PRIMARY KEY IDENTITY(1,1) not null,
	Name varchar(30) not null,
	Street varchar(30) not null,
	Number varchar(10) not null,
	Neighborhood varchar(30) not null,
	ZipCode varchar(10) not null,
	City varchar(30) not null,
	State varchar(30) not null,
	Country varchar(30) not null,
	CURP varchar(19) not null
		foreign key (CURP)
		references PersonalInfo (CURP)
		on delete cascade 
);

INSERT INTO Users values ('marco','marco123')


INSERT INTO PersonalInfo values ('123', '456', '789', 'Marco Antonio', 'Picapiedra', '1998-09-08')
INSERT INTO PersonalInfo values ('234', '456', '789', 'Judith', 'Picapiedra', '1998-09-08')
INSERT INTO PersonalInfo values ('345', '456', '789', 'Gretelsilla', 'Picapiedra', '1998-09-08')
INSERT INTO PersonalInfo values ('456', '456', '789', 'Aldayr', 'Picapiedra', '1998-09-08')
INSERT INTO PersonalInfo values ('567', '456', '789', 'Adal', 'Picapiedra', '1998-09-08')
INSERT INTO PersonalInfo values ('678', '456', '789', 'Johana', 'Picapiedra', '2007-06-24')
INSERT INTO PersonalInfo values ('789', '456', '789', 'Alejandra', 'Picapiedra', '1999-09-08')
INSERT INTO PersonalInfo values ('891', '456', '789', 'Teresa', 'Picapiedra', '1998-09-08')
INSERT INTO PersonalInfo values ('910', '456', '789', 'Rosy', 'Picapiedra', '1998-09-08')
INSERT INTO PersonalInfo values ('101', '456', '789', 'Abigail', 'Picapiedra', '1998-09-08')


INSERT INTO Phones values ('85154651','Casa','123')
INSERT INTO Phones values ('54164616','Oficina','123')
INSERT INTO Phones values ('87651515','Oficina','234')
INSERT INTO Phones values ('54164616','Casa','101')
INSERT INTO Phones values ('64684184','Oficina','345')
INSERT INTO Phones values ('9878515','Casa','891')
INSERT INTO Phones values ('87963245','Oficina','678')
INSERT INTO Phones values ('9884545','Casa','678')
INSERT INTO Phones values ('7896255','Casa','910')
INSERT INTO Phones values ('658418','Casa','789')


INSERT INTO Emails values ('123@hotmail.com','123')
INSERT INTO Emails values ('123@gmail.com','123')
INSERT INTO Emails values ('234@gmail.com','234')
INSERT INTO Emails values ('345@outlook.com','345')
INSERT INTO Emails values ('345@gmail.com','345')
INSERT INTO Emails values ('345correo3@gmail.com','345')
INSERT INTO Emails values ('567@gmail.com','567')
INSERT INTO Emails values ('678@gmail.com','678')
INSERT INTO Emails values ('891@gmail.com','891')
INSERT INTO Emails values ('101@hotmail.com','101')
INSERT INTO Emails values ('101@gmail.com','101')
INSERT INTO Emails values ('910@gmail.com','910')


INSERT INTO Addresses values ('Casa','Estrella','1208','Luna','66478','Marte','Nuevo León','México','123')
INSERT INTO Addresses values ('Casa','Planeta','45','Celeste','6647','Jupiter','Nuevo León','México','234')
INSERT INTO Addresses values ('Casa','Anillo','120','Rosa','678','Saturno','Nuevo León','México','234')
INSERT INTO Addresses values ('Casa','Santos','1208','Angeles','664','Urano','Nuevo León','México','345')
INSERT INTO Addresses values ('Casa','Mora','1208','Sara','478','Neptuno','Nuevo León','México','123')
INSERT INTO Addresses values ('Casa','Felicidad','569','Tristeza','1010','Mercurio','Nuevo León','México','101')
INSERT INTO Addresses values ('Casa','Fuentes','1408','Luces','987','Venus','Nuevo León','México','891')
INSERT INTO Addresses values ('Casa','Oso','700','Peluche','478','Tierra','Nuevo León','México','678')




