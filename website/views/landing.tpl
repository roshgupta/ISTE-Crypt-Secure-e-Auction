<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Poppins&display=swap" rel="stylesheet">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Roboto&display=swap" rel="stylesheet">

    <link rel="stylesheet" href="static/css/styles.css">
    <title>Secure Auction</title>
</head>

<body>
    <div class='mainland'>
        <div class="glass">
            <div class="title">
                <div class="text">
                    <h1 class="LandingHeader">
                        <img src="static/assets/block.png" class="bugimg"> Secure e-Auction
                    </h1>
                    <h3 class="headerSmol">
                        Blockchain
                    </h3>
                </div>
            </div>
            <div class="sections">
                <section class="org">
                    <img src="static/assets/seller.svg" class="orgImg"></img>
                    <a href="/login-seller">
                        <button class="create-btn">Login as Seller</button>
                    </a>
                    <p>
                        Are you a seller? Login here <br /> and sell your items for the <br /> highest bid
                    </p>
                    <!-- <a href="/login-seller">
                        <button class="login-btn">Admin Login</button>
                    </a> -->
                </section>
                <section class="user">
                    <img src="static/assets/bid.svg" class="userImg"></img>

                    <a href="/login-bidder">
                        <button class="create-btn2"> Login as Bidder </button>
                    </a>
                    <p>
                        Are you a bidder? Login here, <br /> and bid your price and buy<br /> various items</p>
                    <!-- <a href="/login-bidder">
                        <button class="login-btn">User Login</button>
                    </a> -->
                </section>
            </div>
        </div>
    </div>
</body>

</html>