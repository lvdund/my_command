## yay

Yay là một công cụ quản lý gói cho Arch Linux và hỗ trợ cả AUR (Arch User Repository). Dưới đây là một số lệnh thường được sử dụng với yay:

1. **Cài đặt một gói từ Arch Linux Repositories:**
   ```bash
   yay -S <package_name>
   ```

2. **Cài đặt một gói từ AUR:**
   ```bash
   yay -S <aur_package_name>
   ```

3. **Cập nhật tất cả các gói:**
   ```bash
   yay -Syu
   ```

4. **Tìm kiếm gói trong AUR:**
   ```bash
   yay -Ss <keyword>
   ```

5. **Xóa một gói:**
   ```bash
   yay -Rns <package_name>
   ```
   Lưu ý rằng `-ns` sẽ xóa cả tệp cấu hình và các phụ thuộc không sử dụng.

6. **Cài đặt một gói mà không cần xác nhận (non-interactive):**
   ```bash
   yay -S --noconfirm <package_name>
   ```

7. **Cập nhật chỉ một số cụ thể của gói (không cập nhật toàn bộ hệ thống):**
   ```bash
   yay -U <package_name>
   ```

8. **Hiển thị thông tin về gói:**
   ```bash
   yay -Qi <package_name>
   ```

9. **Hiển thị thông tin về một gói từ AUR:**
   ```bash
   yay -Si <aur_package_name>
   ```

10. **Kiểm tra xem có cập nhật nào không:**
    ```bash
    yay -Qu
    ```

11. **Xóa tất cả các gói không cần thiết (orphans):**
    ```bash
    yay -Rns $(yay -Qtdq)
    ```

12. **Kiểm tra tất cả các gói đã cài đặt:**
    ```bash
    yay -Q
    ```

Nhớ kiểm tra tài liệu của yay hoặc nhập `yay --help` để xem danh sách đầy đủ các tùy chọn và lựa chọn.