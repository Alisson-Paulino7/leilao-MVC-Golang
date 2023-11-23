CREATE TABLE IF NOT EXISTS usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    senha VARCHAR(20) NOT NULL,
    cpf VARCHAR(15) NOT NULL,
    telefone VARCHAR(15) NOT NULL
);

CREATE TABLE IF NOT EXISTS produtos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome_product VARCHAR(100) NOT NULL,
    desc_product VARCHAR(255) NOT NULL,
    lance_inicial VARCHAR(10) NOT NULL,
    tmp_expiracao INT NOT NULL,
    foto_product BLOB
);