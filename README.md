# Ports and initiation 
port 9000 is the talk-port, port 9001 is the update-port.

The first thing sent ovver the update port must be 32 bit integer from server to the cliant.

The first thing sent ovver the the talk port must start by a 32 bit enteger cliant sends to server, it must be the same value the client resived from the server ovver the update port. This integer vill be revered to as the client ID.

# Update Add
update add updates the Advertisment and the language package

	| 0x00	| Message	|
	| 8 bits| aligned to 8 	|


## Message  
Message starts with a sequens of UTF-8 characters terminated by ETB(0x17) and is followed by a Languagepack

	| Add message in UTF-8 	| 0x17		| Languagepack	|
	| aligned to 8 bits 	| 8 bits 	| aligened to 8 bit 	|

### Languagepack
Languagepack is an list of an arbitrary amount of ellements

ellement:
	| Reffnumber 	| Text			| 0x17	 	|
	| 8 bits 	| aligend to 8 bits 	| 8 bits 	|
element contains an refference number witch is an 8 bit unsigned integer folowed by an UTF-8 string terminated by a EOT(0x17) byte



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
	port 9000				port 9001
	C	S				C	S
	|<-	| S sentding client id		|	| 
	|<-	| S sending UpdateAdd		|->	| confirming client id
	|	|				|	|
	| 	|				|<-	| server sends 
	|	| standing by			|	| "Print to client screen comand" 
	|	|				| 	| with arguments ponter homescreen
	|	|				|	|		
	|	| standing by 			|	| standing by
	|	|				|	|
	|	|				|	| user chuses to change language
	|	|				|->	| client sends a "send integer data package"
	|	|				|	|
	|<-	| server sends updateAdd	|	| server runs logic
	|	|				|	|
	|	|				|<-	| server sends a new 
	|	|				|	| print to client screen comand
	|	|				|	| with argument to print homescreen
	|	|				|	|
	|	|				|	| user chuses to login 
	|	|				|->	| client sends a "send integer data" package to server
	|	|				|	| 
	|	|				|	| server runs some logic
	|	|				|	| 
	|	|				|<-	| server sends a "print to client screen" package 
	|	|				|	| with arguments to print login screen
	|	|				|	| 
	|	|				|	| user enters account number
	|	|				|->	| client sends "send integer data" packag with the account number 
	|	|				|	|
	|	|				|<-	| server sends "print to user screen" package 
	|	|				|	| with the argument to print password page
	|	|				| 	| user enters the password 
	|	|				|->	| client sends pasword with "send ASCII data package"
	|	|				|	|
	|	|				|	| server runs logic and decides the pasword was corect
	|	|				|<-	| server sends "Print to client screen" package
	|	|				|	| with the balance as argument and arguments to printe saldoscreen
	|	|				|<-	| server sends "Print to client screen "
	|	|				|	| with client name as argument and arguments to print Nameblock
	|	|				|	| 
	|	|				|->	| user chuses to withdraw and the client sends "integer data" package with the amount withdrown
	|	|				|	| 
	|	|				|	| server runs some logic, withdraws the amount and logs out the user
	|	|				|<-	| server sends "print to user screen packag" with arguments to print initial screen
	
	
# Client State machine

All the logic is ment to run on the server and every implementatin of the client 
is suposed to be a simple state machine so that the client dosent ever need to be updated.

this is how the clien shoud work 


-------
|start|
-------
   |                                       _____________________________________________________________
   |  					  /	  _________________________	  		 	|
   |					 /	 /                         |				|
   \/					\/	\/ 		           |				|
------------------------------	   -----------------------------------	   |				|
|initilize conection to servr|---->|wait for print message from server|____| if server 			|
------------------------------   ->|and print acordingly              |      dosent request 		|
				|   -----------------------------------	     dosent request data 	|
				|	| server requests  |server requests 				|
				|	| integer data     |Ascii data					|
			        |	|		   |						|
			        |   	|		   |						|
			        |       \/                 \/                 				/
		          -----------------------------    ----------------------------		       /
		          |ask user for integer data  |    | ask user for ascii data  |_______________/
		          |send integer data to server|    | send ascii data to server|
		          -----------------------------    ----------------------------
