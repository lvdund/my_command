## Mongodb

**How to install**
  ```
  $ docker pull mongo
  $ docker run -d -p 2718:27017 -v ~/localPath:/data/db --name mongo-test mongo:latest
  ```
**2718 - port of localhost**
