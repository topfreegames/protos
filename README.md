TFGCo Go Protobuf definitions
=============================

This package contains the common .proto files and related generated code we use in our projects.


### Changelog

#### v1.6.2
* Maestro
  * Added "roomType" to Room 
  * Added "SendRoomResync" to GRPCForwarder 

#### v1.6.1
* Events Gateway
  * Added "retry" to SendEventsRequest 

#### v1.6.0
* Events Gateway
  * `make eventsgateway` to build proto and mock
  * Renamed some types
  * Added SendEvents 
* Maestro
  * `make maestro` to build proto

#### v1.5.0

#### v1.4.4
* Events Gateway
  * Changed package

#### v1.4.3
* Events Gateway
  * Added mock file

#### v1.4.2
* Events Gateway
  * Empty response

#### v1.4.1
* Events Gateway
  * Added timestamp in Event

#### v1.4.0

* Events Gateway
  * Grpc events proto and golang generated file

#### v1.3.0

* Maestro
  * Added tags in SendRoomInfo

#### v1.2.0

* Maestro
  * Added SendRoomPing

#### v1.1.1

* Maestro
  * Fixed missing game in SendRoomInfo
#### v1.1.0

* Maestro
  * Added SendRoomInfo
  * Added SendRoomEvent

#### v1.0.0

* Maestro
  * Grpc events proto and golang generated file

### TODOs

[ ] Add instructions on how to run the example server
