Arch Linux là một hệ điều hành Linux dựa trên việc xây dựng từ đầu (from scratch) và thiết kế để cung cấp sự linh hoạt và kiểm soát cao đối với hệ thống của người dùng. Dưới đây là một số lệnh quan trọng mà bạn có thể cần biết khi sử dụng Arch Linux:

1. **Pacman:**
   - `pacman -Syu`: Cập nhật toàn bộ hệ thống.
   - `pacman -S <package>`: Cài đặt gói phần mềm.
   - `pacman -R <package>`: Gỡ cài đặt một gói.
   - `pacman -Ss <keyword>`: Tìm kiếm gói phần mềm.
   - `pacman -Q <package>`: Hiển thị phiên bản của một gói đã được cài đặt.

2. **Systemctl:**
   - `systemctl start <service>`: Khởi động một dịch vụ.
   - `systemctl stop <service>`: Dừng một dịch vụ.
   - `systemctl restart <service>`: Khởi động lại một dịch vụ.
   - `systemctl enable <service>`: Kích hoạt một dịch vụ để chạy khi khởi động hệ thống.
   - `systemctl disable <service>`: Tắt kích hoạt tự động cho một dịch vụ.

3. **Journalctl:**
   - `journalctl`: Hiển thị nhật ký hệ thống.
   - `journalctl -xe`: Hiển thị nhật ký với thông báo lỗi gần đây.

4. **Pacstrap:**
   - `pacstrap /mnt base`: Cài đặt hệ điều hành Arch Linux trên một hệ thống mới.

5. **mkinitcpio:**
   - `mkinitcpio -p linux`: Tạo initramfs cho kernel Linux.

6. **Networkctl:**
   - `networkctl`: Hiển thị trạng thái mạng.

7. **User Management:**
   - `useradd -m -g users -G wheel,user -s /bin/bash <username>`: Tạo một người dùng mới.
   - `passwd <username>`: Đặt mật khẩu cho người dùng.

8. **File System:**
   - `ls`: Liệt kê các tập tin và thư mục.
   - `cd <directory>`: Di chuyển đến thư mục.
   - `cp <source> <destination>`: Sao chép tập tin hoặc thư mục.
   - `mv <source> <destination>`: Di chuyển hoặc đổi tên tập tin hoặc thư mục.
   - `rm <file>`: Xóa một tập tin.
   - `mkdir <directory>`: Tạo một thư mục mới.

9. **Arch Wiki:**
   - `man <command>`: Hiển thị hướng dẫn cho một lệnh cụ thể.
   - `arch-wiki`: Truy cập Wiki của Arch Linux trên trình duyệt web.

Lưu ý rằng đối với nhiều lệnh, bạn có thể cần quyền root hoặc sử dụng sudo để thực hiện các tác vụ quản lý hệ thống. Để biết thêm chi tiết và tùy chọn, bạn nên tham khảo tài liệu và hướng dẫn chính thức của Arch Linux.