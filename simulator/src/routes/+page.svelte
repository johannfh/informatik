<script lang="ts">
    import { onMount } from "svelte";
    import * as ex from "excalibur";
    import { loader } from "$lib/resources";
    import { inputEvents, MainScene, type AnimalOption, type SpawnAnimalEvent } from "$lib/main-scene";
    import { screenHeight, screenWidth } from "$lib/constants";

    let gameCanvas: HTMLCanvasElement;

    onMount(() => {
        let game = new ex.Engine({
            width: screenWidth,
            height: screenHeight,
            canvasElement: gameCanvas,
            scenes: {
                start: MainScene,
            },
        })

        game.start("start", {
            // name of the start scene 'start'
            loader, // Optional loader (but needed for loading images/sounds)
            inTransition: new ex.FadeInOut({
                // Optional in transition
                duration: 200,
                direction: "in",
                color: ex.Color.ExcaliburBlue,
            }),
        }).then(() => console.log("game started"));

        function spawn({
            animal, position
        }: Omit<SpawnAnimalEvent, "type">) {
            console.log(`spawn: ${animal} @ ${position.x}:${position.y}`)
            inputEvents.update(v => [...v, ({ type: "spawnAnimalEvent", animal, position })])
        }

        game.input.pointers.on("down", ({ pointerId, worldPos }) => {
            spawn({ animal: spawnOption, position: worldPos })
        })
    })

    let spawnOption: AnimalOption = "wolf";
</script>

<div class="flex flex-col items-center justify-center w-screen h-screen gap-4">
    {@render gameControls()}
    <div class="p-2 border-2 rounded-md border-zinc-800">
        <canvas bind:this={gameCanvas}></canvas>
    </div>
</div>

{#snippet gameControls()}
    <div class="flex flex-row gap-4">
        <select bind:value={spawnOption} class="p-2 rounded-md">
            <option value="wolf">Wolf</option>
            <option value="sheep">Sheep</option>
            <option value="fox">Fox</option>
            <option value="chicken">Chicken</option>
        </select>
        <div>
            <h1 class="text-2xl font-bold underline">Controls</h1>
            <ul class="p-2">
                <li>LMB to spawn animal at cursor position</li>
            </ul>
        </div>
    </div>
{/snippet}