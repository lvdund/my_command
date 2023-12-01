`tcpdump` là một công cụ mạng mạnh mẽ trên hệ điều hành Unix/Linux được sử dụng để theo dõi và phân tích gói tin trên mạng. Dưới đây là một số ví dụ về cách sử dụng `tcpdump`:

1. **Hiển thị tất cả các gói tin trên một giao diện cụ thể:**
   ```bash
   sudo tcpdump -i eth0
   ```

   Thay `eth0` bằng tên giao diện mạng mà bạn muốn theo dõi.

2. **Lưu kết quả vào một tệp:**
   ```bash
   sudo tcpdump -i eth0 -w output.pcap
   ```

   Ghi tất cả các gói tin vào tệp `output.pcap` để sau này bạn có thể mở và phân tích chúng bằng các công cụ như Wireshark.

3. **Hiển thị các gói tin với cú pháp đọc được:**
   ```bash
   sudo tcpdump -A -i eth0
   ```

   Hiển thị nội dung của các gói tin một cách đọc được.

4. **Hiển thị gói tin DNS:**
   ```bash
   sudo tcpdump -i eth0 -n port 53
   ```

   Hiển thị gói tin DNS trên cổng 53.

5. **Lọc theo địa chỉ nguồn và đích:**
   ```bash
   sudo tcpdump -i eth0 src host 192.168.1.1 and dst host 192.168.1.2
   ```

   Lọc gói tin từ địa chỉ nguồn `192.168.1.1` đến địa chỉ đích `192.168.1.2`.

6. **Lọc theo cổng và giao thức:**
   ```bash
   sudo tcpdump -i eth0 port 80 and tcp
   ```

   Lọc gói tin trên cổng 80 và sử dụng giao thức TCP.

7. **Lọc theo kích thước gói tin:**
   ```bash
   sudo tcpdump -i eth0 greater 100
   ```

   Hiển thị các gói tin có kích thước lớn hơn 100 byte.

8. **Hiển thị các gói tin ARP:**
   ```bash
   sudo tcpdump -i eth0 arp
   ```

   Hiển thị các gói tin ARP.

Lưu ý rằng việc sử dụng `tcpdump` yêu cầu quyền đặc quyền root hoặc sudo. Để biết thêm chi tiết và tùy chọn, bạn có thể kiểm tra tài liệu của `tcpdump` (`man tcpdump`).