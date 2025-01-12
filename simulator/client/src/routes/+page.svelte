<script lang="ts">
    import Client, { type Message } from "$lib/connection";
    import type { Entity, EntitySchema, EntityType } from "$lib/connection/schemas/messages";
    import { onMount } from "svelte";
    const serverURL = new URL("http://localhost:8080/ws");

    let client: Client = new Client(serverURL);
    client.onMessage((message) => {
        switch (message.messageType) {
            case "server.game.water.updated":
                console.info("Water updated!");
                game.water = message.water;
                break;
            case "server.game.entities.updated":
                game.entities = message.entities;
                break;
        }
    });

    type Game = {
        water: number;
        entities: Entity[];
    };

    function GetEntities(game: Game, type: EntityType) {
        return game.entities.filter((e) => e.type == type)
    }

    let game = $state<Game>({
        water: 0,
        entities: [],
    });

    let water = $state(5);

    onMount(async () => {
        client.connect();
    });

    function addWater(water: number) {
        console.info("adding water");
        client.send({
            messageType: "client.game.water.add",
            water: water,
        });
    }
</script>

<div class="flex size-min flex-col items-center p-4 bg-gray-100 dark:bg-zinc-700 bg-opacity-80 backdrop-blur-xl rounded-lg gap-y-2">
    <h1 class="text-2xl">Water</h1>
    <div class="p-2">
        <p class="h-12 flex items-center justify-center p-4 w-full bg-gray-50 rounded-md dark:bg-zinc-900">{game.water}</p>
        <div class="flex flex-row items-center">
            <input class="h-12 rounded-md bg-gray-50 dark:bg-zinc-900 p-4" min={0} type="number" bind:value={water} />
            <dir class="flex flex-col gap-2">
                <button onclick={() => addWater(water)}>Add</button>
                <button onclick={() => addWater(-water)}>Subtract</button>
            </dir>
        </div>
    </div>
</div>

<div>
    <h1 class="text-2xl">Entities: {game.entities.length}</h1>
    <div class="w-full flex flex-row justify-around flex-wrap">
        <div>Wolves: {GetEntities(game, "wolf").length}</div>
        <div>Sheeps: {GetEntities(game, "sheep").length}</div>
        <div>Foxes: {GetEntities(game, "fox").length}</div>
        <div>Chickens: {GetEntities(game, "chicken").length}</div>
    </div>
    <div class="grid grid-cols-4">
        {#each game.entities as { type }}
            <div class="py-2 px-4">
                <div class="text-lg">{type}</div>
            </div>
        {/each}
    </div>
</div>
