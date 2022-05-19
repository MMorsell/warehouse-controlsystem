# Walle [![Build Status](https://dev.azure.com/martinmorsell18/Current%20Projects/_apis/build/status/Gophers.walle?branchName=master)](https://dev.azure.com/martinmorsell18/Current%20Projects/_build/latest?definitionId=8&branchName=master)

## Azure DevOps
This repository is connected to external services in "Azure DevOps". Specifically, all resources will be built and test-cases will be run automatically after each commit. This improves the stability of the codebase, since passing tests and successful build is required to merge into the Master branch. 


## Description:
The Wall-e project is a control system for robots proceeding tasks in a 2d grid, Walle-e is inspired by storage 
facilities using robots to perform simple storage tasks.

Wall-e consists of the 2 programs "the hive" and "robot", which will communicate via a gRPC system.

The hive is a central hub, essentially being the center of all operations and communication. It provides orders for the robot, which is in the form of "go to these (x,y) coordinates by taking this specific path.\
\
The robot will move to these coordinates by a received path from the hive once the robot has been given confirmation that the path is available, meaning there is no risk of a collision with another robot. Once the robot has arrived at its desired spot the task is completed and the robot can be given further orders.


![Board](/projectStructure.png)

## Requirements:

- [A go installation](https://go.dev/doc/install) (version 1.18 or higher) 
- [Protocol Buffer Compiler installed for go](https://grpc.io/docs/protoc-installation/) (version 3 or higher)

## Getting started:

1. Check that requirements are correcly installed, by running both the commands "``go version``" and "``protoc --version``"
2. Clone the repository
3. In the root folder of the repository, run "``go mod tidy``"
4. If you are using vscode, start the project with all its resources by running the "Start program" command in the _"run and debug"_ menu. Otherwise, you can start the project by running "``go build ./theHive/app/; go run ./theHive/app/; go build ./walle/app; go run ./walle/app/;``" from the root folder of the repo. Then you can navigate to [http://localhost:8000/](http://localhost:8000/) to view the GUI



## Authors

Dante Wesslund <br>
Victor Hernadi <br>
Martin Mörsell

## License

TBD
