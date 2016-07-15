## Setup 

* clone the repo 
* create a free account at [papertrail](https://papertrailapp.com/) 
* After signup, click on `Add your First System` 
* Modify the port specified in `syslog.go` to what is displayed. 

## Running the source 
* `go run syslog.go` 
* Back in papertrail, click on events to see logs streaming in. 
  * If the protocol specified in `syslog.go`is `tcp` then change that to `udp` to get it working. 

#### TCP

To get TCP working: 

* `git fetch`
* `git checkout tls` 
* `go run syslog.go` 
