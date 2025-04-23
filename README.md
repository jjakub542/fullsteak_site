# FullSteak

Website (with blog) in Go Echo + postgres + redis
All important commands are in Makefile (with description)
Docker is required to run databases
-----
### Install go migration cli tool:
```
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```