# gum

gum is short for Git User Manager. It can help you to switch git user account conveniently when working on several git repositories, for example, your job and side projects.


# Installing

```sh
go install github.com/chenlujjj/gum@latest
```
# Usage


List all git user accounts.
```sh
$ gum list
name: chenlujjj, email: 123456789@qq.com
name: chenluxin, email: 987654321@qq.com
```

Add a git user account.

```sh
$ gum add clx clx@sina.com
```


Delete a git user account.
```sh
$ gum delete clx
```


Set a user to current git repository local config.
```sh
$ gum set chelujjj
```
