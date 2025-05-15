## My Command

### Capture package

- [tcpdump](./tcpdump_wireshark/tcpdump.md)
- [wireshark](./tcpdump_wireshark/wireshark.md)

### Python

- [Create virtual env](./python/python-virtualenv.md)

- (Recommend)[Read](./python/pyenv.md)

### Arch Linux

- [Command](./arch_linux/command.md)
- [Delete](./arch_linux/delete_pkg.md)
- [yay](./arch_linux/packed%20AUR%20helper%20yay.md)
- [paru](./arch_linux/packed%20AUR%20helper%20paru.md)
- [process](./arch_linux/process.md)

### minikube use local image

https://www.baeldung.com/ops/docker-local-images-minikube

### setting tmux
```bash
nano .tmux.conf
```
```bash
set -g mouse on
bind -n C-k clear-history
```
```bash
tmux source .tmux.conf
```

### Setting crow translate
```bash
dbus-send --type=method_call --dest=io.crow_translate.CrowTranslate /io/crow_translate/CrowTranslate/MainWindow io.crow_translate.CrowTranslate.MainWindow.translateSelection
```

### Old kernel version

- On [Debian 12](./debian12/old-kernel-version.md)

### Learn k8s

- [github tranductrinh](https://github.com/tranductrinh/k8s)

### begin config
```bash
git config --global user.name "lvdund"
git config --global user.email lvdund@gmail.com
```
