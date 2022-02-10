<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="static/css/login.css">
  <title>Registration for Seller</title>
</head>

<body>
  <main class="main-container">
    <img src="static/img/element.png" alt="" class="side-image left">
    <aside class="sidebar-image-container">
    </aside>
    <section class="container-box">
      <div class="container">
        <h1>Registration for Seller</h1>
        <div class="inputs">
          <div class="input">
            <label for="name">Name</label>
            <input type="text" id="name" name="name" placeholder="Enter your name" class="">
          </div>
          <div class="input">
            <label for="email">Email</label>
            <input type="email" id="email" name="email" placeholder="Enter your email id" class="">
          </div>
          <div class="input">
            <label for="password">Password</label>
            <input type="password" id="password" name="password" placeholder="Create a password" class="">
          </div>
        </div>
        <div class="controls">
          <span>Already have an account?</span>
          <div class="buttons">
            <button id="register-seller" class="button">Register</button>
            <a href="/login-seller" class="button">Login</a>
          </div>
        </div>

      </div>
    </section>
    <img src="static/img/element.png" alt="" class="side-image right">
  </main>
  <script src="static/js/registerSeller.js"></script>
</body>

</html>