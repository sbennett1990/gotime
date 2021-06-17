all: gotime

gotime: *.go go.mod
	go build -o gotime

clean:
	rm -f gotime

user:
	doas user add -c 'daytime daemon' \
            -d /var/empty \
            -p * \
            -s /sbin/nologin \
            _daytimed
