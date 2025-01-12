<script lang="ts">
    import Client, { type Message } from "$lib/connection";
    import { onMount } from "svelte";
    const serverURL = new URL("http://localhost:8080/ws");

    let client: Client = new Client(serverURL);
    client.onMessage((message) => {
        switch (message.messageType) {
            case "server.game.water.updated":
                console.info("Water updated!")
                game.water = message.water;
                break;
        }
    });

    let game = $state({
        water: 0,
    });

    onMount(async () => {
        client.connect();
    });
</script>

<div class="p-4">
    <h1 class="text-2xl">Water</h1>
    <p>{game.water}</p>
</div>
