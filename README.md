# 📘 `env-info` API 명세서

> DeFi 생태계 컨트랙트 분석을 위한 모듈  
> 주요 dApp 또는 Token 컨트랙트 주소를 기준으로 메타정보를 제공합니다.

---

## 🧩 Base URL

```
http://localhost:8023/env-info
```

---

## ✅ 1. [GET] `/ping`

### 📌 설명  
모듈의 정상 작동 여부를 확인하는 간단한 핑 테스트입니다.

### 🔹 Request

- **Method**: `GET`
- **URL**: `/env-info/ping`

### ✅ 성공 응답: `200 OK`

```json
{
  "module": "env-info",
  "message": "pong"
}
```

---

## ✅ 2. [GET] `/contract/:address`

### 📌 설명  
지정된 이더리움 컨트랙트 주소에 대해  
- dApp 또는 Token인지 판별하고,  
- 관련 메타정보를 반환합니다.

### 🔹 Request

- **Method**: `GET`
- **URL**: `/env-info/contract/0x1234...abcd`
- **Path Parameter**:
  - `address` (string, required): Ethereum 컨트랙트 주소

### ✅ 응답 예시 – dApp

```json
{
  "address": "0x...",
  "name": "aave",
  "category": "lending",
  "chain": ["Ethereum", "Polygon", "Arbitrum"],
  "tvl": 1200000000,
  "tx_count_total": 123000,
  "tx_count_30d": 8200,
  "users_count": 6700,
  "deployed_at": "2020-12-03T00:00:00Z"
}
```

### ✅ 응답 예시 – Token

```json
{
  "address": "0x...",
  "name": "Aave Token",
  "symbol": "AAVE",
  "decimals": 18,
  "chain": ["Ethereum"]
}
```

### ❌ 에러 응답 예시

#### 400 Bad Request

```json
{
  "error": "address is required"
}
```

#### 500 Internal Server Error

```json
{
  "error": "failed to fetch contract info: <error details>"
}
```

---
추가로 어떤 포맷이 더 필요할까? Markdown 그대로 README에 붙여도 돼.
