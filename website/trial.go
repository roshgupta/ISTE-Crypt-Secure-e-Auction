package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func proposeBid(app, arg0, arg1, arg2, arg3, arg4 string) (toprint string, err error) {
	fmt.Println(app, arg0, arg1, arg2, arg3, arg4)
	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	stdout, err := cmd.Output()
	return string(stdout), err
}

func addBid(app, arg0, arg1, arg2, arg3, arg4 string) (toprint string, err error) {
	fmt.Println(app, arg0, arg1, arg2, arg3, arg4)
	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	stdout, err := cmd.Output()
	return string(stdout), err
}

func main() {
	bidder_user_id := 1
	amount := 2500
	app := "node"
	arg0 := "auction/auction-simple/application-javascript/bid_web.js"
	arg1 := "org2"
	arg2 := "bidder" + strconv.Itoa(bidder_user_id)
	arg3 := "CarAuction1"
	arg4 := strconv.Itoa(amount)
	fmt.Println(app, arg0, arg1, arg2, arg3, arg4)
	toprint, err := proposeBid(app, arg0, arg1, arg2, arg3, arg4)
	bidId := strings.Split(toprint, "\n")
	bidId = bidId[:len(bidId)-1]
	fmt.Println(toprint)
	fmt.Println(bidId)
	fmt.Println(err)

	app = "node"
	arg0 = "auction/auction-simple/application-javascript/submitBid.js"
	arg1 = "org2"
	arg2 = "bidder" + strconv.Itoa(bidder_user_id)
	arg3 = "CarAuction1"
	arg4 = bidId[0]
	toprint, err = addBid(app, arg0, arg1, arg2, arg3, arg4)
	print(err)
	print(toprint)
}
