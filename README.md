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

	 W X X X X X X X X X
	w = 0x01 
	folowing 8 bytes are one 64 bit encoded unsigned integer
	w = 0x02
	folowing 8 bytes are one double
	W = 0x4 
	folowing 4 bytes are one signed integer
	w = 0x5
	folowing 9 bytes are 9 ASCII characters
	w = 0x6 


## Print to client screen

	
	W Y X X X X X X X X
	
	how W is interpreted:
	1 x x x x x x x = print command
	x x x x 1 x x x = flag to tell the client to clear the scrin before printing
	x 1 0 0 x x x x = flgg to request data (integer if not ascii)
	x 1 1 0 x x x x = flgg to request ASCII data 
	x 1 1 1 x x x x = flag to request login data

	Y is the ref to the string to be printed
	
	Ys make up an integer ref to the string in the languagepack to be printed

# Senario
	port A					port B
	C	S				C	S
	|<-	| print scrien			|<-	| Update Client
	|	| clear screen and 		|	|	
	| 	| request integer input 	|	|
