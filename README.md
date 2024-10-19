# go-config
Just a very small go package, I am just playing around with GoLang.

With config.Load() you can load ini files

first param is a required string array of required ini settings

second param is not required and can be any number of ini file names to load




With config.Get() you can get individual settings

param is name of setting, ie. if you have this in your ini file

SETTING=SOME VALUE

"SETTING" is the param you send to config.Get() go get the value.



