
### 1. Install Pyenv

- [pyenv + virtualenv](../python/pyenv.md)
- Remember set default active virtualenv.

### 2. Install old kernel version

- Check `uname -r`

- Need install linux kernel version [5.4.x or 5.10.x](../debian12/old-kernel-version.md).

### 3. Install mininet

```bash
sudo apt install mininet
```

### 4. Script and Config for etrib5gc

- Copy [this folder](../run%20etrib5gc/mininet/) to etrib5gc repo.
- Run:
    ```bash
    ./runtest.sh
    ```

### 5. Some conflict driver

- Missing `ovs-controller`:
    ```bash
    sudo ln -s /usr/bin/ovs-testcontroller /usr/bin/controller
    ```

- Error cgroup:

    ```bash
    sudo nano /etc/default/grub
    ```
    Then modify `GRUB_CMDLINE_LINUX`
    ```bash
    GRUB_CMDLINE_LINUX="cgroup_enable=memory systemd.unified_cgroup_hierarchy=0"
    ```