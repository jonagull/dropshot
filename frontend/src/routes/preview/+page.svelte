<script lang="ts">
    import type { PageData } from "../$types";
    import Gallery from "$lib/components/Gallery.svelte";

    export let data: PageData;
</script>

<div class="wrapper">
    <Gallery images={data.files} />
</div>

<style>
    .wrapper {
        display: flex;
        justify-content: center;
        padding-top: 100px;
    }

    :root {
        background-color: black;
    }
</style>
