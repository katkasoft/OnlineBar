document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("login-button").addEventListener("click", submitLogin);
    document.getElementById("register-button").addEventListener("click", redirectToRegister);
  });
  
  function submitRegister() {
    var email = document.getElementById("email").value;
    var name = document.getElementById("username").value;
    var password = document.getElementById("password").value;
  
    var credentials = {
      name: name,
      password: password,
      email: email,
      os: "Arch Linux" //I use arch btw
    };
  
    fetch("http://localhost:8080/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(credentials)
    })
      .then(response => response.text())
      .then(data => {
        console.log("Server response:", data);
        window.location.href = "/login"
      })
      .catch(error => {
        console.error("Error:", error);
      });
  }
  
  function redirectToRegister() {
    console.log("Redirecting to registration page");
    window.location.href = "/registration"
  }