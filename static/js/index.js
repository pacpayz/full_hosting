function infoEntered() {
	event.preventDefault(); // prevent form from submitting normally
  
	// get form data
	const email = document.getElementById("email").value;
	const firstName = document.getElementById("firstName").value;
	const lastName = document.getElementById("lastName").value;
	const phoneNumber = document.getElementById("phoneNumber").value;
	const billingAddress = document.getElementById("billingAddress").value;
  
	// create object with form data
	const formData = {
	  email: email,
	  firstName: firstName,
	  lastName: lastName,
	  phoneNumber: phoneNumber,
	  billingAddress: billingAddress
	};
  
	// send form data to server
	const xhr = new XMLHttpRequest();
	xhr.open("POST", "/basic-submit");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	xhr.send(JSON.stringify(formData));
  
	// handle server response
	xhr.onload = function() {
	  if (xhr.status === 200) {
		alert(xhr.responseText);
	  } else {
		alert("Error submitting form. Please try again.");
	  }
	};
  }
  