hatter
======

An attempt at writing a simple sysvinit replacement in Go. I'm writing
this primarily to see if I can. It is not intended to ever be fully
usable, although it likely will be eventually if it actually works.

Goals
-----

- [ ] runit-inspired simplicity combined with some useful systemd-style features.
	- [ ] Controlled through file creation. Simple control utility will be provided.
	- [ ] Works as a cron replacement.
	- [ ] Handles mounts.
	- [ ] Configuration-free services for simple cases. A service can be a single executable.
- [ ] No monopolization of `/etc`. All files are stored in a single directory tree.
- [ ] Very small.
- [ ] Completely static and dependency free.
- [ ] Heavy differentiation between services, cronjobs, and mounts.

Possible Features
-----------------

- [ ] User session support.
