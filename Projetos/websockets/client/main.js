const WebSocket = require('ws');

// Criando a conexão WebSocket
const ws = new WebSocket('ws://localhost:3330');

// Definindo o manipulador para quando a conexão estiver aberta
ws.on('open', () => {
    console.log('Conexão aberta');
    // Enviando uma mensagem
    ws.send('Olá mundo!');
});

// Definindo o manipulador para quando uma mensagem for recebida
ws.on('message', (data) => {
    console.log('Mensagem recebida: ' + data);
});

// Definindo o manipulador para quando ocorrer um erro
ws.on('error', (error) => {
    console.error('Erro na conexão WebSocket: ', error);
});

// Definindo o manipulador para quando a conexão for fechada
ws.on('close', (code, reason) => {
    console.log('Conexão fechada: ', code, reason);
});
