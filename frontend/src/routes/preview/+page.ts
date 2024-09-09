import type { PageLoad } from "./$types";

export const load = (async () => {
    let files: any = [];
    try {
        const response = await fetch("http://localhost:8080/uploads"); // Replace with your API endpoint
        if (!response.ok) {
            throw new Error("Failed to fetch uploads");
        }
        const data = await response.json(); // Parse the JSON response
        files = data.images; // Access the "images" field in the response
        console.log("files", files);
    } catch (err) {
        console.error(err.message); // Handle error
    }

    return {
        files: files,
    };
}) satisfies PageLoad;
