<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="static/css/login.css">
  <title>Login for Bidder</title>
</head>

<body>
  <main class="main-container">
    <img src="static/img/element.png" alt="" class="side-image left">
    <aside class="sidebar-image-container-login">
    </aside>
    <section class="container-box">
      <div class="container">
        <h1>Login for Bidder</h1>
        <form action="" id="bidder-login" method="POST">
          <div class="inputs">
            <div class="input">
              <label for="email">Email</label>
              <input type="email" id="email" name="email" placeholder="Enter your email id">
            </div>
            <div class="input">
              <label for="password">Password</label>
              <input type="password" id="password" name="password" placeholder="Enter password">
            </div>
          </div>

          <div class="controls">
            <span>Don't have an account?</span>
            <div class="buttons">
              <button id="login-bidder" class="button">Login</button>
              <a href="/registration-bidder" class="button">Register</a>
            </div>
        </form>
      </div>



      </div>
    </section>
    <img src="static/img/element.png" alt="" class="side-image right">
  </main>
  <script src="static/js/loginBidder.js"></script>

</body>

</html>