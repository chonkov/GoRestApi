# GO REST API
## Summary:
- Simple rest api written in go that accepts a shell command and returns the output of that command.
The task was part of the interview process for "EnableIT"

## Description:

- Write a simple GO HTTP REST API that will accept a shell command and return the output of that command.
- Create an endpoint with POST method. api/cmd POST. Accept the command via query param or JSON body. Return the output 
of the command as a response.
- Plus point: Return error if command is not found, with proper status code.