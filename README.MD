## Secure-e-Auction

A Secure e-Auction using Hyperleadger Fabric and Beego Framework
---
This Project was done under the <b>Crypt</b> Special Interest Group of <a href="https://iste.nitk.ac.in/">ISTE-NITK</a>

The Project was done entirely in <b>golang</b> and its various libraries<br>
---

### To get Hyperleadger Fabric Running
<ul>
<li>Clone Repo
<li>Go to "Website" folder
<li>Open Terminal
<li>cd auction
<li>cd test-network
<li>./network.sh up createChannel -ca
<li>./network.sh deployCC -ccn auction -ccp ../auction-simple/chaincode-go/ -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')"
<li>cd ..
<li>cd auction-simple/application-javascript
<li>npm install
<li>node enrollAdmin.js org1
<li>node enrollAdmin.js org2
<\ul>

### To get the Website Running
<ul>
<li>Go to "Website" Folder
<li>Open Terminal
<li>go mod tidy
<li>go run main.go
<\ul>
  
### To end the Hyperleadger Fabric Blockchain
<ul>
<li>Go to Simple Auction Folder
<li>Open Terminal
<li>cd auction-simple/application-javascript
<li>rm -rf wallet
<li>cd ..
<li>cd ..
<li>cd test-network
<li>./network.sh down
<\ul>
