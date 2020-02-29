# WhatsAppAuto 💬

A Proof of Concept program which uses WhatsApp web services (by making request to web.whatsapp), and perform the required operations
such as authenticating the user and sending the messages on user behalf.

A ```.csv``` file is used to provide the data, and the required contact info. Messages are constructed using the data provided in the ```.csv``` file.
Columns in data table are :
| Name |
|---|
|Student Name|
|Contact Info|
|Timing|
|Date|
|Subject Code|

Since this is just a proof of concept program the original application is deployed in a DigitalOcean Droplet with supervisor.
