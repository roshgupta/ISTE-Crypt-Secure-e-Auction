const completedButton = document.getElementById("complete-btn");


const handleCompleteAuction = async () => {
    let data = {
        completed: completed.value
    }
    let response = await fetch("/seller", {
        method: "PUT",
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify(data);
    })
    console.log(response);

} else {
    console.log("Error");
}


completedButton.addEventListener("click", handleCompleteAuction);

