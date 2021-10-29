CREATE TABLE IF NOT EXISTS persona (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'Id de la persona',
  	nombre VARCHAR(100) COMMENT 'Nombre de la persona',
 	apellidos VARCHAR(100) COMMENT 'Apellidos del persona',
  	PRIMARY KEY (id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='Tabla de Personas';