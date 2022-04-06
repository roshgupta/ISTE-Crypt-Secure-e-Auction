## Secure-e-Auction

A Secure e-Auction using Hyperleadger Fabric and Beego Framework
---
This Project was done under the <b>Crypt</b> Special Interest Group of <a href="https://iste.nitk.ac.in/">ISTE-NITK</a>

The Project was done entirely in <b>golang</b> and its various libraries<br>
---

### To get Hyperleadger Fabric Running
1.Clone Repo
2.Go to "Website" folder
3.Open Terminal
4.cd auction
5.cd test-network
6../network.sh up createChannel -ca
7../network.sh deployCC -ccn auction -ccp ../auction-simple/chaincode-go/ -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')"
8.cd ..
9.cd auction-simple/application-javascript
10.npm install
11.node enrollAdmin.js org1
12.node enrollAdmin.js org2

### To get the Website Running
1.Go to "Website" Folder
2.Open Terminal
3.go mod tidy
4.go run main.go

### To end the Hyperleadger Fabric Blockchain
1.Go to Simple Auction Folder
2.Open Terminal
3.cd auction-simple/application-javascript
3.rm -rf wallet
4.cd ..
5.cd ..
6.cd test-network
7../network.sh down
