-- Cria o banco de dados caso ele não exista
CREATE
    DATABASE IF NOT EXISTS finplanning;

-- Usar os próximos comandos dentro do banco finplanning
USE
finplanning;

-- Se tiver uma tabela item dentro do banco finplanning, vou dropar ela
DROP TABLE IF EXISTS item;

-- Criando a tabela de informações
CREATE TABLE itens
(
    id                 int auto_increment primary key,
    item               varchar(50) not null,
    valor              int         not null,
    quantidadeDeParcelas int         not null,
    beneficiario       varchar(50) not null,
    compradoEm         timestamp default current_timestamp()
)ENGINE=INNODB;