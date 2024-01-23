const express = require('express')
const mysql = require('mysql2')
const falso = require('@ngneat/falso')

const app = express()
const port = 3000
const config = {
    host: 'db',
    user: 'root',
    password: 'root',
    database: 'nodedb'
};

const connection = mysql.createConnection(config)

app.get('/', (req, res) => {
    const name = falso.randFullName()
    connection.query(`INSERT INTO people (nome) VALUES ('${name}')`)
  
    connection.query(`SELECT nome FROM people`, (error, results, fields) => {
      res.send(`
        <h1>Full Cycle Rocks!</h1>
        <ol>
          ${!!results.length ? results.map(el => `<li>${el.nome}</li>`).join('') : ''}
        </ol>
      `)
    })
  })

app.listen(port, () => {
    console.log('Rodando na porta ' + port)
})