## How To Run

### I. Setup:

1. Chạy NFs core trên `máy thật`.
2. Chạy các UPF, gnb+ue thành các `máy ảo` khác nhau.
3. Chỉ dùng 1 network là `Bridged Adapter`

   Máy khác sẽ có name `enp0s31f6` khác - nó chính là tên default network máy chính - dùng `ifconfig` để xem.

   <img src="./images/network%20adapter%20virtualbox.png" alt="Image" style="width: 70%; max-width: 500px;">


3. Folder `./run` - **copy vào folder etrib5gc**:
   ```
   ├── run
   │   ├── amf.sh
   │   ├── aufs.sh
   │   ├── controller.sh
   │   ├── damf.sh
   │   ├── gnb.sh
   │   ├── pcf.sh
   │   ├── pran.sh
   │   ├── run.sh **
   │   ├── smf.sh
   │   ├── stop.sh **
   │   ├── udm.sh
   │   ├── ue.sh
   │   ├── upf.sh
   │   └── upmf.sh
   ```
4. Config:

   Setup 2 máy ảo. 
   Example:
   
   - NFs core:
      - ip: 192.168.1.56
      - giữ nguyên file config
   - UPF:
      - ip: 192.168.1.75
      - mẫu [upf.json](./upf.json)
         - ip dòng 12 === ip của `máy thật`
         - ip dòng 24, 30 === ip của `máy upf`
   - gnb+ue:
      - ip: 192.168.1.74
      - mẫu [gnb.json](./free5gc-gnb.yaml)
         - ip dòng 8, 9, 10 === ip của `máy gnb+ue`
         - ip dòng 14 === ip của `máy thật`
      - mẫu [ue.json](./free5gc-ue.yaml#L23)
         - ip dòng 23 === ip của `máy gnb+ue`


### II. Run:

0. Trước khi chạy:

   - B1: phân quyền file *.sh
      
      ```bash
      chmod +x ./run/*
      ```
      
   - B2: setup route

      ```bash
      ./run/before_run.sh
      ```

   **Lưu ý:** delay 1-2s giữa các bước.

1. Chạy các NFs core:

   **Bao gồm:** controller, aufs, udm, pcf, amf, damp, pran, smf, upmf.
    
   `Lưu ý:` không chạy upf.

   ```bash
   ./run/run.sh
   ```

2. Chạy UPF:

   - B1: 5G GTP-U kernel module

      ```bash
      git clone -b v0.8.3 https://github.com/free5gc/gtp5g.git

      cd gtp5g
      make
      sudo make install
      ```
   - B2: chạy upf

      ```bash
      ./run/upf1.sh
      ```

3. Chạy gnb, ue:

   - B1: chạy gnb

      ```bash
      ./run/gnb.sh
      ```

   - B2: chạy ue

      ```bash
      ./run/ue.sh
      ```

4. Dùng [tcpdump](../tcpdump_wireshark/tcpdump.md) để bắt các gói tin

5. Terminate:

   ```bash
   ./run/stop.sh
   ```