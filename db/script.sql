CREATE TABLE produtos
(
    id         serial primary key,
    nome       varchar,
    descricao  varchar,
    preco      decimal,
    quantidade integer
)

INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Miniatura de Ferrari da Fórmula 1', 'Miniatura detalhada do carro de corrida da Fórmula 1 da Ferrari de 2007', 29.99, 50);
INSERT INTO produtos (nome, descricao, preco, quantidade)
VALUES ('Boné da Ferrari', 'Boné oficial da equipe Ferrari de Fórmula 1, com logotipo bordado', 19.99, 100);
INSERT INTO produtos (nome, descricao, preco, quantidade)
VALUES ('Poster autografado pelo piloto de Fórmula 1 Lewis Hamilton',
        'Poster de alta qualidade autografado por Lewis Hamilton', 49.99, 20);
INSERT INTO produtos (nome, descricao, preco, quantidade)
VALUES ('Camiseta de Fórmula 1', 'Camiseta de algodão com estampa exclusiva inspirada na Fórmula 1', 24.99, 80);

INSERT INTO produtos (nome, descricao, preco, quantidade)
VALUES ('Luvas de sandá profissional', 'Luvas de sandá de couro genuíno, ideais para treinamento e competição', 99.99,
        30);
INSERT INTO produtos (nome, descricao, preco, quantidade)
VALUES ('Uniforme de Kung Fu', 'Uniforme de alta qualidade para praticantes de Kung Fu, feito com material resistente',
        129.99, 25);
INSERT INTO produtos (nome, descricao, preco, quantidade)
VALUES ('Protetor bucal para Sandá',
        'Protetor bucal moldável para praticantes de Sandá, garantindo segurança durante os combates', 14.99, 50);
INSERT INTO produtos (nome, descricao, preco, quantidade)
VALUES ('DVD de filme do Bruce Lee', 'DVD de filme antigo do Bruce Lee com técnicas avançadas de Kung Fu', 19.99, 15);