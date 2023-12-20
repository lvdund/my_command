Để xóa một gói (package) trên Arch Linux, bạn có thể sử dụng lệnh `pacman -R` hoặc `pacman -Rs`. Dưới đây là cách sử dụng các lệnh này:

1. **Xóa gói không cần giữ các tệp cấu hình (Remove):**
    ```bash
    sudo pacman -R <package_name>
    ```
    Thay `<package_name>` bằng tên của gói bạn muốn xóa.

2. **Xóa gói và các phụ thuộc không cần thiết khác (Remove and clean up dependencies):**
    ```bash
    sudo pacman -Rs <package_name>
    ```
    Lệnh này sẽ xóa gói cùng với các phụ thuộc không cần thiết khác mà không ảnh hưởng đến các phụ thuộc mà bạn đã cài đặt cho các gói khác.

3. **Xóa gói và tất cả các phụ thuộc không sử dụng (Remove and clean up all dependencies):**
    ```bash
    sudo pacman -Rns <package_name>
    ```
    Lệnh này xóa gói và tất cả các phụ thuộc không sử dụng, giải phóng không gian trên ổ đĩa.

4. **Xóa gói và tất cả các phụ thuộc không sử dụng:**
    ```bash
    yay -Rns $(yay -Qdtq)

    sudo pacman -Rns $(pacman -Qdtq)
    ```

Lưu ý rằng khi bạn xóa một gói, nó có thể ảnh hưởng đến các ứng dụng khác mà có thể phụ thuộc vào nó. Hãy kiểm tra cẩn thận trước khi xóa bất kỳ gói nào.

Nếu bạn muốn xóa các tệp cấu hình của gói cùng với nó, bạn có thể sử dụng lệnh `pacman -Rns`.
