# Calendar Event

## Domain
Domain: sebuah batasan masalah
A bounded context (BC) is the space in which a term has a definite and unambiguous meaning.

Entity: objek didalam domain yg punya lifecycle

- Scheduling (meeting-schedule) -> blocking calendar for event
- Reminder

### Schedule
- blocking schedule (tanpa meeting)
- bisa invite meeting (lebih ke blocking calendar)
- meeting online

### Reminder
- email subscription (e.g. for activity notification)
    - event
        - reschedule
        - delete & cancelation event
        - invitation -> to receiver
        - confirmation -> to sender
- reminder before x minutes times

## Architecture

-----------     API      ------------     -------------
| client  | ---GATEWAY-- |    API   |-----| scheduler |
-----------     |        ------------     -------------
                |              | 
-----------     |              |
| client  | -----        ------------
-----------              | database |
                         ------------

## Technology
- Languange -> go v1.20 (gin)
- PROTOCOL  -> graphql (karena pengen), REST
- Database  -> postgresql (consistency & availability)
- ORM       -> gorm
- client    -> sementara, CLI

- Architectural pattern -> hexagonal pattern
- Software development process -> TDD

## Data
- user langsung seed

DB consideration -> Cap Theorem
https://youtu.be/vIA8VJdnpcs
https://stackoverflow.com/questions/12346326/cap-theorem-availability-and-partition-tolerance
https://medium.com/@hnasr/postgres-vs-mysql-5fa3c588a94e
https://hasura.io/learn/graphql/backend-stack/languages/go/
https://www.educative.io/courses/grokking-modern-system-design-interview-for-engineers-managers

## Layout
https://github.com/golang-standards/project-layout

## repository
https://learn.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/infrastructure-persistence-layer-design

https://orkhan.gitbook.io/typeorm/docs/active-record-data-mapper