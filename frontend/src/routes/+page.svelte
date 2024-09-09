<script lang="ts">
    import Images from "$lib/components/Gallery.svelte";
    import { BASE_URL } from "$lib/constants";
    import { onMount } from "svelte";

    let file: any = null;
    let message = "";
    let errorMessage = "";
    let images: any = [];
    let files: any = [];
    let error: any;

    onMount(() => {
        fetchImages();
    });

    // Function to handle form submission
    async function uploadImage(event: any) {
        event.preventDefault();

        if (!file) {
            message = "Please select an image.";
            return;
        }

        const formData = new FormData();
        formData.append("file", file);

        try {
            const response = await fetch(`${BASE_URL}/upload`, {
                method: "POST",
                body: formData,
            });

            if (response.ok) {
                message = "Image uploaded successfully!";
                fetchImages();
            } else {
                const errorMsg = await response.text();
                message = `Error: ${errorMsg}`;
                fetchImages();
            }
        } catch (error) {
            fetchImages();
            message = `Error: ${(error as Error).message}`;
        }
    }

    async function fetchImages() {
        try {
            const response = await fetch("http://localhost:8080/uploads"); // Replace with your API endpoint
            if (!response.ok) {
                throw new Error("Failed to fetch uploads");
            }
            const data = await response.json(); // Parse the JSON response
            files = data.images; // Access the "images" field in the response
            console.log("files", files);
        } catch (err) {
            error = err.message; // Handle error
        }
    }

    // Function to handle file input
    function handleFileChange(event: any) {
        file = event.target.files[0];
        message = ""; // Clear previous messages
    }
</script>

<!-- Form for uploading an image -->
<div class="form__wrapper">
    <form on:submit={uploadImage}>
        <label for="file">Choose an image:</label>
        <input
            type="file"
            id="file"
            accept="image/*"
            on:change={handleFileChange}
        />
        <button type="submit">Upload</button>
    </form>
</div>

{#if images.length > 0}
    <div><h1>REEEEEEEEEE</h1></div>
{/if}

<p class="message">{message}</p>

<style lang="scss">
    .form__wrapper {
        margin-top: 20px;
        display: flex;
        justify-content: center;
        form {
            display: flex;
            flex-direction: column;
            border: 1px solid black;
            padding: 20px;
        }
    }
    .message {
        margin-top: 10px;
        color: red;
    }
</style>
