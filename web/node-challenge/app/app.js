const fs = require('fs')
const app = require('express')()
const server = require('http').createServer(app)
const io = require('socket.io')(server)

app.get('/hello', (req, res) => {
  res.type('txt').send('Hello, expressjs works!')
})

app.use((req, res, next) => {
  res.status(404)
  res.type('txt').send('Expressjs: Not found')
})

io.on('connection', socket => {
  console.log('user connected:', socket.id)
  socket.emit('message', 'socket.io works!')
  socket.on('message', d => {
    console.log('message:', d)
  })
})

const unixSocket = '/home/run/.unix.sock'
if (unixSocket && fs.existsSync(unixSocket)) fs.unlinkSync(unixSocket)
server.listen(unixSocket || process.env.PORT || 80, () => {
  if (unixSocket) fs.chmodSync(unixSocket, 755)
  console.log('Server running at ' + (unixSocket || process.env.PORT || 80))
})
