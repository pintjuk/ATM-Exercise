# format

| Opcode	| Arguments						|
| xxxxxxxx	| 9 bits ore variable leangth	|

# Update Add

| 0x00	| Message 		  |
| 8 bits| aligned to 8 	  |


## Message  
Message starts with a sequens of UTF-8 characters terminated by 0xFF 0x00

| Add message in UTF-8 	| 0xFF 0x00	| Languagemap 			|
| aligned to 8 bits 	| 16 bits 	| aligened to 8 bit 	|

### Language Map
Language Map is an list of an arbitrary amount of ellements


Map ellement:
| Reffnumber	| Text 				| 0xFF 0x00 |
| 2 byte	 	| aligend to 8 bits | 16 bits 	|


# Print
	| 0x01		| UTF-8 data	|
	| xxxxxxxx	| 9 byte 		|

# 

## Message 
UTF-8 encoded html

#Vad behöver jag för operationer?

## Client
* Send integer data	

## Server
* Update Add
* Print to the client Screen, inserting real time data from server
* Clear Client screen 
* Request integer input data

Senario

	C 	S

