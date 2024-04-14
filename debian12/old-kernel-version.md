How to install linux kernel version `5.10` on Debian 12.

1. **Add repo:**
   ```bash
   sudo nano /etc/apt/sources.list.d/bullseye-repo.list
   ```
   Then add source:
   ```bash
   deb http://deb.debian.org/debian bullseye main
   deb-src http://deb.debian.org/debian bullseye main
   ```

2. **Install:**
   ```bash
   sudo apt update
   sudo apt install linux-image-5.10.0-28-amd64 linux-headers-5.10.0-28-amd64
   ```

3. **Remember rm source after install**

