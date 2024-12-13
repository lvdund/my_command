Để thiết lập cấu hình toàn cầu cho Git, bạn có thể sử dụng lệnh `git config --global`. Đây là một số tùy chọn cấu hình phổ biến mà bạn có thể bắt đầu cài đặt:

1. **Init:**
   ```bash
   git config --global user.name "lvdund"
   git config --global user.email "lvdund@gmail.com"
   ssh-keygen -t rsa -C "lvdund@gmail.com"
   ```
   
2. **Mặc định trình soạn thảo (chẳng hạn là VSCode, Sublime Text, Nano, Vim):**
   ```bash
   git config --global core.editor "code --wait"  # Đối với VSCode
   ```

   Lưu ý: Bạn có thể thay thế `"code --wait"` bằng lệnh khác tương ứng với trình soạn thảo bạn muốn sử dụng.

3. **Định dạng mặc định cho lịch sử commit (có thể dùng để xem lịch sử đẹp hơn):**
   ```bash
   git config --global log.decorate auto
   ```

4. **Hiển thị tên nhánh hiện tại trong dấu nhắc command line:**
   ```bash
   git config --global bash.showCurrentBranch true
   ```

5. **Cấu hình mặc định cho fast-forward khi merge:**
   ```bash
   git config --global merge.ff only
   ```

6. **Hiển thị màu sắc trong output của Git:**
   ```bash
   git config --global color.ui auto
   ```

7. **Hiển thị tên branch trong prompt command line:**
   ```bash
   git config --global bash.showCurrentBranch true
   ```

8. **Thiết lập công cụ diff mặc định (chẳng hạn, sử dụng `meld`):**
   ```bash
   git config --global diff.tool meld
   ```

Lưu ý rằng bạn có thể sửa đổi các tùy chọn này bất kỳ lúc nào bằng cách chạy lại lệnh `git config --global` với giá trị mới.

