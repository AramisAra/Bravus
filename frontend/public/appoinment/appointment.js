document.getElementById('appointment-form').addEventListener('submit', function(event) {
    event.preventDefault();

    const loadingSpinner = document.getElementById('loading-spinner');
    const messageDiv = document.getElementById('message');

    loadingSpinner.style.display = 'block';

    const formData = new FormData(this);

    fetch('https://example.com/submit-appointment', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        loadingSpinner.style.display = 'none';
        if (data.success) {
            messageDiv.innerHTML = '<p class="success">Appointment booked successfully!</p>';
            showConfirmationPopup();
        } else {
            messageDiv.innerHTML = `<p class="error">${data.message}</p>`;
        }
    })
    .catch(error => {
        loadingSpinner.style.display = 'none';
        messageDiv.innerHTML = `<p class="error">An error occurred: ${error.message}</p>`;
    });
});

function showConfirmationPopup() {
    const popup = document.getElementById('popup');
    popup.style.display = 'block';
    setTimeout(() => {
        popup.style.display = 'none';
    }, 3000);
}