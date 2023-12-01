## 1 repo main - 1 repo for back up

```bash
git clone --branch=<name_branch> <Repo1_URL>
```

```bash
git remote add backup_repo <Repo2_URL>
```

```bash
git add .
git commit -m "Thực hiện các thay đổi"
```

```bash
git push <name_repo> <name_branch>
```
Example:

- Push to backup repo

```bash
git push backup_repo <name_branch>
```

- Push to main repo

```bash
git push origin <name_branch>
```