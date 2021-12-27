USE db_global;

DROP PROCEDURE IF EXISTS add_id;
CREATE PROCEDURE add_id(count INT)

BEGIN
	WHILE count>0 DO
		INSERT INTO id_pool(invalid) VALUE(0);
		SET count=count-1;
	END WHILE;
	
END;
	