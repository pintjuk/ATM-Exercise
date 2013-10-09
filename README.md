 # format
	
	| Opcode	| Arguments						|
	| xxxxxxxx	| 9 bits ore variable leangth	|

# Update Add

	| 0x00	| Message	|
	| 8 bits| aligned to 8 	|


## Message  
Message starts with a sequens of UTF-8 characters terminated by 0xFF 0x00 and is followed by a Languagepack

	| Add message in UTF-8 	| 0xFF 0x00	| Languagepack	|
	| aligned to 8 bits 	| 16 bits 	| aligened to 8 bit 	|

### Languagepack
Languagepack is an list of an arbitrary amount of ellements

ellement:
	| Reffnumber 	| 0xFF	0x7F	| Text			| 0xFF 0x01 	|
	| aligned to 8 	| 16 bits	| aligend to 8 bits 	| 16 bits 	|




# 10 byte Comands 
## Send integer data

	0x01 X Y Y Y Y Y Y Y Y 
	
Xses are ether 2 32 bit integers ore one 64 bit ingteger

	0x02 Y Y Y Y Y Y Y Y Y
Ys are 
	

## Print to client screen

	
	W X X X X X X X X X
	
	how W is interpreted:
	1 x x x x x x x = print command
	x 1 x x x x x x = flag to tell the client to clear the scrin before printing
	x x 1 x x x x x = flgg to request data (integer if not ascii)
	x x 1 1 x x x x = flgg to request ASCII data 
	
	Ys make up an integer ref to the string in the languagepack to be printed

# Senario
	port A					port B
	C	S				C	S
	|<-	| print scrien			|<-	| Update Client
	|	| clear screen and 		|	|	
	| 	| request integer input 	|	|
