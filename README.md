# go-config
Just a very small go package for loading values from .ini files, I am just playing around with GoLang so excuse me if you dont like the code :-)

## With config.Load() you can load ini files

first param is a required string array of required ini settings

second param is not required and can be any number of ini file names to load (NOTE: if you provide no file names then it will look for an ini with the name of the executable plus ".ini")




## With config.Get() you can get individual settings

param is name of setting, ie. if you have this in your ini file

SETTING=SOME VALUE

"SETTING" is the param you send to config.Get() go get the value.



