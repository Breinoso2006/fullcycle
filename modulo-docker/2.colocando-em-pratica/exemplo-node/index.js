const express = require('express')
const app = express()
const port = 3000

// O req e res do Express, são os objetos de requisição e resposta em uma função de middleware do Express.
app.get('/', (req,res) => {
    res.send('<h1>Olá! Testando aqui o node hehehe </h1>')
})

app.listen(port, ()=> {
    console.log('Rodando na porta ' + port)
})