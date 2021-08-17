# test-api

This is a very crude http server that expects to receive a simple json at the root endpoint ("/"); the json is expected to have a single key, "text", alongside its value which should be comprised of a sentence. The server parses the text input and performs a few changes on the, then returns the formatted text back to the caller.


## Known issues
* in some situations, the code may remove certain punctuations; e.g. "I have worked at ING." becomes "I have worked at ING"
* It's assumed that when starting a new sentence, you will always have a space. "I have an account at ABN.Rabo is not my bank of choice" will result in neither of the strings being parsed.
