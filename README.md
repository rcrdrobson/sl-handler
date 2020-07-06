## Running the SL-Handler

### Dependencies
Docker - https://docs.docker.com/get-docker/
Go - https://golang.org/dl/

### Run the binary file

You can directly run the binary file, or compile and then run.

To run directly run the main file in the **/src** folder use the command: **./main**.

To compile and run manually follow the steps below.

### Download the necessary SL-Handler modules 
Run:

**go get github.com/ricardorobson/sl-handler/src/database**

and after

**go get github.com/ricardorobson/sl-handler/src/docker**

### Build
Before running main.go you need to compile. To do this, with Golang installed, go to the **/src** directory and run the command: **go build main.go**.

This command will create the **main** binary.

### Running
To run the SL-Handler, after creating your torque in the previous step, run the following command or execute the binary file:
**./main**.

## Contribute and license
**sl-handler** was created by **Ricardo Robson**, you can follow him on [Twitter](https://twitter.com/rcrdrobson) or [Facebook](https://www.facebook.com/rcrdrobson) for updates and other projects!

**sl-handler** use [ISC Lincense](https://en.wikipedia.org/wiki/ISC_license). Feel free to fork **sl-handler** on GitHub if you have any features that you want to add!

You can contact me by email too: rcrdrobson@gmail.com
