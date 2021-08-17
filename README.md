# test-api

This is a very crude http server that expects to receive a simple json at the root endpoint ("/"); the json is expected to have a single key, "text", alongside its value which should be comprised of a sentence. The server parses the text input and performs a few changes on the, then returns the formatted text back to the caller.

## Building

There is a Dockerfile supplied, which can be used to build the image and run it locally (make sure to expose the correct ports)

## Testing
This is a pure API and offers no http webpage. You should use curl to test, see an example below (and feel free to provide your own text):

```
curl -X POST ec2-3-248-170-100.eu-west-1.compute.amazonaws.com:8080 -H "Content-Type: application/json" --data '{"text": "The analysts of Abn did a great job! rabo Rabo rAbO VolkSBANK volksbanK TrioDOS."}'
```

The app is deployed in AWS, on an EC2 instance, the URL above will work for as long as this repository will remain public.

## Known issues
* in some situations, the code may remove certain punctuations; e.g. "I have worked at ING." becomes "I have worked at ING"
* It's assumed that when starting a new sentence, you will always have a space. "I have an account at ABN.Rabo is not my bank of choice" will result in neither of the strings being parsed. "I have an account at ABN. Rabo is not my bank of choice", however, will  work.


