# GoShorURL
由 Go 语言撰写的短链服务，适合私有部署。

# 发号器
一、 `MurmurHash`，由 Google 出品的非加密性哈希函数，在这个场景下，对比 `MD5`/`SHA` 性能更好（十倍或以上），因为只关心运算速度和冲突概率。  

二、自增序列
暂不考虑。

默认使用 `MurmurHash 32bit`。

# 重定向
`301`， 更高性能，统计失效。  
`302`，影响性能，便于统计。（推荐）

# 其它
## 传统布隆过滤器
存在误差，不支持删除，暂不采用方案。

