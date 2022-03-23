const completedButton = document.getElementById("complete-btn");
const logout = document.getElementById('logout');

const handleCompleteAuction = async () => {
    let data = {
        completed: toggle
    }
    let response = await fetch("/seller", {
        method: "PUT",
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify(data)
    })
    console.log(response);

}
// const handleLogout = async () => {
//     let data = {
//         bidder_id: -1,
//         seller_id: -1
//     }

//     let response = await fetch("/auction", {
//         method: "PUT",
//         headers: { 'Content-Type': 'application/json'},
//         body: JSON.stringify(data)

//     })
//     console.log(response)
// } 


completedButton.addEventListener("click", handleCompleteAuction);
// logout.addEventListener("click", handleLogout);