# mqtt-server

**Package server implements basic functionality for launching an [MQTT 3.1.1](http://docs.oasis-open.org/mqtt/mqtt/v3.1.1/) server that use multiple listeners and protocols.**

Installation
=============

Get it using go's standard toolset:

```bash
mkdir -p $GOPATH/src/github.com/adminbaintex
mv mqtt-server $GOPATH/src/github.com/adminbaintex
# import "github.com/adminbaintex/mqtt-server/server"
```


Dependencies 
=============

Dependencies are managed with govendor.

* Initialize "vendor" directory
```
govendor init
```

* List packages used in the application
```
govendor list
```

* Add external packages in GOPATH to vendor folder
```
govendor add +external
```
