# 자동 로그 JSON 파서

로그 문자열에서 JSON을 자동으로 추출하고 분석하는 웹 애플리케이션입니다.

## 실행 방법

### Option 1: Node.js (추천)

```bash
node server.js
```

또는 npm 스크립트 사용:

```bash
npm start
```

### Option 2: Go

```bash
go run server.go
```

## 접속

서버 실행 후 브라우저에서 접속:

```
http://localhost:9990
```

## 기능

- 로그 문자열에서 JSON 자동 추출
- 중첩된 JSON 파싱
- output 배열 타입 추정 및 매핑
- JSON 다운로드 및 복사
- Go Struct 코드 자동 생성
- 샘플 데이터 테스트

## 포트 변경

포트를 변경하려면:
- Node.js: `server.js`의 `PORT` 변수 수정
- Go: `server.go`의 `PORT` 상수 수정

