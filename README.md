# clone

**clone a specific repo**

Will clone to `~/src/github.com/stayradiated/clone

```shell
clone github.com/stayradiated/clone
```

**shallow clone**

Just get the latest commit.

```shell
clone --shallow github.com/stayradiated/clone

**shallow ref clone**

Just get a specific commit or tag

```shell
clone --shallow --ref='441c8c55' github.com/stayradiated/clone
```

**use https (instead of ssh)**

By default, `clone` will use `git@github.com:` prefix.
Pass this flag to use `https://` prefix.

```shell
clone --https github.com/stayradiated/clone
```
