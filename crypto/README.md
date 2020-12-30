# ctf-deploy
How to deploy challenges for a CTF event?
Làm thế nào để triển khai các challenges crypto cho một giải CTF?

# crypto challenges
## Vấn đề
- Qua một số lần deploy crypto challenges, mình thấy vấn đề chính của công việc này là tạo một proxy để map từ stdin/stdout <=> network, việc scan challenges cũng không quan trọng lắm. Bởi vì không giống như web, crypto đa số cần source và cũng bật container là chạy luôn, không cần phải test nhiều.
- Do đó, mình để public ports các challenges này và khi nào mở bài thi thì mới bật container

## Giải quyết
- Để giải quyết vấn đề về stdin/stdout <=> network, mình đã viết một proxy có sẵn [main.go](../proxy-cmd/main.go)
- Vậy thôi, rất đơn giản

# Chạy
Đây là 3 bài crypto của tác giả (NDH)[https://github.com/nguyenduyhieukma] được dùng trong KMACTF-2020
- `./affine/run.sh`
- `./nonsense/run.sh`

# Truy cập
- `nc -v 127.0.0.1 9991`
- `nc -v 127.0.0.1 9992`
