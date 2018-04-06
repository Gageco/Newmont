import requests
from requests.auth import HTTPDigestAuth
import json
import os
import time
# import Adafruit_DHT

onRaspberryPi = False

sensorData = {'temperature': 1.0, 'humidity': 1.0}

def getTempData():
    print('Getting temperature data')
    # Be sure to use DHT22 Temperature Sensor
    # sensor = Adafruit_DHT.DHT22

    # Change Pin as necessary
    pin = 23

    # Try and get sensor reading, will retry
    # humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
    humidity = 100
    temperature = 100
    # If no reading is recieved try again
    if True:#  humidity is not None and temperature is not None:
        print('Temp={0:0.1f}*C  Humidity={1:0.1f}%'.format(temperature, humidity))
        sensorData['temperature'] = temperature
        sensorData['humidity'] = humidity

    else:
        print('Failed to get reading. Try again!')

def writeDataToFile():
    print('Writing Data to File')

    file = open("./data.txt", "w")
    file.write("{\"temperature\": " + str(sensorData['temperature']) + ", \"humidity\": " + str(sensorData['humidity']) + "}")
    file.close()

def uploadToSia():
    print('Uploading to Sia')

    if onRaspberryPi:
        url = "http://localhost:9980/daemon/version"
        data = '{}'
        response = requests.get(url, data=data,headers={"User-Agent":"Sia-Agent"})
        resp = response.json()

    else:
        cwd = os.getcwd()
        url = "http://localhost:9980/renter/upload/Newmont/data.txt"
        # source = cwd + "data.txt"
        data = '{"datapieces": 1, "paritypieces": 1, "source": "data.txt"}'
        response = requests.post(url, data=data,headers={"User-Agent":"Sia-Agent"})
        resp = response.json()

    # Do some error handling here

def checkUpload():
    print('Checking upload')
    data = '{}'
    url = "http://localhost:9980/renter/files"
    response = requests.get(url, data=data,headers={"User-Agent":"Sia-Agent"})
    resp = response.json()

    if resp["files"] != None:
        while resp["uploadprogress"] < 100:
            time.sleep(5)
            print('Progress: ' + str(resp["uploadprogress"]))
            checkUpload()
    else:
        print("No files uploading")


x=0
while x < 1:
    x=1
    getTempData()
    writeDataToFile()
    uploadToSia()
    checkUpload()

    # time.sleep(300) # wait for 5 minutes
