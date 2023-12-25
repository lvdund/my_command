Để merge một nhánh vào nhánh hiện tại (ví dụ: `main`), bạn có thể sử dụng lệnh `git merge`. Dưới đây là hướng dẫn cơ bản:

### Merge thông thường:

1. **Chuyển đến nhánh mà bạn muốn merge vào:**
   ```bash
   git checkout main
   ```

2. **Thực hiện merge:**
   ```bash
   git merge <other-branch>
   ```
   Hoặc nếu bạn đang ở trên nhánh mà bạn muốn merge vào, bạn có thể chỉ cần gõ:
   ```bash
   git merge <other-branch>
   ```

3. **Giải quyết xung đột (nếu có):**
   Nếu có xung đột giữa các commit, Git sẽ dừng ở đây và yêu cầu bạn giải quyết xung đột trước khi tiếp tục. Sử dụng `git status` và `git diff` để xác định vị trí của xung đột, và sau đó chỉnh sửa các tệp để giải quyết xung đột. Sau khi giải quyết xung đột, thêm tệp đã sửa đổi vào staging area (`git add`) và tiếp tục quá trình merge (`git merge --continue`).

4. **Commit merge:**
   Sau khi giải quyết xung đột, bạn cần tạo một commit để hoàn thành quá trình merge:
   ```bash
   git commit -m "Merge <other-branch> into main"
   ```

### Merge với lịch sử tuyến tính (Fast-forward):

Nếu không có xung đột và lịch sử commit của nhánh cần merge là một phần của lịch sử tuyến tính của nhánh hiện tại, thì quá trình merge sẽ là fast-forward, tức là chỉ cần di chuyển con trỏ nhánh.

```bash
git checkout main
git merge <other-branch>
```

Lệnh trên có thể được thực hiện mà không cần commit mới.

Lưu ý: Trong trường hợp sử dụng fast-forward, không có commit merge mới được tạo ra, và lịch sử commit giống nhau giữa các nhánh.

Xử lý xung đột trong quá trình merge là một quá trình quan trọng để đảm bảo tính nhất quán của mã nguồn trong Git. Dưới đây là các bước chi tiết để xử lý xung đột khi bạn đang thực hiện merge:

## Xử lý Conflict

#### 1. Thực Hiện Merge:

Đầu tiên, bạn cần thực hiện lệnh merge. Đảm bảo bạn đã checkout vào nhánh mà bạn muốn merge vào (thường là `main` hoặc `master`). Sử dụng lệnh sau để thực hiện merge:

```bash
git checkout main  # hoặc git checkout master
git merge <other-branch>
```

#### 2. Phát Hiện Xung Đột:

Sau khi thực hiện lệnh merge, Git có thể phát hiện xung đột. Sử dụng lệnh sau để kiểm tra trạng thái xung đột:

```bash
git status
```

#### 3. Mở Trình Soạn Thảo để Giải Quyết Xung Đột:

Sử dụng một trình soạn thảo để mở các tệp tin có xung đột. Git sẽ đánh dấu nơi xảy ra xung đột trong tệp.

```bash
git mergetool
```

#### 4. Giải Quyết Xung Đột:

Trong trình soạn thảo, bạn sẽ thấy các phần được đánh dấu là `<<<<<<<`, `=======`, và `>>>>>>>`. Các phần giữa giữa `<<<<<<<` và `=======` là từ nhánh hiện tại, và giữa `=======` và `>>>>>>>` là từ nhánh bạn đang merge vào. Bạn cần quyết định làm thế nào để giải quyết xung đột.

Chỉnh sửa các phần xung đột, sau đó lưu và đóng trình soạn thảo.

#### 5. Đánh Dấu Xung Đột Đã Giải Quyết:

Sử dụng lệnh sau để đánh dấu rằng xung đột đã được giải quyết:

```bash
git add <file-path>
```

#### 6. Hoàn Thành Quá Trình Merge:

Tiếp tục quá trình merge:

```bash
git merge --continue
```

#### 7. Commit Kết Quả Merge:

Sử dụng lệnh commit để lưu kết quả của quá trình merge:

```bash
git commit -m "Merge <other-branch> into main"
```

#### 8. Kết Thúc Quá Trình Merge:

Kiểm tra xem có công việc merge nào còn lại không:

```bash
git status
```

Nếu không có xung đột nào khác, quá trình merge đã hoàn thành.

### Lưu Ý:

- Nếu bạn muốn hủy bỏ quá trình merge, bạn có thể sử dụng lệnh `git merge --abort`.
- Luôn kiểm tra kết quả sau mỗi bước để đảm bảo tính nhất quán của mã nguồn.