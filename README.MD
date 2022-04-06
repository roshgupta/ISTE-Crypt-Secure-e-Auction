## Secure-e-Auction

A Secure e-Auction using Hyperleadger Fabric and Beego Framework
---
This Project was done under the <b>Crypt</b> Special Interest Group of <a href="https://iste.nitk.ac.in/">ISTE-NITK</a>

The Project was done entirely in <b>golang</b> and its various libraries<br>
---

### To get Hyperleadger Fabric Running
<ul>
<li>Clone Repo<\li>
<li>Go to "Website" folder<\li>
<li>Open Terminal<\li>
<li>cd auction<\li>
<li>cd test-network<\li>
<li>./network.sh up createChannel -ca<\li>
<li>./network.sh deployCC -ccn auction -ccp ../auction-simple/chaincode-go/ -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')"<\li>
<li>cd ..<\li>
<li>cd auction-simple/application-javascript<\li>
<li>npm install<\li>
<li>node enrollAdmin.js org1<\li>
<li>node enrollAdmin.js org2<\li>
<\ul>

### To get the Website Running
<ul>
<li>Go to "Website" Folder<\li>
<li>Open Terminal<\li>
<li>go mod tidy<\li>
<li>go run main.go<\li>
<\ul>
  
### To end the Hyperleadger Fabric Blockchain
<ul>
<li>Go to Simple Auction Folder<\li>
<li>Open Terminal<\li>
<li>cd auction-simple/application-javascript<\li>
<li>rm -rf wallet<\li>
<li>cd ..<\li>
<li>cd ..<\li>
<li>cd test-network<\li>
<li>./network.sh down<\li>
<\ul>
