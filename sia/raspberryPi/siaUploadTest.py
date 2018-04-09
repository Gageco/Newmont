import requests
import json
import os
import time
from datetime import datetime
# import Adafruit_DHT

onRaspberryPi = False

wallet = {'password': ''}



sensorData = {'time': '','temperature': 10.0, 'humidity': 5.0}

def getSensorData():
    print('Getting Sensor Data')
    # Be sure to use DHT22 Temperature Sensor
    # sensor = Adafruit_DHT.DHT22

    # Change Pin as necessary
    # pin = 23

    # Try and get sensor reading, will retry
    # humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
    humidity = humidity + 2
    temperature = temperature + 1

    #get current time
    sensorData['time'] = datetime.now().strftime('%Y-%m-%d %H:%M').replace(" ", "-")

    if True: #humidity is not None and temperature is not None:
        print('Temp={0:0.1f}*C  Humidity={1:0.1f}%'.format(temperature, humidity))
        sensorData['temperature'] = temperature
        sensorData['humidity'] = humidity

    else:
        print('Failed to get reading. Try again!')
        time.sleep(5)
        getSensorData()

def writeDataToFile():
    print('Writing Data to File')

    file = open("./data.txt", "w")
    file.write("{\"temperature\": \"" + str(sensorData['temperature']) + "\", \"humidity\": \"" + str(sensorData['humidity']) + "\", \"time\": \"" + str(sensorData['time']) + "\"}")
    file.close()

def getWalletPassword():
    print('Getting Wallet Password')
    wallet['password'] = open("./password.txt").read()
    print('Password: ' + wallet['password'])

def checkNetwork():
    print('Checking network status')
    url = "http://localhost:9980/consensus"
    data = '{}'
    response = requests.get(url, data=data,headers={"User-Agent":"Sia-Agent"})
    resp = response.json()
    # print(resp)
    if resp['synced']:
        print('Sia Network Synced')
    else:
        print('Sia Network Not Synced')

def unlockWallet():


    url = "http://localhost:9980/wallet/unlock"
    response = requests.post(url, data={"encryptionpassword": wallet['password']},headers={"User-Agent":"Sia-Agent"})
    resp = response.json()

    if resp['message'] == 'error when calling /wallet/unlock: wallet has already been unlocked':
        print('Wallet Already Unlocked')
    elif resp['message'] == 'error when calling /wallet/unlock: provided encryption key is incorrect':
        print('Password Incorrect')
    else:
        print(resp)

def uploadToSia():
    print('Uploading to Sia')

    cwd = os.getcwd()
    url = "http://localhost:9980/renter/upload/data.txt"
    source = cwd + "/data.txt"
    response = requests.post(url, data={"datapieces":2, "paritypieces":12,"source":source},headers={"User-Agent":"Sia-Agent"})
    # print response.status_code
    if response.status_code == 204:
        print('File Uploading')
    elif response.json()['message'] == 'upload failed: a file already exists at that location':
        print('File Already Exists')
    else:
        print response.json()

def checkUpload():
    data = '{}'
    url = "http://localhost:9980/renter/files"
    response = requests.get(url, data=data,headers={"User-Agent":"Sia-Agent"})
    resp = response.json()
    # print resp
    if resp['files'] != None:
        # print resp
        if resp['files'][0]['uploadprogress'] < 100:
            time.sleep(1)
            print('Progress: ' + str(resp['files'][0]['uploadprogress']))
            checkUpload()
        elif resp['files'][0]['uploadprogress'] >= 100:
            print('File Uploaded')
        else:
            print("No files uploading")

x=0
while x < 5:
    x=1
    getSensorData()
    writeDataToFile()
    checkNetwork()
    # getWalletPassword()
    # unlockWallet()
    uploadToSia()
    print('Checking Upload')
    checkUpload()

    time.sleep(10) # wait for 150 second
