-- Estrutura para tabela `carrinho`

CREATE TABLE carrinho (
  id SERIAL PRIMARY KEY,
  product_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  id_vendedor INTEGER,
  quantidade INTEGER NOT NULL
);

-- Estrutura para tabela `produto`

CREATE TABLE produto (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  nome VARCHAR(100) NOT NULL,
  descricao TEXT,
  imagem TEXT,
  preco NUMERIC(10,2) NOT NULL,
  quantidade INTEGER DEFAULT 0,
  codigo VARCHAR(255),
  garantia INTEGER,
  categoria VARCHAR(255),
  marca VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

-- Estrutura para tabela `usuario`

CREATE TABLE usuario (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  senha VARCHAR(255) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  tipo VARCHAR(255) NOT NULL,
  tel VARCHAR(20),
  endereco TEXT,
  cpf VARCHAR(30) UNIQUE,
  cep VARCHAR(45),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

-- Estrutura para tabela `vendas`

CREATE TABLE vendas (
  id SERIAL PRIMARY KEY,
  vendedor_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  product_id INTEGER NOT NULL,
  total NUMERIC(10,2) NOT NULL,
  quantidade INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  endereco TEXT,
  num_residencia INTEGER NOT NULL,
  cpf VARCHAR(30),
  cep VARCHAR(45),
  mtd_pay VARCHAR(45),
  sts_venda VARCHAR(45) DEFAULT 'Confirmada'
);

-- Restrições para tabela `carrinho`

ALTER TABLE carrinho
  ADD CONSTRAINT fk_client FOREIGN KEY (user_id) REFERENCES usuario (id),
  ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES produto (id),
  ADD CONSTRAINT fk_vendedor FOREIGN KEY (id_vendedor) REFERENCES usuario (id);

-- Restrições para tabela `produto`

ALTER TABLE produto
  ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES usuario (id) ON DELETE CASCADE ON UPDATE CASCADE;

-- Restrições para tabela `vendas`

ALTER TABLE vendas
  ADD CONSTRAINT fk_client_vendas FOREIGN KEY (user_id) REFERENCES usuario (id) ON UPDATE CASCADE,
  ADD CONSTRAINT fk_vendedor_vendas FOREIGN KEY (vendedor_id) REFERENCES usuario (id) ON UPDATE CASCADE,
  ADD CONSTRAINT fk_product_vendas FOREIGN KEY (product_id) REFERENCES produto (id) ON UPDATE CASCADE;
