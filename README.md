GOTIME(8) - System Manager's Manual

# NAME

**gotime** - simple Daytime Protocol daemon written in go

# SYNOPSIS

**gotime**
\[**-d**]

# DESCRIPTION

**gotime**
is a privilege-dropping server that implements the Daytime Protocol as specified in
RFC 867.
It responds to TCP connections on port
*13*,
returns a human-readable date string to the client,
then closes the connection.
Time portion of the date string is reported in UTC.

The options are as follows:

**-d**

> Debug mode.
> **gotime**
> will run in the foreground and output debug messages to stdout.
> root privileges are not required for startup and the server will listen on
> *localhost*
> port
> *13013*.

# BUILD AND INSTALL

To build
**gotime**:

	$ make

To clean things up:

	$ make clean

To install, first create the necessary user account,
then install the program and the
rc.d(8)
script:

	# make user
	# make install

To hook
**gotime**
into the
rc(8)
system, simply add
**gotime**
to the
*pkg\_scripts*
variable in
rc.conf(8):

	pkg_scripts=gotime

# EXAMPLES

Request current date and time from a local
**gotime**
instance running in debug mode:

	$ nc localhost 13013

Make a request to a remote server running
**gotime**:

	$ nc puffy.example.com 13

# SEE ALSO

rc(8),
rc.conf(8),
rc.d(8),

[daytimed(8)](https://github.com/sbennett1990/daytimed)

# AUTHORS

Scott Bennett

# CAVEATS

**gotime**
has only been built and tested on
OpenBSD 6.8
and newer, because it relies on
pledge(2)
and the
rc(8)
system.

**gotime**
only listens on IPv4 addresses.

OpenBSD 6.8 - June 24, 2021
