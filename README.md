# hack-salesforce
## Discription
hack-salesforce input daily works to Attendance Sheet in salesforce.

## Setup
### Install Golang
http://golang.jp/install

### Install ChromeDriver
Install [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/downloads) and set Path.

### Make build
```bash
$ make build
```

## How to use
```bash
$ hack-salesforce --config_path {config_path} --jsonfile {jsonfile}
```

* config_path
  * Set configuration with reference to the `template_config.ini`

* jsonfile
  * Set Json file for a month's attendance with reference to the `template_works.json`
