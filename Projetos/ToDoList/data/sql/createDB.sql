-- criando banco de dados
CREATE TABLE tarefas ( ID INT AUTO_INCREMENT PRIMARY KEY,conteudo VARCHAR(100),data_termino DATE,status_conclusao BOOLEAN);
-- pegar todos os dados de tarefa
SELECT * FROM tarefas
-- pegar dados especificos de tarefas
SELECT t.conteudo,t.data_termino,t.status_conclusao from tarefas t