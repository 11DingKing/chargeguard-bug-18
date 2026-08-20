# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

已删除站点再次上报隐患时，repository 返回的是“站点不存在”，HTTP 接口却变成 409“状态冲突”。请纠正错误分类：不存在应为 404，真正的乐观锁冲突仍是 409，存储故障继续保留原始错误链。错误映射测试用例不得改动，更不能放宽状态码断言。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-18
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-18.git
- parent SHA：96b5d694063b2fd946c6d6131e4459135e5f5143

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-18.git bug-repro
cd bug-repro
git checkout --detach 96b5d694063b2fd946c6d6131e4459135e5f5143
go test ./internal/httpapi -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:13: status=409 body=state conflict
FAIL
FAIL	chargeguard/internal/httpapi	0.063s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:13: status=409 body=state conflict
FAIL
FAIL	chargeguard/internal/httpapi	0.003s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，在题面描述的触发条件下应得到预期业务结果且不再出现原始症状；定向验证命令修复前必须失败、应用修复后必须通过，相关回归和仓库全量测试必须通过；不得新增、删除或修改测试文件，不得跳过测试、降低断言或绕过目标逻辑。
