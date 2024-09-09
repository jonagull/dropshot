<script lang="ts">
    import { fade } from "svelte/transition"; // Import the fade transition
    import { onDestroy, onMount } from "svelte";

    export let images: string[];

    let currentImage = 0;
    let interval: any;

    onMount(() => {
        console.log("IMAGES INSIDE GALLERY", images);
        // Automatically switch images every 3 seconds
        interval = setInterval(() => {
            currentImage = (currentImage + 1) % images.length;
        }, 10000); // Adjust the interval time for faster or slower switching
    });

    // Clear interval when component is destroyed
    onDestroy(() => {
        clearInterval(interval);
    });
</script>

<div class="gallery">
    {#each images as image, i}
        {#if i === currentImage}
            <!-- svelte-ignore a11y-img-redundant-alt -->
            <img
                src={`http://localhost:8080/images/${image}`}
                alt="Gallery image"
                transition:fade={{ duration: 1000 }}
            />
        {/if}
    {/each}
</div>

<style>
    .gallery {
        position: relative;
        width: 600px; /* Adjust based on your desired gallery size */
        height: 400px; /* Adjust based on your desired gallery size */
        overflow: hidden;
    }

    .gallery img {
        position: absolute;
        width: 100%;
        height: 100%;
        object-fit: cover;
    }
</style>
