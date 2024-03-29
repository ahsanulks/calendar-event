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


# doc peer
https://docs.google.com/document/d/1py_jUw93c5448axUF9EBrXUOWYMuSNVdAjK5K3bC5DM/edit#

// facade pattern https://refactoring.guru/design-patterns/facade
// facade -> clean architecture https://medium.com/@mammimia/clean-architecture-part-iv-architecture-5eee3cb5656a
// func (su *SchedulerUsecase) GetSchedule() {
// 	a := su.ABC
// 	b := su.MicroServX
// 	return a + b
// }


// Test methodology
// fake
// dummy
// stub
// spy -> kita akan melakukan pengecekan kalo sebuah function di panggil
//			e.g write.create di call 1x
// mock -> kita akan melakukan mocking function tersebut
// 			e.g kita melakukan ekspektasi kalo function tersebut di call dengan parameter dan ngemock response

// https://jesusvalerareales.com/testing-with-test-doubles/