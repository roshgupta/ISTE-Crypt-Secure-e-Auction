const email = document.getElementById("email");
const password = document.getElementById("password");
const loginSellerButton = document.getElementById("login-seller");

const emailRegex = /^([a-zA-Z0-9_\.\-]+)@([a-zA-Z0-9_\.\-]+)\.([a-zA-Z]+){1,10}$/;

const handleloginSeller = async () => {
  if (password.value.length > 3 && emailRegex.test(email.value)) {
    email.classList.remove("error");
    password.classList.remove("error");
    let data = {
      password: password.value,
      email: email.value
    }
    let response = await fetch("/login-seller", {
      method: "POST",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    })

    console.log("Request complete! response:", response);

  } else {
    alert("Enter Valid Input");
  }

}

password.addEventListener("blur", () => {
  if (password.value.length < 3) {
    password.classList.add("error");
  } else {
    password.classList.remove("error");
  }
});
email.addEventListener("blur", () => {
  if (!emailRegex.test(email.value)) {
    email.classList.add("error");
  } else {
    email.classList.remove("error");
  }
});

loginSellerButton.addEventListener("click", handleloginSeller);