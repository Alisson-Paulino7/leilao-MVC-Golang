CREATE TABLE IF NOT EXISTS usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    senha VARCHAR(255) NOT NULL,
    cpf VARCHAR(15) NOT NULL,
    telefone VARCHAR(15) NOT NULL
);

CREATE TABLE IF NOT EXISTS produtos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome_product VARCHAR(100) NOT NULL,
    desc_product VARCHAR(255) NOT NULL,
    lance_inicial INT(10) NOT NULL,
    tmp_expiracao INT NOT NULL,
    -- data_cadastro DATE NOT NULL,
    -- data_expires DATE NOT NULL,
    lance_atual INT(10) NOT NULL DEFAULT 0,
    foto_product MEDIUMBLOB
);