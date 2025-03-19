## paru

Paru là một công cụ quản lý gói cho Arch Linux, dựa trên yay, và nó hỗ trợ cả AUR (Arch User Repository). Dưới đây là một số câu lệnh thường được sử dụng với Paru:

1. **Cài đặt một gói từ Arch Linux Repositories:**
   ```bash
   paru -S <package_name>
   ```

2. **Cài đặt một gói từ AUR:**
   ```bash
   paru -S <aur_package_name>
   ```

3. **Cập nhật tất cả các gói:**
   ```bash
   paru -Syu
   ```

4. **Tìm kiếm gói trong AUR:**
   ```bash
   paru -Ss <keyword>
   ```

5. **Xóa một gói:**
   ```bash
   paru -Rns <package_name>
   ```
   Lưu ý rằng `-ns` sẽ xóa cả tệp cấu hình và các phụ thuộc không sử dụng.

6. **Cài đặt một gói mà không cần xác nhận (non-interactive):**
   ```bash
   paru -S --noconfirm <package_name>
   ```

7. **Cập nhật chỉ một số cụ thể của gói (không cập nhật toàn bộ hệ thống):**
   ```bash
   paru -U <package_name>
   ```

8. **Hiển thị thông tin về gói:**
   ```bash
   paru -Qi <package_name>
   ```

9. **Hiển thị thông tin về một gói từ AUR:**
   ```bash
   paru -Si <aur_package_name>
   ```

10. **Kiểm tra xem có cập nhật nào không:**
   ```bash
   paru -Qu
   ```

11. **Xóa tất cả các gói không cần thiết (orphans):**
   ```bash
   paru -Rns $(paru -Qtdq)
   ```

12. **Kiểm tra tất cả các gói đã cài đặt:**
   ```bash
   paru -Q
   ```

13. **Update Mirrors**
    ```bash
    sudo reflector --verbose --sort rate -l 20 --save /etc/pacman.d/mirrorlist
    ```

Nhớ kiểm tra tài liệu của Paru hoặc nhập `paru --help` để xem danh sách đầy đủ các tùy chọn và lựa chọn.
