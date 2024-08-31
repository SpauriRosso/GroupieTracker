function submitSearch(artistName) {
    //  Makes sure the form exist
    var form = document.querySelector('form');
    // Same thing for the search bar
    var input = document.querySelector('#search-bar');
    //  Fill form with given variable
    input.value = artistName;
    // Bake cookies
    form.submit();
}