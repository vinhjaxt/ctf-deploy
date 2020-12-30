## Website sử dụng ngôn ngữ PHP
- Mỗi project sẽ được định nghĩa bởi 1 thư mục trong /home/public_html/.
- Ví dụ: /home/public_html/project1
         /home/public_html/project2
- Để truy cập vào project1: https://username-project1.luôn.vn
- Để truy cập vào project2: https://username-project2.luôn.vn

## Website sử dụng ngôn ngữ khác (NodeJS, Go,..)
- Mỗi project sẽ được để ở 1 thư mục trong /home/run/.
- Ví dụ: /home/run/nodejs-project1
         /home/run/go-project2
- Các chương trình của project1 phải listen unix socket file với đường dẫn:
         /home/run/.project1-unix.sock
- Các chương trình của project2 phải listen unix socket file với đường dẫn:
         /home/run/.project2-unix.sock
- Để truy cập vào project1: https://unix-username-project1.luôn.vn
- Để truy cập vào project2: https://unix-username-project2.luôn.vn
### Demo nodejs
```js
const fs = require('fs')
const unixSocket = '/home/run/.project1-unix.sock'
const http = require('http')
const server = http.createServer((req, res) => {
  res.statusCode = 200
  res.setHeader('Content-Type', 'text/plain')
  res.end('Hello world! Nodejs works!\n')
})
if (unixSocket && fs.existsSync(unixSocket)) fs.unlinkSync(unixSocket)
server.listen(unixSocket || process.env.PORT || 80, () => {
  if (unixSocket) fs.chmodSync(unixSocket, 755)
  console.log('Server running at ' + (unixSocket || process.env.PORT || 80))
})
```

## Kết nối cơ sở dữ liệu
- MySQL:
  - Unix socket path: /var/run/mysqld/mysqld.sock
  - Hostname: localhost
  - User: username
  - Password: password
- pgsql:
  - Unix socket path: /var/run/pgsql/.s.PGSQL.5432
  - User: username
  - Password: password

## Quản lý cơ sở dữ liệu
- phpMyAdmin: https://dev-pma.luôn.vn/ (Đăng nhập với user, password MySQL ở trên)
- pgAdmin: https://dev-pga.luôn.vn/ (Đăng nhập với user, password pgsql ở trên)

## Lưu trữ data
- Tất cả data sẽ được lưu trữ ở đường dẫn /home/www-data/.
- Nếu lưu trữ ở một đường dẫn khác, data có thể bị mất nếu restart server (hoặc có sự cố sảy ra)
- Nên lưu trữ data theo từng project, ví dụ như:
    - /home/www-data/project1
    - /home/www-data/project2
