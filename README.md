# Calendar Event

## Feature
- bisa invite meeting
- blocking schedule (tanpa meeting)
- meeting online
- notif
    - reminder
        - email subscription (e.g. for activity notification)
        - 
    - push notif
- user langsung seed

## Domain
Domain: sebuah batasan masalah
A bounded context (BC) is the space in which a term has a definite and unambiguous meaning.

Entity: objek didalam domain yg punya lifecycle

- Scheduling (meeting-schedule) -> blocking calendar for event
- Reminder

### Scheduling
// TODO

### Reminder
// TODO

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

DB consideration -> Cap Theorem
https://youtu.be/vIA8VJdnpcs
https://stackoverflow.com/questions/12346326/cap-theorem-availability-and-partition-tolerance
https://medium.com/@hnasr/postgres-vs-mysql-5fa3c588a94e
https://hasura.io/learn/graphql/backend-stack/languages/go/
https://www.educative.io/courses/grokking-modern-system-design-interview-for-engineers-managers