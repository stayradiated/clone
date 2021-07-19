# clone

**clone a specific repo**

Will clone to `~/src/github.com/stayradiated/clone`

```shell
clone github.com/stayradiated/clone
```

**shallow clone**

Just get the latest commit.

```shell
clone --shallow github.com/stayradiated/clone
```

**checkout specific commit**

```shell
clone --ref='441c8c55' github.com/stayradiated/clone
```

**checkout specific tag**

```shell
clone --tag='v1.4.0' github.com/stayradiated/clone
```

**use https (instead of ssh)**

By default, `clone` will use `git@github.com:` prefix.
Pass this flag to use `https://` prefix.

```shell
clone --https github.com/stayradiated/clone
```
