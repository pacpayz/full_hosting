## Handling Git Submodules

**Make Sure to Clone Submodules**

```
git clone --recurse-submodules https://github.com/pacpayz/full_hosting.git
```

**Updating Submodules (Pulling)**

```
git submodule update --remote --merge
```

**Updating Submodules (Pushing)**

```
cd ./submodule-folder
git add .
git commit -m "Updated submodule"
git push
```

## Handling Go Submodules

**Using Go Work To Setup**

```
go work init
```

**Adding Go Submodules**

```
go work use /path/to/submodule
```

**Example go submodules** 
> These submodules already exist in this project

```
go work use ./blockchain
go work use ./extensions
```