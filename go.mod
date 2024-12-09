module github.com/knagadevara/GoDSA

go 1.23.3

replace github.com/knagadevara/GoDS/hashmap v0.0.2 => ./hashmap

replace github.com/knagadevara/GoDS/linkedlist v0.0.2 => ./linkedlist

replace github.com/knagadevara/GoDS/queue v0.0.2 => ./queue

replace github.com/knagadevara/GoDS/stack v0.0.2 => ./stack

require (
	github.com/knagadevara/GoDS/hashmap v0.0.2
	github.com/knagadevara/GoDS/linkedlist v0.0.2
	github.com/knagadevara/GoDS/queue v0.0.2
	github.com/knagadevara/GoDS/stack v0.0.2
)
