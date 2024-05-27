## My Command

### All my config files

- Check this [folder](./config-file/)

### Capture package

- [tcpdump](./tcpdump_wireshark/tcpdump.md)
- [wireshark](./tcpdump_wireshark/wireshark.md)

### Git command

- [Begin config](./git_command/begin_config.md)
- [2 repo](./git_command/2_repo.md)
- [branch](./git_command/branch.md)
- [err](./git_command/err.md)

### Python

- [Create virtual env](./python/python-virtualenv.md)

- (Recommend)[Read](./python/pyenv.md)

### Docker command

- [Command](./docker_command/README.md)
- [mongodb](./docker_command/mongodb.md)

### Run Etrib5gc

- [Read config](./run%20etrib5gc/readme.md)
- [Download run/](./run%20etrib5gc/run/) to etrib5gc/
- [Setup vm virtualbox (không có bước config)](./run%20etrib5gc/setup_vms.md)
- [mininet](./mininet/etrib5gc.md)

### Nodejs nvm

- [nvm command](./nodejs/nvm%20command.md)

### Arch Linux

- [Command](./arch_linux/command.md)
- [Delete](./arch_linux/delete_pkg.md)
- [yay](./arch_linux/packed%20AUR%20helper%20yay.md)
- [paru](./arch_linux/packed%20AUR%20helper%20paru.md)
- [process](./arch_linux/process.md)

### mininet

- [connect host to external](./mininet/h1-to-google.md)
- [setup for etrib5gc testing](./mininet/etrib5gc.md)

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
