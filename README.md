This projects solves 2 puzzles.  The first public function merges 2 sets of sorted arrays.  The second function finds the largest int32 product that can be generated with an array of integers.


To run the tests:

1) Install Golang: https://golang.org/doc/install

2) Create a directory.  I called my workspace.

3) Create a directory called "src" inside of that directory

4) Clone this project into the src folder

Your workspace should look like this:

    -workspace

        - src

            -codepuzzle
    
   
5) Set the gopath: export GOPATH=absolute/path/to/workspace
 
6) from the codepuzzle directory, run:"go get -t ./ ..."

This should download the dependencies to run the tests to the src directory.

7) from the codepuzzle dir run "go test".

   
