# GoCoronaVirusAPIParser
Go program that receives JSON from corona virus API at https://covid-api.mmediagroup.fr/v1/cases and parses the JSON for global and state info

# INSTALL
To install simply issue the go build command as all code is native to Golang and requires no
external dependencies. go build will link state-struct.go and main.go together automatically.

cd src

go build -o anyname

# Program state
Please note this software is working , but still needs some more work.
As of this moment, it will give you Global Corona Virus information and Tennessee information

you may change the tennessee code section of the program to match your US state.


# Special Thanks
Thank you to the people who created the public API for tracking the Corona Virus !

Their API can be reached at https://covid-api.mmediagroup.fr/v1/cases
