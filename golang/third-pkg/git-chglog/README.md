# git-chglog

git-chglog 是一款用于根据 git（标签） 提交记录自动生成提交日志的工具；


# 安装

```bash
$ go get github.com/git-chglog/git-chglog/cmd/git-chglog
```


## 使用

```bash
$ git-chglog --init
```


选项（设置日志格式的模版）：

- What is the URL of your repository?: https://github.com/xxx/xxx
- What is your favorite style?: github
- Choose the format of your favorite commit message: <type>(<scope>): <subject> -- feat(core): Add new feature
- What is your favorite template style?: standard
- Do you include Merge Commit in CHANGELOG?: n
- Do you include Revert Commit in CHANGELOG?: y
- In which directory do you output configuration files and templates?: .chglog


# 输出日志

```bash
$ git-chglog  # 默认情况下输出到标准输出中，通常会记录在日志文件中；

$ git-chglog -o file # 将指定文件输出到指定路径
# $ git-chglog -o CHANGELOG/CHANGELOG-0.1.md
```