PROG=	gotime

all: gotime

gotime: *.go go.mod
	go build -o ${PROG}

install:
	install -o root -g bin ${PROG} /usr/local/bin

clean:
	rm -f ${PROG}

user:
	doas user add -c 'daytime daemon' \
            -d /var/empty \
            -p * \
            -s /sbin/nologin \
            _daytimed
