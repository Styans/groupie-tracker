function handleSearch() {
    var input, filter, artists, img, i, txtValue, resultContainer;
    input = document.getElementById('search');
    filter = input.value.trim().toUpperCase();
    artists = document.getElementById('artist-container').getElementsByTagName('a');
    resultContainer = document.getElementById('search-results');

    // Clear the list of links under the search bar if the search input is empty
    if (filter === '') {
        resultContainer.innerHTML = '';
        showArtists(); // Show the artist links
        return;  // Exit the function early if the search input is empty
    }

    // Clear previous results
    resultContainer.innerHTML = '';

    for (i = 0; i < artists.length; i++) {
        img = artists[i].getElementsByTagName("img")[0];
        txtValue = img.alt || img.getAttribute('alt');
        if (txtValue.toUpperCase().indexOf(filter) > -1) {
            artists[i].style.display = "";

            // Create a dynamic list item (li) for each artist
            var artistListItem = document.createElement('li');
            artistListItem.classList.add('artist-list-item');

            // Create a dynamic anchor element for each artist
            var artistLink = document.createElement('a');
            artistLink.href = '/artist?id=' + artists[i].getAttribute('id');
            artistLink.textContent = txtValue;
            artistLink.classList.add('artist-link', 'your-other-class');

            // Add a click event listener to navigate to the artist page
            artistLink.addEventListener('click', function(event) {
                event.preventDefault();
                window.location.href = this.href;
            });

            // Append the anchor element to the list item
            artistListItem.appendChild(artistLink);

            // Append the list item to the result container
            resultContainer.appendChild(artistListItem);
        } else {
            artists[i].style.display = "none";
        }
    }
}

// Function to show the artist links
function showArtists() {
    var artistLinks = document.getElementById('artist-container').getElementsByTagName('a');
    for (var i = 0; i < artistLinks.length; i++) {
        artistLinks[i].style.display = "";
    }
}
