# git

## install git

安装略……

## git config

git 的基本配置：首先 git 配置分为三层，分别是系统级、用户级、仓库级

对应的配置文件如下（Unix 系列）：

- 系统级 system：`/etc/gitconfig`
- 用户级 user：`~/.gitconfig`
- 仓库级 repository：`<repo>/.git/config`

windows 的配置文件路径：（不一定对，我没验证，模型一键生成的）

- 系统级 system：`C:\Program Files\Git\etc\gitconfig`
- 用户级 user：`C:\Users\<username>\.gitconfig`
- 仓库级 repository：`<repo>/.git/config`

git 配置的优先级：

- 仓库级 > 用户级 > 系统级

### git 配置文件的格式

git 使用的是 INI 格式，[section] 是 section 名，section 中的 key=value 是配置项。通俗易懂来讲，section 是配置的分类，到某个分类下，再配置具体的配置项。key 是配置的名称，value 是配置的值。

```ini
[section]
key=value
```

具体的 git 配置文件示例如下：

```ini
[user]
    name = ZhouThreeJin
    email = zhouthreejin@gmail.com
[core]
    editor = code --wait
    autocrlf = input
[alias]
    st = status
    ci = commit
    co = checkout
    br = branch
    lg = log --oneline --graph --decorate --all
```

上述配置文件中，user 是配置分类，name 和 email 是配置项。user 这个 section 配置了 git 用户的用户名和邮箱。当你在提交代码时，git 就会用你的 user.name 和 user.email 来标识你的提交。

举个例：

```bash
# 运行这行命令查看提交日志
git log
```

git log 的输出就会显示每一个 commit 的作者和邮箱。如果有你提交的 commit，那么就会显示你的用户名和邮箱。

```
commit 1234567890 (HEAD -> master)
Author: ZhouThreeJin <zhouthreejin@gmail.com> # 你的用户名和邮箱
Date:   Mon Dec 30 10:00:00 2024 +0800
```

还例如，如果你在 commit 的时候，没有配置 user.name 和 user.email，那么 git 就会使用系统级的配置。也就是你的电脑用户的 username 作为 user.name，你的电脑用户的 username + hostname 作为 user.email。

```bash
commit 03eb4e69fd696c37ce8f044aa0b6a0934e585512 (HEAD -> scoheart, origin/scoheart)
Author: Scoheart <scoheart@Scohearts-Laptop.local> # 没有配置 user.name 和 user.email 时候的默认配置
Date:   Mon Dec 30 22:44:31 2024 +0800

    chore: 武林秘籍
```

这里就仅仅用 user 这个 section 的 name 和 email 配置来示例，其他的一些 section 也是比较重要的配置，例如 core、alias 等。需要全面的了解一下 git 的配置。

### git config 配置方式

git config 的配置方式有两种：

- 命令行命令配置
- 配置文件配置

1. 命令行命令配置

```bash
# 设置用户名 和 邮箱，--global 表示全局配置，即对所有仓库生效
git config --global user.name "ZhouThreeJin" 
git config --global user.email "zhouthreejin@gmail.com"
```

```bash
# 设置仓库级配置，只对当前仓库生效
git config user.name "ZhouThreeJin"
git config user.email "zhouthreejin@gmail.com"
```

2. 配置文件配置

直接编辑配置文件，添加配置项。
