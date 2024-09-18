

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
CREATE INDEX idx_username ON users(username);



CREATE TABLE departments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    parent_id INT DEFAULT NULL,
    flags TINYINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES departments(id)
);
CREATE INDEX idx_department_name ON departments(name);


-- CreateDepartment procedure
DELIMITER $$
CREATE PROCEDURE CreateDepartment(IN dept_name VARCHAR(255), IN parent_id INT, IN dept_flags INT)
BEGIN
    INSERT INTO departments (name, parent_id, flags)
    VALUES (dept_name, parent_id, dept_flags);
END$$

DELIMITER ;


-- UpdateDepartment procedure
DELIMITER $$
CREATE PROCEDURE UpdateDepartment(IN dept_id INT, IN dept_name VARCHAR(255), IN parent_id INT, IN flags TINYINT)
BEGIN
    UPDATE departments
    SET name = dept_name, parent_id = parent_id, flags = flags, updated_at = CURRENT_TIMESTAMP
    WHERE id = dept_id;
END$$

DELIMITER ;



-- DeleteDepartment procedure
DELIMITER $$
CREATE PROCEDURE DeleteDepartment(IN p_id INT)
BEGIN
    UPDATE departments
    SET flags = flags & ~1 | 2
    WHERE id = p_id;
END$$

DELIMITER ;

-- GetDepartmentByName procedure
DELIMITER $$
CREATE PROCEDURE GetDepartmentByID(IN dept_id INT)
BEGIN
    SELECT *
    FROM departments
    WHERE id = dept_id;
END$$

DELIMITER ;


-- GetDepartmentHierarchy procedure
DELIMITER $$
CREATE PROCEDURE GetDepartmentHierarchy(IN departmentId INT)
BEGIN
    WITH RECURSIVE DepartmentHierarchy AS (

        SELECT 
            d.id,
            d.name,
            d.parent_id,
            d.flags,
            d.created_at,
            d.updated_at
        FROM 
            departments d
        WHERE 
            d.id = departmentId
        
        UNION ALL
        
        SELECT 
            dp.id,
            dp.name,
            dp.parent_id,
            dp.flags,
            dp.created_at,
            dp.updated_at
        FROM 
            departments dp
        INNER JOIN DepartmentHierarchy dh ON dp.id = dh.parent_id
    )
    -- Select the entire hierarchy
    SELECT 
        id,
        name,
        parent_id,
        flags,
        created_at,
        updated_at
    FROM 
        DepartmentHierarchy;
END$$

DELIMITER ;

-- GetAllDepartments procedure
DELIMITER $$
CREATE PROCEDURE GetAllDepartments()
BEGIN
    SELECT *
    FROM departments;
END$$
DELIMITER ;


-- CreateUser procedure
DELIMITER $$
CREATE PROCEDURE CreateUser(IN username VARCHAR(255), IN password VARCHAR(255))
BEGIN
    INSERT INTO users (username, password)
    VALUES (username, password);
END$$

DELIMITER ;


-- GetUserByUsername procedure
DELIMITER $$
CREATE PROCEDURE GetUserByUsername(IN username VARCHAR(255))
BEGIN
    SELECT *
    FROM users
    WHERE username = username;
END$$

DELIMITER ;
