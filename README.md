# groupie-tracker
[![Made with Go](https://img.shields.io/badge/Go-1-blue?logo=go&logoColor=white)](https://golang.org "Go to Go homepage")  [![Made with Docker](https://img.shields.io/badge/Made_with-Docker-blue?logo=docker&logoColor=white)](https://www.docker.com/ "Go to Docker homepage")

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

It will be given an API, that consists in four parts:

- The first one, artists, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

- The second one, locations, consists in their last and/or upcoming concert locations.

- The third one, dates, consists in their last and/or upcoming concert dates.

- And the last one, relation, does the link between all the other parts, artists, dates and locations.

## Installation

To install the program into your own computer in order to run it
```shell
$ git clone https://zone01normandie.org/git/scointin/groupie-tracker-visualizations.git
```

### Installation using docker
This project can be run without any hassle by running dockerize.sh
````shell
$ chmod +x dockerize.sh
$ ./dockerize.sh
````
You will be prompted unto installation, uninstallation or cancel if you ever change your mind.


## Usage

If you used the shell script (dockerize.sh) everything should be working just fine,
you can visit this url : http://localhost:5826  

To launch the server :
```shell
$ go run .
```
Then go to http://localhost:5826

## Collaborators

- A. Nassuif (Back dev and head of project)
- S. Cointin (Front & back dev)
- M. Soumare (Front dev)


<a href="https://www.digitalocean.com/?refcode=d52bcc90ccc2&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge"><img src="https://web-platforms.sfo2.cdn.digitaloceanspaces.com/WWW/Badge%201.svg" alt="DigitalOcean Referral Badge" /></a>
