# 术语表

> 状态：速查表

用于统一 `Suphx` 相关术语，避免不同文档中出现错误翻译或含义漂移。

## 模型

- `discard model`：弃牌模型
- `riichi model`：立直模型
- `chi model`：吃模型
- `pon model`：碰模型
- `kan model`：杠模型
- `reward predictor`：奖励预测器

## 训练与评估

- `global reward prediction`：全局奖励预测
- `oracle guiding`：Oracle 引导
- `pMCPA`：参数化蒙特卡洛策略适配
- `supervised learning`：监督学习
- `self-play reinforcement learning`：自博弈强化学习
- `offline evaluation`：离线评估
- `online evaluation`：在线评估

## 平台与指标

- `Tenhou`：天凤
- `phoenix room`：凤凰桌
- `record rank`：记录段位
- `stable rank`：稳定段位

## 不建议的写法

- `全局奖励塑形`
- `论文主线是 MCTS`
- `Suphx 使用单一 policy/value 网络覆盖全部动作`
- `训练完成后默认进入雀魂验证`
- `Rust 验证层已经存在`

## 推荐写法

- `当前训练设计以五个动作模型和 reward predictor 为主线`
- `在线评估若要与论文对齐，应优先参考 Tenhou 语境`
- `雀魂方向属于实验性扩展`

## 相关文档

- [04-TRAINING.md](./04-TRAINING.md)
- [02-ARCHITECTURE.md](./02-ARCHITECTURE.md)
- [03-DATA_COLLECTION.md](./03-DATA_COLLECTION.md)
- [07-MAJSOUL_VALIDATION.md](./07-MAJSOUL_VALIDATION.md)
