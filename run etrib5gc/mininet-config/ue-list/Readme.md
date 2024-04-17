Install and start MongoDb

Import the test UE subsciber database:

```bash
mongoimport --db "etrib5gc" --collection "authsub" --file etrib5gc-uelist.json
```
Configuting UDM with db.name = "etrib5gc", db.auth="authsub" and db.mock=false

