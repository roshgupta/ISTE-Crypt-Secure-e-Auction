<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <link rel="stylesheet" href="https://pro.fontawesome.com/releases/v5.10.0/css/all.css"
        integrity="sha384-AYmEC3Yw5cVb3ZcuHtOA93w35dYTsvhLPVnYs9eStHfGJvOvKxVfELGroGkvsg+p" crossorigin="anonymous" />
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Poppins&display=swap" rel="stylesheet">
    <link rel="stylesheet" href="../static/css/styles.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Bidder's Dashboard</title>
</head>

<body>
    <div class="nav">
        <a href="/" class="brand">
            <img src="../static/assets/block.png" alt="" class="bugimg">
            <h1 class="navBrand">Secure e-Auction</h1>
        </a>

        <div class="userProfile">
            <i class="fas fa-user-circle user-pro"></i>
            <a href="/" id="logout">
                
                <i class="fas fa-power-off power"></i>
            </a>

        </div>
    </div>
    <br>
    <h1 class="head">Bidder's Dashboard</h1>
    <h2 class="bidHead"><i class="fas fa-gavel"></i>Bids</h2>
    {{range .auctions}}
    <div class="Product1">
        <span class="productTitle">{{.Name}}</span>
        <span class="productDesc">{{.Description}}</span>
        <form action="" method="POST">
            <input name="bidAmount" id="bidAmount" class="amount-container" type="number"
                placeholder="Enter a value"></input>
            <div class="buttonDiv"> <input type="submit" value="Bid" class="bidbtn">
        </form>
    </div>
    </div>
    {{end}}

</body>

</html>