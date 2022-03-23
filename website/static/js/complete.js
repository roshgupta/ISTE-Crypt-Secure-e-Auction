const completedButton = document.getElementById("complete-btn");


const handleCompleteAuction = async () => {
    let data = {
        completed: completed.value
    }
    let response = await fetch("/seller", {
        method: "PUT",
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify(data)
    })
    console.log(response);

} 


completedButton.addEventListener("click", handleCompleteAuction);

