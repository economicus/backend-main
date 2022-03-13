# Economicus backend

> 이코노미쿠스 프로젝트에 대한 설명 (작성 예정)

## Requirements

- `docker & docker-compose`
- `.env` `.env.mysql` `.env.mongo` `nginx.conf`
- 실행 방법

  ```shell
  git clone https://github.com/economicus/backend-main

  docker-compose up -d --build
  ```

## Overview
Main 서버는 이코노미쿠스 프로젝트에서 전반적인 기능을 담당하는 서버입니다. RESTFUL API 인터페이스의 
구조로 작성되었으며, 하위 서버인 퀀트 서버와 gRPC를 통해 데이터를 주고 받습니다.

## Software Used
- Docker & docker-compose
- Golang
    - gin
    - gorm
- logrus
- jwt
- Swagger
- mysql DB
- mongo DB
- AWS S3

## More