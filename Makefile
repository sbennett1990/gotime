PROG=	gotime
MAN=	gotime.8

all: gotime

gotime: *.go go.mod
	go build -o ${PROG}

md:
	mandoc -Tmarkdown ${MAN} > README.md

install:
	install -o root -g bin ${PROG} /usr/local/bin
	install -o root -g bin ${PROG}.rc /etc/rc.d/${PROG}

clean:
	rm -f ${PROG}

user:
	doas user add -c"daytime daemon" \
            -d/var/empty \
            -p* \
            -s/sbin/nologin \
            _daytimed
