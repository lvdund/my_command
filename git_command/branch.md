## Branch

### Create branch

Để tạo một nhánh mới và đẩy nó lên kho lưu trữ từ xa (remote repository), bạn có thể thực hiện các bước sau:

1. **Tạo Nhánh Mới:**
   ```bash
   git checkout -b tên_nhánh_mới
   ```

   Hoặc sử dụng cú pháp ngắn gọn:
   ```bash
   git branch tên_nhánh_mới
   git checkout tên_nhánh_mới
   ```

2. **Tạo Nhánh Mới Copy Từ 1 Nhánh:**

   Để tạo một nhánh mới và sao chép nó từ một nhánh hiện tại, bạn có thể sử dụng lệnh `git checkout` với tùy chọn `-b`. Dưới đây là cách để thực hiện điều này:

   ```bash
   git checkout -b tên_nhánh_mới tên_nhánh_nguồn
   ```

   Lệnh trên có ý nghĩa là bạn đang tạo một nhánh mới với tên `tên_nhánh_mới` và sao chép toàn bộ lịch sử và thay đổi từ nhánh `tên_nhánh_nguồn`.

   Ví dụ, nếu bạn muốn tạo một nhánh mới có tên là `feature-branch` và sao chép từ nhánh `main`, bạn sẽ sử dụng lệnh sau:

   ```bash
   git checkout -b feature-branch main
   ```

   Sau khi tạo nhánh mới, bạn có thể thực hiện các thay đổi, commit, và đẩy nhánh mới lên kho lưu trữ từ xa (nếu cần thiết), sử dụng các lệnh `git add`, `git commit`, và `git push`.

3. **Thực Hiện Các Thay Đổi và Commit:**
   Thực hiện các thay đổi trên nhánh mới và thực hiện commit:

   ```bash
   git add .
   git commit -m "Mô tả commit"
   ```

4. **Đẩy Nhánh Lên Kho Lưu Trữ Từ Xa (Remote):**
   ```bash
   git push origin tên_nhánh_mới
   ```

   Đối với trường hợp đầu tiên, bạn có thể cần thêm `-u` để thiết lập nhánh từ xa theo dõi (tracking branch) để sau này có thể sử dụng `git push` và `git pull` mà không cần chỉ định tên nhánh:

   ```bash
   git push -u origin tên_nhánh_mới
   ```

   Sau khi thiết lập, bạn có thể đơn giản sử dụng `git push`:

   ```bash
   git push
   ```

Lưu ý rằng bạn cần có quyền ghi (write) đối với nhánh trên kho lưu trữ từ xa để thực hiện thao tác đẩy (push).


### List branch

Để liệt kê tất cả các nhánh (branches) trong một kho lưu trữ Git, bạn có thể sử dụng lệnh `git branch`. Dưới đây là cách để thực hiện điều này:

1. **Liệt kê tất cả các nhánh:**
   ```bash
   git branch
   ```

   Lệnh trên sẽ hiển thị danh sách tất cả các nhánh trong kho lưu trữ. Nhánh hiện tại sẽ được đánh dấu bằng dấu `*`.

2. **Liệt kê cả nhánh từ xa (remote branches):**
   ```bash
   git branch -a
   ```

   Lệnh trên sẽ hiển thị danh sách tất cả các nhánh cả từ xa (remote branches) trong kho lưu trữ.

3. **Liệt kê nhánh hiện tại và thể hiện commit gần đây:**
   ```bash
   git branch -v
   ```

   Lệnh trên sẽ hiển thị danh sách các nhánh cùng với commit gần đây mà mỗi nhánh đang trỏ tới.

Nếu bạn muốn xem tất cả các nhánh (bao gồm cả những nhánh ẩn), bạn có thể sử dụng tùy chọn `-a` hoặc `--all` khi sử dụng lệnh `git branch`.

```bash
git branch -a
```

Nhớ rằng các nhánh từ xa (remote branches) sẽ bắt đầu bằng `remotes/` khi được liệt kê.


### Delete branch

Để xóa một nhánh trong Git, bạn có thể sử dụng lệnh `git branch -d` hoặc `git branch -D`. Dưới đây là sự khác biệt giữa hai lựa chọn này:

1. **Xóa Nhánh Đã Merge (Safe Delete):**
   ```bash
   git branch -d tên_nhánh
   ```
   Lựa chọn `-d` (hoặc `--delete`) được sử dụng để xóa nhánh đã merge. Git sẽ kiểm tra xem các thay đổi trên nhánh đã được merge vào nhánh khác hay chưa. Nếu đã merge, nhánh sẽ bị xóa. Nếu chưa merge, Git sẽ cảnh báo và không xóa nhánh.

   Ví dụ:
   ```bash
   git branch -d feature-branch
   ```

2. **Xóa Nhánh Chưa Merge (Force Delete):**
   ```bash
   git branch -D tên_nhánh
   ```
   Lựa chọn `-D` (hoặc `--force-delete`) sẽ xóa nhánh mà không kiểm tra xem nó đã được merge hay chưa. Sử dụng lựa chọn này khi bạn chắc chắn rằng bạn muốn xóa nhánh mà không cần kiểm tra merge.

   Ví dụ:
   ```bash
   git branch -D feature-branch
   ```

Lưu ý rằng khi xóa nhánh, mọi thay đổi chưa commit trên nhánh đó sẽ bị mất. Hãy chắc chắn rằng bạn đã commit và lưu trữ tất cả các thay đổi cần thiết trước khi xóa nhánh.

3. **Xóa Nhánh Remote:**

   Để xóa một nhánh từ kho lưu trữ từ xa (remote repository), bạn có thể sử dụng lệnh `git push` với tùy chọn `--delete`. Dưới đây là cách để thực hiện điều này:

   ```bash
   git push origin --delete tên_nhánh
   ```

   Trong đó:
   - `origin` là tên của kho lưu trữ từ xa (remote repository). Thay `origin` bằng tên từ xa của bạn nếu có tên khác.
   - `--delete` là tùy chọn cho biết bạn đang thực hiện một hoạt động xóa.
   - `tên_nhánh` là tên của nhánh bạn muốn xóa.

   Ví dụ, để xóa một nhánh từ xa có tên là `feature-branch`, bạn có thể thực hiện lệnh sau:

   ```bash
   git push origin --delete feature-branch
   ```

Lưu ý rằng việc xóa nhánh từ xa sẽ không ảnh hưởng đến bản sao local của bạn. Để đồng bộ hóa với những thay đổi từ xa, bạn có thể sử dụng lệnh `git fetch -p` để loại bỏ các nhánh từ xa đã bị xóa.