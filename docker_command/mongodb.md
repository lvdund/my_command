## Mongodb

**How to install**
  ```
  $ docker pull mongo
  $ docker run -d -p 2718:27017 -v ~/localPath:/data/db --name mongo-test mongo:latest
  ```
**2718 - port of localhost**


https://flathub.org/setup/Ubuntu

sudo apt install flatpak

flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo

restart

flatpak install flathub com.mongodb.Compass
