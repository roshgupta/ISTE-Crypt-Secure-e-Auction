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
      <div class="bidder-names">
        <div class="bidder-name">
          <h2 class="bidder-names-heading">Bidder names</h2>
          <h2 class="bidder-names-heading">Bidding amount</h2>
        </div>
      {{range .bidders}}
        <div class="bidder-name">
          <p>{{.Bidder_id}}</p>
          <p>{{.Amount}}<span> &#8377; </span></p>
        </div>
      {{end}}
      <form class="end-auction" action="" method="POST">
        <button name="completed" value="1">End auction</button>
      </form>



    </section>
    <img src="static/img/element.png" alt="" class="side-image right">
  </main>
</body>

</html>