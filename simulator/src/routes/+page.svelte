<script lang="ts">
    import { onMount } from "svelte";
    import * as ex from "excalibur";
    import { loader } from "$lib/resources";
    import {
        inputEvents,
        MainScene,
        type KillAnimalEvent,
        type SpawnAnimalEvent,
    } from "$lib/main-scene";
    import { type AnimalOption } from "$lib/types";
    import { screenHeight, screenWidth } from "$lib/constants";
    import { Chicken, Fox, Sheep, Wolf } from "$lib";

    let gameCanvas: HTMLCanvasElement;

    onMount(() => {
        let game = new ex.Engine({
            width: screenWidth,
            height: screenHeight,
            canvasElement: gameCanvas,
            scenes: {
                start: MainScene,
            },
        });

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

        function spawn({ animal, position }: Omit<SpawnAnimalEvent, "type">) {
            console.log(`spawn: ${animal.name} @ ${position.x}:${position.y}`);
            inputEvents.update((v) => [...v, { type: "spawnAnimalEvent", animal, position }]);
        }

        function kill({ position }: Omit<KillAnimalEvent, "type">) {
            console.log(`kill: ${position.x}:${position.y}`);
            inputEvents.update((v) => [...v, { type: "killAnimalEvent", position }]);
        }

        game.input.pointers.on("down", ({ button, pointerId, worldPos }) => {
            switch (button) {
                case "Left":
                    spawn({ animal: spawnOption, position: worldPos });
                    break;

                case "Right":
                    kill({ position: worldPos });
                    break;
            }
        });
    });

    let spawnOption: AnimalOption = Wolf;
</script>

<div class="flex flex-col items-center justify-center w-screen h-screen gap-4">
    {@render gameControls()}
    <div class="p-2 border-2 rounded-md border-zinc-800">
        <canvas bind:this={gameCanvas}></canvas>
    </div>
</div>

{#snippet gameControls()}
    <div class="flex flex-col gap-2">
        <h1 class="text-2xl font-bold underline w-[500px]">Controls</h1>
        <ul class="control-list">
            <li>
                <code>LMB</code>
                <span>to spawn </span>
                <select
                    bind:value={spawnOption}
                    class="p-2 rounded-md bg-gray-200 cursor-pointer hover:bg-gray-300 outline-none"
                >
                    <option value={Wolf}>Wolf</option>
                    <option value={Sheep}>Sheep</option>
                    <option value={Fox}>Fox</option>
                    <option value={Chicken}>Chicken</option>
                </select>
                <span>at cursor position</span>
            </li>
            <li class="unimplemented"><code>RMB</code> to kill animals at cursor position</li>
        </ul>
    </div>
{/snippet}

<style>
    ul.control-list {
        @apply p-2 flex flex-col gap-2;
        > li {
            @apply bg-gray-100 p-2 flex flex-row items-center gap-x-1;
            > code {
                @apply font-bold bg-gray-200 p-2 rounded-full;
            }
        }

        > li.unimplemented {
        }
    }
</style>
