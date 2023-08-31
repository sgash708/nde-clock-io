# nde-clock-io

automate clock in/out for a certain site....

## TODO

```bash
$ cp config.sample.yml config.yml
$ vi config.yml
##### startline ######
Secret:
  URL: https://example.com
  UserID: 1111111
  Password: password
  Channel: test-channel
  Token: xoxb-xxxxxxx-xxxxx
##### endline ######
```

## Require
- Chromedriver
  - [check your version](https://googlechromelabs.github.io/chrome-for-testing/)
- Slack
  - channel
  - token

### linux
```bash
$ CHROMEDRIVER_VERSION=`curl -sS chromedriver.storage.googleapis.com/LATEST_RELEASE`
$ curl -sS -o /tmp/chromedriver_linux64.zip http://chromedriver.storage.googleapis.com/$CHROMEDRIVER_VERSION/chromedriver_linux64.zip
$ unzip /tmp/chromedriver_linux64.zip
$ mv chromedriver /usr/local/bin/
```

https://chromedriver.chromium.org/downloads

### mac
```bash
$ brew install chromedriver
```

## Run

To run `ndeio` binaries, you need `config.yml` in the same directory.

```bash
$ make m-gobuild
$ go install ./bin/ndeio
$ ndeio clockin
```

or
```go
$ make gobuild
$ cd bin
$ ./ndeio clockin
```

### crontab
in ubuntu server...

```
$ cd
$ wget https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv
$ crontab -e

# clockin
00 8 * * 1-5 bash -c "sleep $((RANDOM \% 1800))s"; grep `date "+\%Y/\%-m/\%-d"`, syukujitsu.csv > /dev/null || ./ndeio clockin
# clockout
00 18 * * 1-5 bash -c "sleep $((RANDOM \% 1800))s"; grep `date "+\%Y/\%-m/\%-d"`, syukujitsu.csv > /dev/null || ./ndeio clockout
```
