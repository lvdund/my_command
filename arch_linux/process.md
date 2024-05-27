1. **Run in background:**
   - Chuyển hướng đầu ra chuẩn của lệnh vào tệp output.log

   ```bash
   command > output.log 2>&1 &
   ```

2. **Check running process:**
   - `ps aux | grep "name"`: Find.
   - `sudo netstat -tunpl`: Check running port with ID process.
   - `kill {id}`: Kill process.
   - `cat /proc/net/sctp/eps`: Check SCTP open port.