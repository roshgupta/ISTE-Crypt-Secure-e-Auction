const userName = document.getElementById("name");
const email = document.getElementById("email");
const password = document.getElementById("password");
const registerBidderButton = document.getElementById("register-bidder");
const nameRegex = /^([a-zA-Z0-9]){3,30}$/;
const emailRegex = /^([a-zA-Z0-9_\.\-]+)@([a-zA-Z0-9_\.\-]+)\.([a-zA-Z]+){1,10}$/;

const handleRegisterSeller = async () => {
  if (nameRegex.test(userName.value) && password.value.length > 3 && emailRegex.test(email.value)) {
    email.classList.remove("error");
    userName.classList.remove("error");
    password.classList.remove("error");
    let data = {
      userName: userName.value,
      password: password.value,
      email: email.value
    }
    let response = await fetch("/registration-bidder", {
      method: "POST",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    })

    console.log("Request complete! response:", response);


  } else {
    console.error("Enter Valid Input");
  }

}
userName.addEventListener("blur", (e) => {
  if (!nameRegex.test(userName.value)) {
    userName.classList.add("error");
  } else {
    userName.classList.remove("error");
  }
});
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

registerBidderButton.addEventListener("click", handleRegisterSeller);