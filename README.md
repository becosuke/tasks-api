# はじめに

動作させるには

```
cd ${WORKDIR}/tasks-docker
docker-compose up -d taksk-golang

cd ${GOPATH}/src/github.com/becosuke/tasks-api
make build
```

# 仕様

## lists

- inbox
- next actions / today
- waiting for / delegate
  - 誰待ちなのか明示して通知を送れると良い
  - 通知だけでなく子タスクを発生させて送りつけられたらなお良い
- tomorrow / later
- someday / maybe / incubate
- calendar
- repeating / scheduled
- trashed
- done

## contexts

- 場所の概念
  - home
  - office
  - commuter route / 通勤路
- 時間の概念
  - 起床
  - 出勤
  - 退勤
  - 就寝
  - 朝食
  - 昼食
  - おやつ
  - 夕食
  - 夜食
  - morning / 朝
  - afternoon / 昼
  - evening / 晩
  - night / 夜

# 設計

## API

| method | endpoint | parameter |
|---|---|---|
| GET | /lists | |
| GET | /list/:id | |
| GET | /contexts | |
| GET | /context/:id | |
| POST | /task | |
| PUT | /task/:id | |
| GET | /tasks/all | limit offset |
| PUT | /task/:id/completed | |
| DELETE | /task/:id/completed | |
| GET | /task/:id/completed | |
| GET | /tasks/completed | limit offset |
| PUT | /task/:id/deleted | |
| DELETE | /task/:id/deleted | |
| GET | /task/:id/deleted | |
| GET | /tasks/deleted | limit offset |
| PUT | /task/:id/description | |
| DELETE | /task/:id/description | |
| GET | /task/:id/description | |

## 型定義

### list

| name | type |
| --- | --- |
| id | uint64 |
| title | string |

### task

| name | type |
| --- | --- |
| id | uint64 |
| list_id | uint64 |
| title | string |
| completed_at | int64 |
| created_at | int64 |
| updated_at | int64 |
| deleted_at | int64 |

### description

| name | type |
| --- | --- |
| id | uint64 |
| task_id | uint64 |
| description | string |
| created_at | int64 |
| updated_at | int64 |
| deleted_at | int64 |
