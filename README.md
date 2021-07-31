# Newmont

edit: README.md updated 2021-07-31 for clarity

The platform and program that was written for the Newmont mining challenge, held by Colorado School of Mines.

The important files are in the 'sia' folder and it contains three different programs, the backend, middle, and front end of the platform developed.

## Purpose
This program is built in three parts. The RaspberryPi that collects data that it shares with Google Cloud Platform, that then the front end utilizes for a temperature collection from a location. It uses three different languages, based on what would be best for the hardware. The front end uses json and html, and so on. Middle end uses Google Cloud Platform and the Google developed language Go, and the backend uses Python, easily implemented on the Rapsberry Pi.

## sia/site - Front End
This is the front end of the whole set up. it includes the csv that contains that data pulled from Google Cloud Platform periodically.

## GCP - Middle
This is the Google Cloud Platform program written in Go. It is a simple program that is more or less just a basic REST server with authentication for the backend Raspberry Pi to talk to. It accepts structs with temperature data in json, and then passes it along when requested.

## raspberryPi - Back End
The Raspberry Pi collected and used GCP to catalog the data so that the front end can display it.
