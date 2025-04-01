# 核心模块设计文档

## 交易发送引擎
```mermaid
sequenceDiagram
    participant User
    participant Scheduler
    participant GasOptimizer
    participant Contract
    participant Blockchain
    
    User->>Scheduler: 启动监控循环
    loop 每个区块周期
        Scheduler->>GasOptimizer: 获取实时Gas价格
        GasOptimizer->>Blockchain: 查询baseFee
        GasOptimizer-->>Scheduler: 返回优化后的Gas参数
        Scheduler->>Contract: 构造frontrun交易
        Contract->>Blockchain: 发送签名交易
        Blockchain-->>Contract: 返回交易哈希
        Contract-->>Scheduler: 监控交易状态
        Scheduler->>User: 报告交易结果
    end
```

## 关键数据结构
```rust
#[derive(Debug, Clone)]
pub struct FrontrunConfig {
    /// 目标合约地址
    pub contract_address: Address,
    /// 最大Gas价格 (Gwei)
    pub max_gas_price: U256,
    /// 交易间隔 (秒)
    pub interval_secs: u64,
    /// 紧急停止开关
    pub emergency_stop: AtomicBool,
}

#[derive(Debug, Clone, Copy)]
pub struct GasParams {
    pub base_fee: U256,
    pub priority_fee: U256,
    pub gas_limit: U256,
}
