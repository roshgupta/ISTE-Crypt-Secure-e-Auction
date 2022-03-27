<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="static/css/bid-details.css">
  <title>Bid Details</title>
</head>

<body>
  <main class="main-container">
    <img src="static/img/element.png" alt="" class="side-image left">
    <section class="container-box">
      <div>
        <h1 class="bid-details-heading">{{.auctions.Name}}</h1>
        <p class="bid-description">{{.auctions.Description}}</p>
      </div>
      <form action="" method="POST" class="bid-form">
        <input type="number" name="bidAmount" id="bidAmount">
        <input type="submit" value="Bid">
      </form>



    </section>
    <img src="static/img/element.png" alt="" class="side-image right">
  </main>
</body>

</html>