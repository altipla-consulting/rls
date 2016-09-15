
# rls

> Release a new version of the repo to production.

### Install

See the latest [GitHub release](https://github.com/altipla-consulting/rls/releases) to download the compiled binary. Leave it somewhere you can call it from the shell (in your PATH). For example download it to `/usr/local/bin/rls`.

### Usage

When you are ready to release a new version to production run the command in the `master` branch of the repo:

```shell
rls
```

*rls* will:
- Confirm the latest change that will be sent.
- Copy all the missing commits from `master` to `release`.
- Push the `master` and `release` branchs to the remote (e.g.: GitHub).
