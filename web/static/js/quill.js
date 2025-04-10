var quill = new Quill('#editor', {
    theme: 'snow',
    modules: {
      toolbar: [
        [{ 'header': '1' }, { 'header': '2' }],
        [{ 'font': [] }],
        [{ 'align': [] }],
        ['bold', 'italic', 'underline', 'strike'],
        ['blockquote', 'code-block'],
        [{ 'list': 'ordered'}, { 'list': 'bullet' }],
        ['link', 'image', 'video'],
        ['clean'], // To remove formatting
        [{ 'color': [] }, { 'background': [] }], // Text color and background
        [{ 'indent': '-1'}, { 'indent': '+1' }], // Indentation controls
        ['blockquote', 'code-block'], // Allow block quotes and code blocks
        [{ 'direction': 'rtl' }] // Right-to-left support
      ]
    }
  });

// Save button logic (optional - for handling saving data)
document.getElementById('save-btn').addEventListener('click', function () {
    const title = document.getElementById('title').value;
    const description = document.getElementById('description').value;
    const coverImage = document.getElementById('cover-image').files[0];
    const content = quill.root.innerHTML; // Get the content from the Quill editor

    // Example of sending form data (including file upload) to server
    const formData = new FormData();
    formData.append('title', title);
    formData.append('description', description);
    formData.append('content', content);
    if (coverImage) {
        formData.append('cover-image', coverImage);
    }

    // Example of submitting the form data to an API endpoint
    // fetch('/api/articles/save', {
    //   method: 'POST',
    //   body: formData
    // })
    // .then(response => response.json())
    // .then(data => {
    //     console.log('Success:', data);
    // })
    // .catch(error => {
    //     console.error('Error:', error);
    // });
});