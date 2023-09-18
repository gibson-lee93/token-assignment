# Test Assignment

### 개요 ###
BSC Testnet 네트워크를 사용하여 Chainlink 컨트랙스와 Bitfinex API로 부터 토큰 가격 정보를 조회하고 데이터베이스에 저장하여 클라이언트에게 토큰 가격 정보를 조회할 수 있도록 API를 제공.

---

### 개발환경
Language: GoLang<br>
Framework: Echo<br>
ORM: Bun<br>
DB: MySQL<br>

---

### 서버 실행 환경
서버 포트: `localhost:1323`<br>
MySQL DB 포트: `3306`

---


### 실행방법
#### 코드를 받아서 실행하는 방법
**Prerequisites:**
- Golang
- MySQL (비밀번호 없이 username:`root`로 localhost:3306으로 실행)

1. `database` 파일에 있는 `token.sql` 스크립트 실행하여 서버 실행에 필요한 데이터베이스와 테이블 생성.
2. 터미널에서 `go run main.go`를 입력하여 서버 실행.
3. 서버가 실행되면 토큰 가격 정보를 조회하는 스케쥴러가 30초 뒤에 30초 간격으로 실행.

---

### API 명세서
**토큰 이름으로 최신의 토큰 정보 조회**
- **Endpoint**: `GET /tokeninfo`
- **Request Body**:
  - `tokenSymbol` (토큰 이름): Not Null (`USDT`, `USDC`, `ETH` 중 택 1 )
  - `source` (가격 출저): Nullable
  - `startTime` (특정 구간 시작 시간): Nullable (`endTime` 값이 있을경우 Not Null)
  - `endTime` (특정 구간 끝 시간): Nullable (`startTime` 값이 있을경우 Not Null)
  - 예시:
``` JSON
{
    "tokenSymbol": "USDC",
    "source": "chainlink",
    "startTime": "2023-09-17T18:12:36Z",
    "endTime": "2023-09-17T18:16:00Z"
}
```
- **Response**:
  - Request Body에 토큰 이름만 입력했을 경우 토큰 이름의 최신 토큰 정보를 조회
  - Request Body에 토큰 이름과 가격 출처를 입력했을 경우 토큰 이름과 가격 출저의 최신의 토큰 정보 조회 
  - Request Body에 특정 시간 구간(`startTime` & `endTime`)을 입력했을 경우 해당 시간동안의 평균 가격 조회
  - 가격 출처가 다수인 경우 각각 출력
  - 예시:
``` JSON
{
    "TokenInfos": [
        {
            "tokenSymbol": "USDC",
            "price": 1.001,
            "source": "bitfinex",
            "timestamp": "2023-09-18T16:13:30Z"
        },
        {
            "tokenSymbol": "USDC",
            "price": 1.0003,
            "source": "chainlink",
            "timestamp": "2023-09-18T16:13:30Z"
        }
    ]
}
```
