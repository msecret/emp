EMP
=====
EMP is a fully encrypted, distributed messaging service designed with speed in mind.
Originally based off of BitMessage, EMP makes modifications to the API to include
both Read Receipts that Purge the network of read messages, and an extra identification field
to prevent clients from having to decrypt every single incoming message.

Submodules
----------

This repository contains submodules.  You will need to run:
```
git submodule init
git submodule update
```

Required Libraries
---------
In order to compile and run this software, you will need:

* The [Go Compiler (gc)](http://golang.org/doc/install)
* For Downloading Dependencies: [Git](http://git-scm.com/book/en/Getting-Started-Installing-Git)
* For Downloading Dependencies: [Mercurial](http://mercurial.selenic.com/wiki/Download)

Building and Launching
---------

* `make build` will install the daemon to ./bin/emp
* `make start` will set up the config directory at ~/.config/emp/, then build and run the daemon, outputting to the log file at ~/.config/emp/log/log_<date>
* `make stop` will stop any existing emp daemon
* `make clean` will remove all build packages and log files
* `make clobber` will also remove all the dependency sources

**Running as root user is NOT recommended!**

Configuration
---------
All configuration is found in `~/.config/emp/msg.conf`, which is installed automatically with `make start`. An example is found in `./script/msg.conf.example`. The example should be good for most users, but if you plan on running a "backbone" node, make sure to add your external IP to msg.conf in order to have it circulated around the network.