document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("login-button").addEventListener("click", submitLogin);
    document.getElementById("register-button").addEventListener("click", redirectToRegister);
  });
  
  function submitLogin() {
    var name = document.getElementById("username").value;
    var password = document.getElementById("password").value;
  
    var credentials = {
      name: name,
      password: password
    };
  
    fetch("http://localhost:8080/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(credentials)
    })
      .then(response => response.text())
      .then(data => {
        console.log("Server response:", data);
        window.location.href = "/"
      })
      .catch(error => {
        console.error("Error:", error);
      });
  }
  
  function redirectToRegister() {и
    console.log("Redirecting to registration page");
    window.location.href = "/registration"
  }
  