	CREATE TABLE IF NOT EXISTS tb_casestudy (
        id SERIAL PRIMARY KEY,
        title VARCHAR, 
        description VARCHAR,
        imageuri VARCHAR,  
        createddate TIMESTAMP
        );

    INSERT INTO tb_casestudy (title, description, imageuri, createddate) 
                    VALUES ('test_title', 'test_description', 'test_imageuri', '2004-10-19 10:23:54');
        