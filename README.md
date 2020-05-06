# myhttp

Simple tool to make http requests in parallel and print MD5 hash of the response body

## Description

- The tool accepts websites as inputs and make http requests, if request is successful, prints the
   sitename allong with MD5 hash of the response body.
- Flag ```-parallel``` is used to set how many requests can be processed parallely for the given input websites.
   e.g. if the flag is set to 2, 2 website request will be processed parallely(default value is 10, if the flag is not provided).

## Usage

- Clone the repository using git clone.
- Build the project using command ```go build```
- Run the program using command ```./myhttp -parallel <no of parallel sites to be processed> <list of sites separated by space>```
- To run, this alternate command also can be used- ```./myhttp -parallel=<no of parallel sites to be processed> <list of sites separated```
- e.g. ```./myhttp -parallel 3 google.com www.fb.com http://yahoo.com```
- To run the test cases, run command- ```go test ./...```

## Assumptions

- The website are fetched using GET request, so success will be checked with http status code 200(OK).
- If any website fetching is failed, the program will keep running untill other websites are finished.




