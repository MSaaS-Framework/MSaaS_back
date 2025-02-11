# 1단계: 빌드 스테이지
FROM golang:alpine AS builder

# 환경 변수 설정
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 작업 디렉토리 설정
WORKDIR /build

# 소스 코드 복사
COPY . .

# 의존성 다운로드
RUN go mod download

# wizcraft 바이너리 빌드
RUN go build -o wizcraft ./cmd/wizcraft/wizcraft.go

# 2단계: 실행 스테이지
FROM scratch

# 빌드된 바이너리를 복사
COPY --from=builder /build/wizcraft /wizcraft

# 웹 서버 포트 설정
EXPOSE 8080

# 컨테이너가 시작될 때 실행할 명령어
ENTRYPOINT ["/wizcraft", "start"]
