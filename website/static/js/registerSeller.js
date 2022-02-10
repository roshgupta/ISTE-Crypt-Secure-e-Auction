const userName = document.getElementById("name");
const email = document.getElementById("email");
const password = document.getElementById("password");
const registerSellerButton = document.getElementById("register-seller");
let regex = /^([a-zA-Z]){3,20}$/

const handleRegisterSeller = async () => {
  if (userName.value.length > 3 && password.value.length > 3 && regex.test(email.value)) {
    let data = {
      userName: userName.value,
      password: password.value,
      email: email.value
    }
    let response = await fetch("/register-seller", {
      method: "POST",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    })

    console.log("Request complete! response:", response);

  }
}