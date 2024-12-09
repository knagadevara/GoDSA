module github.com/knagadevara/GoDSA

go 1.23.3

replace github.com/knagadevara/GoDS/hashmap => ./hashmap

replace github.com/knagadevara/GoDS/linkedlist => ./linkedlist

replace github.com/knagadevara/GoDS/queue => ./queue

replace github.com/knagadevara/GoDS/stack => ./stack

require (
	github.com/knagadevara/GoDS/hashmap v0.0.0-00010101000000-000000000000
	github.com/knagadevara/GoDS/linkedlist v0.0.0-00010101000000-000000000000
	github.com/knagadevara/GoDS/queue v0.0.0-00010101000000-000000000000
	github.com/knagadevara/GoDS/stack v0.0.0-00010101000000-000000000000
)
