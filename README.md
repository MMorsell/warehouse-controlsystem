# walle

Basic project "flowchart"

[![Build Status](https://dev.azure.com/martinmorsell18/Current%20Projects/_apis/build/status/Gophers.walle?branchName=refs%2Fpull%2F1%2Fmerge)](https://dev.azure.com/martinmorsell18/Current%20Projects/_build/latest?definitionId=8&branchName=refs%2Fpull%2F1%2Fmerge)

![Board](/projectStructure.png)

# Project Description:
The Wall-e project is a control system for robots proceeding tasks in a 2d grid, Walle-e is inspired by storage 
facilities using robots to perform simple storage tasks.

Wall-e consists of the 2 programs "the hive" and "robot", which will communicate via a gRPC system.

The hive is a central hub, essentially being the center of all operations and communication. It provides orders for the robot, which is in the form of "go to these (x,y) coordinates by taking this specific path.\
\
The robot will move to these coordinates by a received path from the hive once the robot has been given confirmation that the path is available, meaning there is no risk of a collision with another robot. Once the robot has arrived at its desired spot the task is completed and the robot can be given further orders.

