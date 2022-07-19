## Conventional package structure for golang 
- cmd: entry point to the application 
- pkg: shareable logic that could be imported 
- internal: code explicitly for the software, this will contain routing logic, middlewares, entities, validation logic etc. 


## External Configuration 
- check out the `viper` library for working with environment variables and external configuration 
- 