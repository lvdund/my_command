Post Install Arch linux (updated May-8-2025)

- Init
```bash
sudo pacman -S base-devel linux-headers git nano wget curl --needed
```

- UI
```bash
paru -S ttf-firacode-nerd
```

- Install virtualbox & Vagrant
```bash
paru -S linux-lts-headers vagrant virtualbox
sudo modprobe vboxdrv
```

- docker
```bash
paru -S docker docker-compose
```

- nvidia
```bash
sudo nano /etc/pacman.conf
# for enable [multilib]
paru
paru -S nvidia-dkms nvidia-utils lib32-nvidia-utils
paru -S nvidia-settings
```
