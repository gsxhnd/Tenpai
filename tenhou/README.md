# Tenhou

```shell
# 凤凰卓对局每小时数据信息 只包含场次ID
# https://tenhou.net/sc/raw/dat/2023/scc20230101.html.gz

# 凤凰卓每年的对局数据 只包含场次ID
# https://tenhou.net/sc/raw/scraw2022.zip

# 单场对局转换成json
# 通过场次ID获取该场次的对局信息
# https://tenhou.net/5/mjlog2json.cgi?2023100405gm-00a9-0000-1c001b06
```

1. 下载年数据
2. 下载当年的每小时数据
3. 解析HTML中的所有牌谱
   1. 分类 4人 3人
   2. 分类 东风战/一个半庄
4. 转换牌谱到具体对局信息

json 命名  年月日小时分钟-场次ID 200902201110-2009022011gm-00a9-0000-d7935c6d
