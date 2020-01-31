module github.com/opentelekomcloud/dockermachinedriver

go 1.13

replace github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2 // indirect

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Sirupsen/logrus v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/machine v0.16.2
	golang.org/x/crypto v0.0.0-20200128174031-69ecbb4d6d5d // indirect
)
