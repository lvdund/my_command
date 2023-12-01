Để lọc các loại gói tin cụ thể trong Wireshark, bạn có thể sử dụng bộ lọc dựa trên các đặc điểm cụ thể của gói tin hoặc dựa trên giao thức. Dưới đây là một số ví dụ:

1. **Lọc theo Giao thức:**
   - Lọc tất cả các gói tin ICMP:
     ```
     icmp
     ```
   - Lọc tất cả các gói tin TCP:
     ```
     tcp
     ```
   - Lọc tất cả các gói tin UDP:
     ```
     udp
     ```

2. **Lọc theo Cổng:**
   - Lọc tất cả các gói tin trên cổng 80 (HTTP):
     ```
     tcp.port == 80
     ```

3. **Lọc theo Loại Dịch vụ:**
   - Lọc tất cả các gói tin DNS:
     ```
     dns
     ```
   - Lọc tất cả các gói tin DHCP:
     ```
     dhcp
     ```

4. **Lọc theo Địa chỉ IP:**
   - Lọc tất cả các gói tin từ hoặc đến một địa chỉ IP cụ thể:
     ```
     ip.addr == 192.168.1.1
     ```

5. **Lọc theo Nội dung:**
   - Lọc tất cả các gói tin chứa chuỗi "CONNECT" trong nội dung:
     ```
     frame contains "CONNECT"
     ```

Nhớ rằng các bộ lọc này có thể được kết hợp để tạo ra các bộ lọc phức tạp hơn để hiển thị chính xác những gói tin bạn quan tâm. Nhập bộ lọc vào ô "Display Filter" và nhấn Enter hoặc nhấn nút "Apply" để áp dụng bộ lọc.