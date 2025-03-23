Để tạo một người dùng mới trên hệ thống Linux, bạn có thể sử dụng lệnh `useradd`. Dưới đây là các bước và lệnh cụ thể để tạo một người dùng mới:

1. **Sử dụng lệnh `useradd`**: Lệnh `useradd` tạo một người dùng mới. Sử dụng tùy chọn `-m` để tạo thư mục home cho người dùng mới.

2. **Thiết lập mật khẩu cho người dùng mới**: Sau khi tạo người dùng mới, sử dụng lệnh `passwd` để thiết lập mật khẩu.

Dưới đây là các bước cụ thể với ví dụ tạo người dùng mới tên là `newuser`:

### Bước 1: Tạo người dùng mới
```bash
sudo useradd -m newuser
```

### Bước 2: Thiết lập mật khẩu cho người dùng mới
```bash
sudo passwd newuser
```
### Bước 3: (Tùy chọn) Thêm người dùng vào các nhóm
Bạn có thể thêm người dùng vào các nhóm cụ thể. Ví dụ, để thêm `newuser` vào nhóm `sudo` để họ có quyền quản trị, sử dụng lệnh sau:
```bash
sudo usermod -aG sudo newuser
```

### Kiểm tra người dùng mới
```bash
id newuser
```

3. **Gnome**
```bash
paru -S gnome-browser-connector
```
### If completely unable to play ANY videos.
```bash
systemctl --user restart pipewire wireplumber
```
